// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./VoterRegistry.sol";

contract Voting {
    // ===== STAGE 3; Voter Registry =====
    VoterRegistry public immutable registry;

    // ===== STAGE 2: Phase Enum =====
    enum Phase { Commit, Reveal, Finalized }

    struct Proposal {
        // Original Stage 1 fields
        uint256 id;
        string  title;
        string  description;
        address creator;
        uint256 startTime;       // unix timestamp (commit start time)
        uint256 endTime;         // unix timestamp (kept for compatibility, but not used)
        bool    finalized;
        uint256 totalVotes;

        // Stage 2: Commit-Reveal fields
        uint256 commitDeadline;  // unix timestamp - end of commit phase
        uint256 revealDeadline;  // unix timestamp - end of reveal phase
        uint256 depositRequired; // Required deposit in wei (prevents last-revealer attack)
        Phase   phase;           // Current voting phase
    }

    struct Candidate {
        uint256 id;
        string  name;
        uint256 voteCount;
    }

    uint256 public proposalCount;

    // proposalId => Proposal
    mapping(uint256 => Proposal) public proposals;

    // proposalId => candidateId => Candidate
    mapping(uint256 => mapping(uint256 => Candidate)) public candidates;

    // per-proposal счётчик кандидатов (НЕ глобальный!)
    mapping(uint256 => uint256) public candidateCount;

    // ===== STAGE 1: hasVoted (kept for backward compatibility) =====
    // proposalId => address => voted
    mapping(uint256 => mapping(address => bool)) public hasVoted;

    // ===== STAGE 2: Commit-Reveal mappings =====
    // Commitment hashes: proposalId => voter => commitHash
    mapping(uint256 => mapping(address => bytes32)) public commitments;

    // Deposits held: proposalId => voter => amount (wei)
    mapping(uint256 => mapping(address => uint256)) public deposits;

    // Reveal status: proposalId => voter => revealed
    mapping(uint256 => mapping(address => bool)) public revealed;

    // Slashed deposits (treasury): proposalId => totalWeiSlashed
    mapping(uint256 => uint256) public treasury;

    // ===== STAGE 3: модификатор регистрации =====
    modifier onlyRegistered() {
        require(registry.isRegistered(msg.sender), "Not registered");
        _;
    }

    // Modifier: проверка активности голосования
    modifier proposalActive(uint256 id) {
        Proposal storage p = proposals[id];
        require(p.id != 0, "Proposal not found");
        require(p.phase != Phase.Finalized, "Already finalized");
        // Phase-specific timing handled in individual functions
        _;
    }

    // ===== STAGE 3: конструктор принимает адрес реестра =====
    constructor(address _registry) {
        registry = VoterRegistry(_registry);
    }

    /// @notice Создаёт голосование с commit-reveal схемой. Возвращает ID.
    /// @param _startDelay Seconds until commit phase starts (min 60)
    /// @param _commitDuration Commit phase duration in seconds (min 300)
    /// @param _revealDuration Reveal phase duration in seconds (min 300)
    /// @param _depositRequired Required deposit in wei to prevent last-revealer attack
    function createProposal(
        string calldata _title,
        string calldata _description,
        uint256 _startDelay,
        uint256 _commitDuration,
        uint256 _revealDuration,
        uint256 _depositRequired
    ) external returns (uint256 proposalId) {
        require(_startDelay >= 60, "Start delay too small");
        require(_commitDuration >= 300, "Commit duration too short");
        require(_revealDuration >= 300, "Reveal duration too short");

        proposalCount++;
        uint256 commitStart = block.timestamp + _startDelay;
        uint256 commitEnd = commitStart + _commitDuration;
        uint256 revealEnd = commitEnd + _revealDuration;

        proposals[proposalCount] = Proposal({
            id:               proposalCount,
            title:            _title,
            description:      _description,
            creator:          msg.sender,
            startTime:        commitStart,
            endTime:          revealEnd,       // kept for backward compatibility
            finalized:        false,
            totalVotes:       0,
            commitDeadline:   commitEnd,
            revealDeadline:   revealEnd,
            depositRequired:  _depositRequired,
            phase:            Phase.Commit     // Start in commit phase
        });

        emit ProposalCreated(
            proposalCount,
            _title,
            msg.sender,
            commitStart,
            commitEnd,
            revealEnd,
            _depositRequired
        );
        return proposalCount;
    }

    /// @notice Только creator может добавить кандидата до начала commit фазы
    function addCandidate(uint256 _proposalId, string calldata _name) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");
        require(p.creator == msg.sender, "Only creator");
        require(block.timestamp < p.startTime, "Already started");

        candidateCount[_proposalId]++;
        uint256 cid = candidateCount[_proposalId];
        candidates[_proposalId][cid] = Candidate({id: cid, name: _name, voteCount: 0});

        emit CandidateAdded(_proposalId, cid, _name);
    }

    // ===== STAGE 2 + STAGE 3: Commit Phase =====

    /// @notice Commit голоса с депозитом. Только зарегистрированные участники.
    /// @param _proposalId Proposal ID
    /// @param _commitHash keccak256(abi.encodePacked(candidateId, salt))
    function commit(uint256 _proposalId, bytes32 _commitHash)
        external
        payable
        onlyRegistered
    {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");
        require(p.phase == Phase.Commit, "Not in commit phase");
        require(block.timestamp >= p.startTime, "Commit phase not started yet");
        require(block.timestamp <= p.commitDeadline, "Commit deadline passed");
        require(commitments[_proposalId][msg.sender] == bytes32(0), "Already committed");
        require(msg.value >= p.depositRequired, "Insufficient deposit");

        commitments[_proposalId][msg.sender] = _commitHash;
        deposits[_proposalId][msg.sender] = msg.value;

        emit CommitMade(_proposalId, msg.sender, _commitHash, msg.value);
    }

    // ===== STAGE 2: + STAGE 3 Reveal Phase =====

    /// @notice Reveal a previously committed vote. Call during reveal phase.
    /// @param _proposalId Proposal ID
    /// @param _candidateId The candidate ID being voted for
    /// @param _salt The salt used when committing (must match!)
    function reveal(uint256 _proposalId, uint256 _candidateId, bytes32 _salt)
        external
    {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");
        require(p.phase == Phase.Reveal, "Not in reveal phase");
        require(block.timestamp <= p.revealDeadline, "Reveal deadline passed");
        require(revealed[_proposalId][msg.sender] == false, "Already revealed");

        // Verify commitment matches
        bytes32 computedHash = keccak256(abi.encodePacked(_candidateId, _salt));
        require(commitments[_proposalId][msg.sender] == computedHash, "Invalid reveal");

        // Mark as revealed and get deposit
        revealed[_proposalId][msg.sender] = true;
        uint256 deposit = deposits[_proposalId][msg.sender];
        deposits[_proposalId][msg.sender] = 0;

        // Record the vote
        require(candidates[_proposalId][_candidateId].id != 0, "Invalid candidate");
        candidates[_proposalId][_candidateId].voteCount++;
        proposals[_proposalId].totalVotes++;

        // Also update hasVoted for Stage 1 compatibility
        hasVoted[_proposalId][msg.sender] = true;

        // Refund deposit
        payable(msg.sender).transfer(deposit);

        emit VoteRevealed(_proposalId, msg.sender, _candidateId);
    }

    // ===== STAGE 2: Slashing =====

    /// @notice Slash a voter who committed but never revealed. Anyone can call after reveal deadline.
    /// @param _proposalId Proposal ID
    /// @param _voter Address of the voter to slash
    function slashNoReveal(uint256 _proposalId, address _voter) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");
        require(block.timestamp > p.revealDeadline, "Reveal deadline not passed");
        require(commitments[_proposalId][_voter] != bytes32(0), "No commitment found");
        require(revealed[_proposalId][_voter] == false, "Already revealed");
        require(deposits[_proposalId][_voter] > 0, "Already slashed");

        uint256 penalty = deposits[_proposalId][_voter];
        deposits[_proposalId][_voter] = 0;
        treasury[_proposalId] += penalty;

        emit VoterSlashed(_proposalId, _voter, penalty);
    }

    // ===== STAGE 2: Phase Transitions =====

    /// @notice Advance the proposal phase. Can be called by anyone after deadlines pass.
    /// @param _proposalId Proposal ID
    function advancePhase(uint256 _proposalId) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");
        require(p.phase != Phase.Finalized, "Already finalized");

        if (p.phase == Phase.Commit && block.timestamp > p.commitDeadline) {
            p.phase = Phase.Reveal;
            emit PhaseAdvanced(_proposalId, uint8(Phase.Reveal));
        } else if (p.phase == Phase.Reveal && block.timestamp > p.revealDeadline) {
            p.phase = Phase.Finalized;
            p.finalized = true;
            uint256 winner = _findWinner(_proposalId);
            emit ProposalFinalized(_proposalId, winner);
        }
    }

    // ===== STAGE 1: Legacy finalizeProposal (kept for compatibility) =====

    /// @notice Публичная финализация — любой может вызвать после reveal deadline
    function finalizeProposal(uint256 _proposalId) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Not found");
        require(!p.finalized, "Already done");
        require(block.timestamp > p.revealDeadline, "Not ended yet");

        p.finalized = true;
        p.phase = Phase.Finalized;
        uint256 winner = _findWinner(_proposalId);
        emit ProposalFinalized(_proposalId, winner);
    }

    /// @notice Внутренняя функция поиска победителя по максимальному voteCount
    /// @return winnerId ID кандидата с наибольшим числом голосов (0 если никто не голосовал)
    function _findWinner(uint256 _proposalId) internal view returns (uint256 winnerId) {
        uint256 count = candidateCount[_proposalId];
        uint256 maxVotes = 0;

        for (uint256 i = 1; i <= count; i++) {
            uint256 votes = candidates[_proposalId][i].voteCount;
            if (votes > maxVotes) {
                maxVotes = votes;
                winnerId = i;
            }
        }
        // winnerId == 0 означает ничью или отсутствие голосов
        return winnerId;
    }

    /// @notice Возвращает результаты голосования — массивы ID и голосов
    function getResults(uint256 _proposalId)
        external
        view
        returns (uint256[] memory ids, uint256[] memory votes)
    {
        uint256 count = candidateCount[_proposalId];
        ids   = new uint256[](count);
        votes = new uint256[](count);
        for (uint256 i = 1; i <= count; i++) {
            ids[i - 1]   = i;
            votes[i - 1] = candidates[_proposalId][i].voteCount;
        }

        return (ids, votes);
    }

    // ===== STAGE 2: Get Phase Info =====

    /// @notice Get detailed proposal information including phase
    function getProposalInfo(uint256 _proposalId)
        external
        view
        returns (
            uint256 id,
            string memory title,
            string memory description,
            address creator,
            uint256 startTime,
            uint256 commitDeadline,
            uint256 revealDeadline,
            uint256 depositRequired,
            uint8 phase,
            bool finalized,
            uint256 totalVotes
        )
    {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0, "Proposal not found");

        return (
            p.id,
            p.title,
            p.description,
            p.creator,
            p.startTime,
            p.commitDeadline,
            p.revealDeadline,
            p.depositRequired,
            uint8(p.phase),
            p.finalized,
            p.totalVotes
        );
    }

    // ===== Events =====

    event ProposalCreated(
        uint256 indexed proposalId,
        string          title,
        address indexed creator,
        uint256         startTime,
        uint256         commitDeadline,
        uint256         revealDeadline,
        uint256         depositRequired
    );

    event CandidateAdded(
        uint256 indexed proposalId,
        uint256 indexed candidateId,
        string          name
    );

    // ===== STAGE 1: Legacy event (kept for compatibility) =====
    event VoteCast(
        uint256 indexed proposalId,
        address indexed voter,
        uint256 indexed candidateId
    );

    // ===== STAGE 2: Commit-Reveal events =====
    event CommitMade(
        uint256 indexed proposalId,
        address indexed voter,
        bytes32         commitHash,
        uint256         deposit
    );

    event VoteRevealed(
        uint256 indexed proposalId,
        address indexed voter,
        uint256 indexed candidateId
    );

    event VoterSlashed(
        uint256 indexed proposalId,
        address indexed voter,
        uint256         penalty
    );

    event PhaseAdvanced(
        uint256 indexed proposalId,
        uint8           newPhase
    );

    event ProposalFinalized(
        uint256 indexed proposalId,
        uint256         winnerCandidateId
    );
}
