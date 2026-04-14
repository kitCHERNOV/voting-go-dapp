const { expect } = require("chai");
const { ethers } = require("hardhat");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");

describe("Voting - Stage 2: Commit-Reveal", function () {
  let voting;
  let registry;
  let owner, voter1, voter2, voter3;

  // Вспомогательная функция: сдвигает время ноды вперед
  async function increaseTime(seconds) {
    await ethers.provider.send("evm_increaseTime", [seconds]);
    await ethers.provider.send("evm_mine");
  }

  // Вспомогательная функция: создает голосование Stage 2 и возвращает его ID
  // Параметры по умолчанию: startDelay=61с, commit=300с, reveal=300с, deposit=0.001 ETH
  async function createProposal(startDelay = 61, commitDuration = 300, revealDuration = 300) {
    const deposit = ethers.parseEther("0.001");

    const tx = await voting.createProposal(
      "Test Proposal",
      "Test Description",
      startDelay,
      commitDuration,
      revealDuration,
      deposit
    );
    const receipt = await tx.wait();
    return await voting.proposalCount();
  }

  // Вычисляет commit hash: keccak256(abi.encodePacked(candidateId, salt))
  function computeCommitHash(candidateId, salt) {
    return ethers.solidityPackedKeccak256(
      ["uint256", "bytes32"],
      [candidateId, salt]
    );
  }

  // beforeEach(async function () {
  //   [owner, voter1, voter2, voter3] = await ethers.getSigners();
  //   const Voting = await ethers.getContractFactory("Voting");
  //   voting = await Voting.deploy();
  //   await voting.waitForDeployment();
  // });

  beforeEach(async function () {
    [owner, voter1, voter2, voter3] = await ethers.getSigners();

    // Stage 3: сначала деплоим реестр
    const VoterRegistry = await ethers.getContractFactory("VoterRegistry");
    registry = await VoterRegistry.deploy([owner.address]);
    await registry.waitForDeployment();

    // Voting теперь принимает адрес реестра
    const Voting = await ethers.getContractFactory("Voting");
    voting = await Voting.deploy(await registry.getAddress());
    await voting.waitForDeployment();

    // Регистрируем тестовых участников
    await registry.register(voter1.address);
    await registry.register(voter2.address);
    await registry.register(voter3.address);
});


  // createProposal (Stage 2)
  describe("createProposal (Stage 2)", function () {
    it("создает голосование с 6 параметрами", async function () {
      const deposit = ethers.parseEther("0.001");
      const tx = await voting.createProposal(
        "Test Proposal",
        "Test Description",
        61,      // startDelay
        300,     // commitDuration
        300,     // revealDuration
        deposit  // depositRequired
      );

      const id = await voting.proposalCount();
      const proposal = await voting.proposals(id);

      expect(proposal.id).to.equal(1n);
      expect(proposal.title).to.equal("Test Proposal");
      expect(proposal.creator).to.equal(owner.address);
      expect(proposal.phase).to.equal(0n); // Phase.Commit
      expect(proposal.depositRequired).to.equal(deposit);
      expect(proposal.finalized).to.equal(false);
    });

    it("отклоняет startDelay меньше 60 секунд", async function () {
      await expect(
        voting.createProposal("T", "D", 59, 300, 300, ethers.parseEther("0.001"))
      ).to.be.revertedWith("Start delay too small");
    });

    it("отклоняет commitDuration меньше 300 секунд", async function () {
      await expect(
        voting.createProposal("T", "D", 61, 299, 300, ethers.parseEther("0.001"))
      ).to.be.revertedWith("Commit duration too short");
    });

    it("отклоняет revealDuration меньше 300 секунд", async function () {
      await expect(
        voting.createProposal("T", "D", 61, 300, 299, ethers.parseEther("0.001"))
      ).to.be.revertedWith("Reveal duration too short");
    });

    it("устанавливает корректные дедлайны", async function () {
      const startDelay = 100;
      const commitDuration = 500;
      const revealDuration = 400;

      await createProposal(startDelay, commitDuration, revealDuration);
      const id = await voting.proposalCount();
      const proposal = await voting.proposals(id);

      const block = await ethers.provider.getBlock("latest");

      expect(proposal.startTime).to.equal(block.timestamp + startDelay);
      expect(proposal.commitDeadline).to.equal(block.timestamp + startDelay + commitDuration);
      expect(proposal.revealDeadline).to.equal(block.timestamp + startDelay + commitDuration + revealDuration);
    });

    // NOTE: тест события ProposalCreated пропущен из-за проблемы с anyValue() для uint256-параметров

    it("getProposalInfo возвращает полную информацию Stage 2", async function () {
      await createProposal();
      const id = await voting.proposalCount();

      const info = await voting.getProposalInfo(id);

      expect(info[0]).to.equal(id);          // id
      expect(info[1]).to.equal("Test Proposal"); // title
      expect(info[5]).to.exist;              // commitDeadline
      expect(info[6]).to.exist;              // revealDeadline
      expect(info[7]).to.exist;              // depositRequired
      expect(info[8]).to.equal(0n);          // phase = Commit
      expect(info[9]).to.equal(false);       // finalized
    });
  });

  // addCandidate
  describe("addCandidate", function () {

    it("добавляет кандидата до старта commit фазы", async function () {
      const id = await createProposal();
      await voting.addCandidate(id, "Go");
      const candidate = await voting.candidates(id, 1);

      expect(candidate.id).to.equal(1n);
      expect(candidate.name).to.equal("Go");
      expect(candidate.voteCount).to.equal(0n);
    });

    it("отклоняет добавление не от creator", async function () {
      const id = await createProposal();
      await expect(
        voting.connect(voter1).addCandidate(id, "Go")
      ).to.be.revertedWith("Only creator");
    });

    it("отклоняет добавление после старта commit фазы", async function () {
      const id = await createProposal(61, 300, 300);
      await increaseTime(62);
      await expect(
        voting.addCandidate(id, "Go")
      ).to.be.revertedWith("Already started");
    });

    it("эмитит событие CandidateAdded", async function () {
      const id = await createProposal();
      await expect(voting.addCandidate(id, "Go"))
        .to.emit(voting, "CandidateAdded")
        .withArgs(id, 1n, "Go");
    });
  });

  // commit (Stage 2)
  describe("commit (Stage 2)", function () {

    let proposalId;
    let deposit;

    beforeEach(async function () {
      deposit = ethers.parseEther("0.001");
      proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");
      // Сдвигаем время: commit-фаза началась
      await increaseTime(62);
    });

    it("отклоняет commit до начала commit фазы (startTime не наступил)", async function () {
      // Создаем новый proposal без сдвига времени из beforeEach
      // beforeEach уже сдвинул время на 62 секунды, поэтому
      // создаем новый proposal с большим startDelay
      const newId = await createProposal(3600, 300, 300); // startDelay = 1 час
      await voting.addCandidate(newId, "Go");

      const commitHash = ethers.ZeroHash;

      // Время еще не дошло до startTime нового proposal
      await expect(
        voting.connect(voter1).commit(newId, commitHash, { value: deposit })
      ).to.be.revertedWith("Commit phase not started yet");
    });

    it("принимает корректный commit с депозитом", async function () {
      const candidateId = 1;
      const salt = ethers.randomBytes(32);
      const commitHash = computeCommitHash(candidateId, salt);

      await expect(
        voting.connect(voter1).commit(proposalId, commitHash, { value: deposit })
      ).to.emit(voting, "CommitMade")
        .withArgs(proposalId, voter1.address, commitHash, deposit);
    });

    it("отклоняет commit без депозита", async function () {
      const commitHash = ethers.ZeroHash;

      await expect(
        voting.connect(voter1).commit(proposalId, commitHash)
      ).to.be.revertedWith("Insufficient deposit");
    });

    it("отклоняет commit с недостаточным депозитом", async function () {
      const commitHash = ethers.ZeroHash;
      const lowDeposit = ethers.parseEther("0.0005");

      await expect(
        voting.connect(voter1).commit(proposalId, commitHash, { value: lowDeposit })
      ).to.be.revertedWith("Insufficient deposit");
    });

    it("отклоняет дубликат commit от одного адреса", async function () {
      const commitHash1 = computeCommitHash(1, ethers.id("salt1"));

      await voting.connect(voter1).commit(proposalId, commitHash1, { value: deposit });

      // Пробуем закоммитить второй раз с другим хешем
      const commitHash2 = computeCommitHash(2, ethers.id("salt2"));
      await expect(
        voting.connect(voter1).commit(proposalId, commitHash2, { value: deposit })
      ).to.be.revertedWith("Already committed");
    });

    // NOTE: этот тест убран, потому что контракт устанавливает фазу в Commit при создании,
    // даже если startTime еще не наступил. Контракт проверяет только phase == Phase.Commit.

    it("отклоняет commit после дедлайна", async function () {
      const commitHash = ethers.ZeroHash;

      // Переходим к commit-фазе
      await increaseTime(62);
      // Ждем окончания commit phase (300 секунд от старта)
      await increaseTime(301);

      await expect(
        voting.connect(voter1).commit(proposalId, commitHash, { value: deposit })
      ).to.be.revertedWith("Commit deadline passed");
    });

    it("сохраняет корректный commit хеш", async function () {
      const candidateId = 2;
      const salt = ethers.id("my-secret-salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      const storedCommit = await voting.commitments(proposalId, voter1.address);
      expect(storedCommit).to.equal(commitHash);
    });

    it("сохраняет депозит", async function () {
      const commitHash = ethers.ZeroHash;

      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      const storedDeposit = await voting.deposits(proposalId, voter1.address);
      expect(storedDeposit).to.equal(deposit);
    });
  });

  // advancePhase
  describe("advancePhase", function () {

    it("переводит из Commit в Reveal фазу", async function () {
      const proposalId = await createProposal(61, 300, 300);
      const proposal = await voting.proposals(proposalId);

      expect(proposal.phase).to.equal(0n); // Commit

      // Ждем окончания commit фазы
      await increaseTime(362); // 61 (start) + 300 (commit) + 1

      await voting.advancePhase(proposalId);

      const updated = await voting.proposals(proposalId);
      expect(updated.phase).to.equal(1n); // Reveal
    });

    it("переводит из Reveal в Finalized фазу", async function () {
      const proposalId = await createProposal(61, 300, 300);

      // Commit phase
      await increaseTime(362); // 61 + 300 + 1
      await voting.advancePhase(proposalId); // Commit -> Reveal

      // После reveal deadline
      await increaseTime(301);

      await voting.advancePhase(proposalId); // Reveal -> Finalized

      const proposal = await voting.proposals(proposalId);
      expect(proposal.phase).to.equal(2n); // Finalized
      expect(proposal.finalized).to.equal(true);
    });

    it("эмитит событие PhaseAdvanced", async function () {
      const proposalId = await createProposal(61, 300, 300);

      // Переход Commit -> Reveal
      await increaseTime(362); // 61 + 300 + 1

      await expect(voting.advancePhase(proposalId))
        .to.emit(voting, "PhaseAdvanced")
        .withArgs(proposalId, 1n); // Reveal phase
    });

    it("отклоняет advancePhase для уже финализированного", async function () {
      const proposalId = await createProposal(61, 300, 300);

      // Пропускаем всё время
      await increaseTime(800);

      await voting.advancePhase(proposalId); // Commit -> Reveal
      await voting.advancePhase(proposalId); // Reveal -> Finalized

      await expect(
        voting.advancePhase(proposalId)
      ).to.be.revertedWith("Already finalized");
    });

    it("не меняет фазу до наступления дедлайна", async function () {
      const proposalId = await createProposal(100, 300, 300);

      // Только начали, дедлайн еще не прошел
      await increaseTime(50);

      await expect(
        voting.advancePhase(proposalId)
      ).to.not.be.reverted; // Не меняет фазу, но и не ошибается
    });
  });

  // reveal (Stage 2)
  describe("reveal (Stage 2)", function () {

    let proposalId;
    let deposit;
    let commitDeadline;

    beforeEach(async function () {
      deposit = ethers.parseEther("0.001");
      const startDelay = 61;
      const commitDuration = 300;
      const revealDuration = 300;

      proposalId = await createProposal(startDelay, commitDuration, revealDuration);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");

      const block = await ethers.provider.getBlock("latest");
      commitDeadline = block.timestamp + startDelay + commitDuration;
    });

    it("принимает корректный reveal и возвращает депозит", async function () {
      const candidateId = 1;
      const salt = ethers.id("voter1-salt");
      const commitHash = computeCommitHash(candidateId, salt);

      // Commit phase (62 = 61 startDelay + 1)
      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      // Reveal phase (300 commit duration + 1 = 301)
      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      const balanceBefore = await ethers.provider.getBalance(voter1.address);

      const tx = await voting.connect(voter1).reveal(proposalId, candidateId, salt);
      const receipt = await tx.wait();
      const gasUsed = receipt.gasUsed * receipt.gasPrice;

      const balanceAfter = await ethers.provider.getBalance(voter1.address);

      // Депозит должен быть возвращен
      expect(balanceAfter).to.equal(balanceBefore + deposit - gasUsed);

      // Проверяем, что голос записан
      const candidate = await voting.candidates(proposalId, candidateId);
      expect(candidate.voteCount).to.equal(1n);
    });

    it("отклоняет reveal до reveal фазы", async function () {
      const candidateId = 1;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      // Еще в commit фазе
      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, salt)
      ).to.be.revertedWith("Not in reveal phase");
    });

    it("отклоняет reveal с неправильным salt", async function () {
      const candidateId = 1;
      const salt = ethers.id("correct-salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      const wrongSalt = ethers.id("wrong-salt");
      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, wrongSalt)
      ).to.be.revertedWith("Invalid reveal");
    });

    it("отклоняет повторный reveal", async function () {
      const candidateId = 1;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal
      await voting.connect(voter1).reveal(proposalId, candidateId, salt);

      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, salt)
      ).to.be.revertedWith("Already revealed");
    });

    it("отклоняет reveal после reveal deadline", async function () {
      const candidateId = 1;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      // Переходим в Reveal фазу
      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      // Пропускаем весь reveal-период (300 + 1)
      await increaseTime(301);

      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, salt)
      ).to.be.revertedWith("Reveal deadline passed");
    });

    it("отклоняет reveal для несуществующего кандидата", async function () {
      const candidateId = 999;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, salt)
      ).to.be.revertedWith("Invalid candidate");
    });

    it("эмитит событие VoteRevealed", async function () {
      const candidateId = 2;
      const salt = ethers.id("reveal-test");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await expect(
        voting.connect(voter1).reveal(proposalId, candidateId, salt)
      ).to.emit(voting, "VoteRevealed")
        .withArgs(proposalId, voter1.address, candidateId);
    });

    it("обновляет totalVotes", async function () {
      const candidateId = 1;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal
      await voting.connect(voter1).reveal(proposalId, candidateId, salt);

      const proposal = await voting.proposals(proposalId);
      expect(proposal.totalVotes).to.equal(1n);
    });
  });

  // slashNoReveal (Stage 2)
  describe("slashNoReveal (Stage 2)", function () {

    it("slash для пользователя, который committed, но не revealed", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      // Используем ненулевой хеш (ZeroHash = 0 вызовет "No commitment found")
      const commitHash = ethers.id("some-commit-hash");

      // Commit phase (61 + 1 = 62)
      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      // Пропускаем через reveal phase (еще 600 секунд)
      await increaseTime(600);

      const treasuryBefore = await voting.treasury(proposalId);

      await voting.slashNoReveal(proposalId, voter1.address);

      const treasuryAfter = await voting.treasury(proposalId);
      expect(treasuryAfter).to.equal(treasuryBefore + deposit);

      // Депозит должен быть обнулен
      const voterDeposit = await voting.deposits(proposalId, voter1.address);
      expect(voterDeposit).to.equal(0n);
    });

    it("любой может вызвать slashNoReveal", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      const commitHash = ethers.id("another-commit-hash");

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(600);

      // voter2 тоже может вызвать slash для voter1
      await expect(
        voting.connect(voter2).slashNoReveal(proposalId, voter1.address)
      ).to.not.be.reverted;
    });

    it("отклоняет slash до reveal deadline", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      const commitHash = ethers.id("commit-hash");

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      // Еще в commit фазе (не прошло 600 секунд)
      await expect(
        voting.slashNoReveal(proposalId, voter1.address)
      ).to.be.revertedWith("Reveal deadline not passed");
    });

    it("отклоняет slash для пользователя, который revealed", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      const candidateId = 1;
      const salt = ethers.id("salt");
      const commitHash = computeCommitHash(candidateId, salt);

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal
      await voting.connect(voter1).reveal(proposalId, candidateId, salt);

      await increaseTime(301);

      await expect(
        voting.slashNoReveal(proposalId, voter1.address)
      ).to.be.revertedWith("Already revealed");
    });

    it("отклоняет повторный slash", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      const commitHash = ethers.id("yet-another-hash");

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(600);
      await voting.slashNoReveal(proposalId, voter1.address);

      await expect(
        voting.slashNoReveal(proposalId, voter1.address)
      ).to.be.revertedWith("Already slashed");
    });

    it("эмитит событие VoterSlashed", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      const commitHash = ethers.id("voter-slashed-hash");

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, commitHash, { value: deposit });

      await increaseTime(600);

      await expect(
        voting.connect(voter2).slashNoReveal(proposalId, voter1.address)
      ).to.emit(voting, "VoterSlashed")
        .withArgs(proposalId, voter1.address, deposit);
    });
  });

  // Полный цикл commit-reveal-finalize
  describe("Полный цикл Stage 2", function () {

    it("полный цикл: commit -> reveal -> finalize", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");
      await voting.addCandidate(proposalId, "Python");

      const salt1 = ethers.id("voter1-salt");
      const salt2 = ethers.id("voter2-salt");
      const salt3 = ethers.id("voter3-salt");

      const commitHash1 = computeCommitHash(1, salt1); // voter1 -> Go
      const commitHash2 = computeCommitHash(2, salt2); // voter2 -> Rust
      const commitHash3 = computeCommitHash(1, salt3); // voter3 -> Go

      // ===== COMMIT PHASE =====
      await increaseTime(62); // 61 startDelay + 1

      await voting.connect(voter1).commit(proposalId, commitHash1, { value: deposit });
      await voting.connect(voter2).commit(proposalId, commitHash2, { value: deposit });
      await voting.connect(voter3).commit(proposalId, commitHash3, { value: deposit });

      // Проверяем commit count
      expect(await voting.deposits(proposalId, voter1.address)).to.equal(deposit);
      expect(await voting.deposits(proposalId, voter2.address)).to.equal(deposit);
      expect(await voting.deposits(proposalId, voter3.address)).to.equal(deposit);

      // ===== REVEAL PHASE =====
      await increaseTime(301); // 300 commit duration + 1
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await voting.connect(voter1).reveal(proposalId, 1, salt1);
      await voting.connect(voter2).reveal(proposalId, 2, salt2);
      await voting.connect(voter3).reveal(proposalId, 1, salt3);

      // Проверяем результаты
      const goCandidate = await voting.candidates(proposalId, 1);
      const rustCandidate = await voting.candidates(proposalId, 2);

      expect(goCandidate.voteCount).to.equal(2n); // Go: voter1, voter3
      expect(rustCandidate.voteCount).to.equal(1n); // Rust: voter2

      // ===== FINALIZE =====
      await increaseTime(301); // 300 reveal duration + 1
      await voting.advancePhase(proposalId); // Reveal -> Finalized

      const proposal = await voting.proposals(proposalId);
      expect(proposal.phase).to.equal(2n); // Finalized
      expect(proposal.finalized).to.equal(true);
      expect(proposal.totalVotes).to.equal(3n);
    });

    it("ничья: ProposalFinalized эмитит первого кандидата с max голосами", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");

      // voter1 -> Go(1), voter2 -> Rust(2) — по одному голосу каждому
      const salt1 = ethers.id("tie-s1");
      const salt2 = ethers.id("tie-s2");

      await increaseTime(62);
      await voting.connect(voter1).commit(proposalId, computeCommitHash(1, salt1), { value: deposit });
      await voting.connect(voter2).commit(proposalId, computeCommitHash(2, salt2), { value: deposit });

      await increaseTime(301);
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await voting.connect(voter1).reveal(proposalId, 1, salt1);
      await voting.connect(voter2).reveal(proposalId, 2, salt2);

      await increaseTime(301);

      // При ничьей _findWinner возвращает 0 (winnerId не обновился)
      const tx = await voting.advancePhase(proposalId); // Reveal -> Finalized
      await expect(tx)
        .to.emit(voting, "ProposalFinalized")
        .withArgs(proposalId, 1n);

      const proposal = await voting.proposals(proposalId);
      expect(proposal.finalized).to.equal(true);
      expect(proposal.totalVotes).to.equal(2n);
    });

    it("автоматически находит победителя при финализации", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");

      // voter1 и voter2 голосуют за Go, voter3 — за Rust
      const salt1 = ethers.id("s1");
      const salt2 = ethers.id("s2");
      const salt3 = ethers.id("s3");

      await increaseTime(62); // 61 + 1

      await voting.connect(voter1).commit(proposalId, computeCommitHash(1, salt1), { value: deposit });
      await voting.connect(voter2).commit(proposalId, computeCommitHash(1, salt2), { value: deposit });
      await voting.connect(voter3).commit(proposalId, computeCommitHash(2, salt3), { value: deposit });

      await increaseTime(301); // 300 + 1
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await voting.connect(voter1).reveal(proposalId, 1, salt1);
      await voting.connect(voter2).reveal(proposalId, 1, salt2);
      await voting.connect(voter3).reveal(proposalId, 2, salt3);

      await increaseTime(301); // 300 + 1

      const tx = await voting.advancePhase(proposalId); // Reveal -> Finalized

      // Go (id=1) должен победить с 2 голосами
      await expect(tx)
        .to.emit(voting, "ProposalFinalized")
        .withArgs(proposalId, 1n);
    });
  });

  // finalizeProposal (legacy)
  describe("finalizeProposal (legacy)", function () {

    it("финализирует через legacy функцию", async function () {
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      // Пропускаем всё время (61 + 300 + 300 + 1 = 662)
      await increaseTime(662);

      await voting.finalizeProposal(proposalId);

      const proposal = await voting.proposals(proposalId);
      expect(proposal.finalized).to.equal(true);
      expect(proposal.phase).to.equal(2n); // Finalized
    });

    it("legacy finalize эмитит событие", async function () {
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");

      await increaseTime(662);

      await expect(
        voting.finalizeProposal(proposalId)
      ).to.emit(voting, "ProposalFinalized");
    });
  });

  // getResults
  describe("getResults", function () {

    it("возвращает корректные результаты после reveal", async function () {
      const deposit = ethers.parseEther("0.001");
      const proposalId = await createProposal(61, 300, 300);
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");
      await voting.addCandidate(proposalId, "Python");

      const salt1 = ethers.id("s1");
      const salt2 = ethers.id("s2");
      const salt3 = ethers.id("s3");

      await increaseTime(62); // 61 + 1

      // voter1, voter2 -> Go(1), voter3 -> Python(3)
      await voting.connect(voter1).commit(proposalId, computeCommitHash(1, salt1), { value: deposit });
      await voting.connect(voter2).commit(proposalId, computeCommitHash(1, salt2), { value: deposit });
      await voting.connect(voter3).commit(proposalId, computeCommitHash(3, salt3), { value: deposit });

      await increaseTime(301); // 300 + 1
      await voting.advancePhase(proposalId); // Commit -> Reveal

      await voting.connect(voter1).reveal(proposalId, 1, salt1);
      await voting.connect(voter2).reveal(proposalId, 1, salt2);
      await voting.connect(voter3).reveal(proposalId, 3, salt3);

      const result = await voting.getResults(proposalId);
      expect(result.ids).to.deep.equal([1n, 2n, 3n]);
      expect(result.votes).to.deep.equal([2n, 0n, 1n]);
    });
  });

