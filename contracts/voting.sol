// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
 
contract Voting {
 
    struct Proposal {
        uint256 id;
        string  title;
        string  description;
        address creator;
        uint256 startTime;      // unix timestamp
        uint256 endTime;        // unix timestamp
        bool    finalized;
        uint256 totalVotes;
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
 
    // proposalId => address => voted
    mapping(uint256 => mapping(address => bool)) public hasVoted;
 
    // Modifier: проверка активности голосования
    modifier proposalActive(uint256 id) {
        Proposal storage p = proposals[id];
        require(p.id != 0,                  'Proposal not found');
        require(!p.finalized,                'Already finalized');
        require(block.timestamp >= p.startTime,'Not started yet');
        require(block.timestamp <= p.endTime,  'Voting ended');
        _;
    }

    /// @notice Создаёт голосование. Возвращает ID для дальнейших вызовов.
    function createProposal(
        string calldata _title,
        string calldata _description,
        uint256 _startDelay,    // секунд до начала (минимум 60)
        uint256 _durationSec    // длительность в секундах
    ) external returns (uint256 proposalId) {
        require(_startDelay >= 60,    'Start delay too small');
        require(_durationSec >= 3600, 'Duration too short');
    
        proposalCount++;
        uint256 start = block.timestamp + _startDelay;
        uint256 end   = start + _durationSec;
    
        proposals[proposalCount] = Proposal({
            id:          proposalCount,
            title:       _title,
            description: _description,
            creator:     msg.sender,
            startTime:   start,
            endTime:     end,
            finalized:   false,
            totalVotes:  0
        });
    
        emit ProposalCreated(proposalCount, _title, msg.sender, start, end);
        return proposalCount;  // читаем из события, не угадываем ID!
    }

    /// @notice Только creator может добавить кандидата до старта
    function addCandidate(uint256 _proposalId, string calldata _name) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0,               'Proposal not found');
        require(p.creator == msg.sender, 'Only creator');        // ← важно!
        require(block.timestamp < p.startTime, 'Already started');
    
        candidateCount[_proposalId]++;
        uint256 cid = candidateCount[_proposalId];
        candidates[_proposalId][cid] = Candidate({id: cid, name: _name, voteCount: 0});
    
        emit CandidateAdded(_proposalId, cid, _name);
    }
 
    /// @notice Голосование — один адрес, один голос
    function vote(uint256 _proposalId, uint256 _candidateId)
        external proposalActive(_proposalId)
    {
        require(!hasVoted[_proposalId][msg.sender], 'Already voted');
        require(candidates[_proposalId][_candidateId].id != 0, 'Bad candidate');
    
        hasVoted[_proposalId][msg.sender] = true;
        candidates[_proposalId][_candidateId].voteCount++;
        proposals[_proposalId].totalVotes++;
    
        emit VoteCast(_proposalId, msg.sender, _candidateId);
    }
 
    /// @notice Публичная финализация — любой может вызвать после endTime
    function finalizeProposal(uint256 _proposalId) external {
        Proposal storage p = proposals[_proposalId];
        require(p.id != 0,              'Not found');
        require(!p.finalized,           'Already done');
        require(block.timestamp > p.endTime, 'Not ended yet');
    
        p.finalized = true;
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
        external view
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

    event ProposalCreated(
        uint256 indexed proposalId,
        string          title,
        address indexed creator,
        uint256         startTime,
        uint256         endTime
    );
    event CandidateAdded(
        uint256 indexed proposalId,
        uint256 indexed candidateId,
        string          name
    );
    event VoteCast(
        uint256 indexed proposalId,
        address indexed voter,
        uint256 indexed candidateId
    );
    event ProposalFinalized(
        uint256 indexed proposalId,
        uint256         winnerCandidateId
    );
}