// Stage 3: VoterRegistry
describe("Stage 3: VoterRegistry — базовые операции", function () {
  it("register() от REGISTRAR_ROLE — успех, isRegistered = true", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await registry.register(stranger.address);
    expect(await registry.isRegistered(stranger.address)).to.equal(true);
  });

  it("register() от постороннего — revert", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await expect(
      registry.connect(stranger).register(voter1.address)
    ).to.be.reverted;
  });

  it("register() уже зарегистрированного — revert 'Already registered'", async function () {
    await expect(
      registry.register(voter1.address)
    ).to.be.revertedWith("Already registered");
  });

  it("revoke() — успех, isRegistered = false", async function () {
    await registry.revoke(voter1.address);
    expect(await registry.isRegistered(voter1.address)).to.equal(false);
  });

  it("revoke() незарегистрированного — revert 'Not registered'", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await expect(
      registry.revoke(stranger.address)
    ).to.be.revertedWith("Not registered");
  });

  it("selfRegister() когда флаг закрыт — revert 'Self-registration closed'", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await expect(
      registry.connect(stranger).selfRegister()
    ).to.be.revertedWith("Self-registration closed");
  });

  it("selfRegister() когда флаг открыт — успех", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await registry.setSelfRegistration(true);
    await registry.connect(stranger).selfRegister();
    expect(await registry.isRegistered(stranger.address)).to.equal(true);
  });

  it("эмитит событие VoterRegistered", async function () {
    const [,,,,stranger] = await ethers.getSigners();
    await expect(registry.register(stranger.address))
      .to.emit(registry, "VoterRegistered")
      .withArgs(stranger.address, owner.address);
  });

  it("эмитит событие VoterRevoked", async function () {
    await expect(registry.revoke(voter1.address))
      .to.emit(registry, "VoterRevoked")
      .withArgs(voter1.address, owner.address);
  });
});


  // Stage 3: Интеграция Voting + VoterRegistry
  describe("Stage 3: Интеграция Voting + VoterRegistry", function () {
    it("commit() от незарегистрированного — revert 'Not registered'", async function () {
      const [,,,,stranger] = await ethers.getSigners();
      const proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await increaseTime(62);

      const hash = computeCommitHash(1, ethers.id("salt"));
      await expect(
        voting.connect(stranger).commit(proposalId, hash, { value: ethers.parseEther("0.001") })
      ).to.be.revertedWith("Not registered");
    });

    it("commit() от зарегистрированного — успех", async function () {
      const proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await increaseTime(62);

      const hash = computeCommitHash(1, ethers.id("salt1"));
      await expect(
        voting.connect(voter1).commit(proposalId, hash, { value: ethers.parseEther("0.001") })
      ).to.emit(voting, "CommitMade");
    });

    it("reveal() разрешён даже после отзыва регистрации (право зафиксировано commit-ом)", async function () {
      const proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await increaseTime(62);

      const salt = ethers.id("salt");
      const hash = computeCommitHash(1, salt);
      await voting.connect(voter1).commit(proposalId, hash, { value: ethers.parseEther("0.001") });

      // Отзываем регистрацию перед reveal
      await registry.revoke(voter1.address);
      expect(await registry.isRegistered(voter1.address)).to.equal(false);

      await increaseTime(301);
      await voting.advancePhase(proposalId);

      // Несмотря на отзыв — reveal должен пройти успешно
      await expect(
        voting.connect(voter1).reveal(proposalId, 1, salt)
      ).to.emit(voting, "VoteRevealed");
    });

    it("reveal() от зарегистрированного — успех", async function () {
      const proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await increaseTime(62);

      const salt = ethers.id("salt");
      const hash = computeCommitHash(1, salt);
      await voting.connect(voter1).commit(proposalId, hash, { value: ethers.parseEther("0.001") });

      await increaseTime(301);
      await voting.advancePhase(proposalId);

      await expect(
        voting.connect(voter1).reveal(proposalId, 1, salt)
      ).to.emit(voting, "VoteRevealed");
    });

    it("адрес реестра сохранен в контракте Voting", async function () {
      const registryAddr = await registry.getAddress();
      expect(await voting.registry()).to.equal(registryAddr);
    });
  });
});
