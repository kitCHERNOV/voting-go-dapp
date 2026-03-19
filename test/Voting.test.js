const { expect } = require("chai");
const { ethers } = require("hardhat");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");

describe("Voting", function () {
  let voting;
  let owner, voter1, voter2, voter3;

  // Вспомогательная функция — сдвигает время ноды вперёд
  async function increaseTime(seconds) {
    await ethers.provider.send("evm_increaseTime", [seconds]);
    await ethers.provider.send("evm_mine");
  }

  // Вспомогательная функция — создаёт голосование и возвращает его ID
  async function createProposal(startDelay = 61, duration = 3600) {
    const tx = await voting.createProposal(
      "Test Proposal",
      "Test Description",
      startDelay,
      duration
    );
    const receipt = await tx.wait();
    return await voting.proposalCount();
  }

  beforeEach(async function () {
    [owner, voter1, voter2, voter3] = await ethers.getSigners();
    const Voting = await ethers.getContractFactory("Voting");
    voting = await Voting.deploy();
    await voting.waitForDeployment();
  });

  // ─── createProposal ───────────────────────────────────────────────
  describe("createProposal", function () {

    it("создаёт голосование с корректными параметрами", async function () {
      const id = await createProposal();
      const proposal = await voting.proposals(id);

      expect(proposal.id).to.equal(1n);
      expect(proposal.title).to.equal("Test Proposal");
      expect(proposal.creator).to.equal(owner.address);
      expect(proposal.finalized).to.equal(false);
      expect(proposal.totalVotes).to.equal(0n);
    });

    it("отклоняет startDelay меньше 60 секунд", async function () {
      await expect(
        voting.createProposal("T", "D", 59, 3600)
      ).to.be.revertedWith("Start delay too small");
    });

    it("отклоняет duration меньше 3600 секунд", async function () {
      await expect(
        voting.createProposal("T", "D", 61, 3599)
      ).to.be.revertedWith("Duration too short");
    });

    it("увеличивает proposalCount", async function () {
      await createProposal();
      await createProposal();
      expect(await voting.proposalCount()).to.equal(2n);
    });

    // it("эмитит событие ProposalCreated", async function () {
    //   await expect(voting.createProposal("T", "D", 61, 3600))
    //     .to.emit(voting, "ProposalCreated")
    //     .withArgs(1n, "T", owner.address, await anyValue(), await anyValue());
    // });
    it("эмитит событие ProposalCreated", async function () {
    const tx = await voting.createProposal("T", "D", 61, 3600);
    const receipt = await tx.wait();
    const block = await ethers.provider.getBlock(receipt.blockNumber);

    await expect(tx)
        .to.emit(voting, "ProposalCreated")
        .withArgs(
        1n,
        "T",
        owner.address,
        block.timestamp + 61,
        block.timestamp + 61 + 3600
        );
    });

  });

  // ─── addCandidate ─────────────────────────────────────────────────
  describe("addCandidate", function () {

    it("добавляет кандидата до старта", async function () {
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

    it("отклоняет добавление после старта", async function () {
      const id = await createProposal();
      await increaseTime(62);
      await expect(
        voting.addCandidate(id, "Go")
      ).to.be.revertedWith("Already started");
    });

    it("отклоняет несуществующий proposal", async function () {
      await expect(
        voting.addCandidate(999, "Go")
      ).to.be.revertedWith("Proposal not found");
    });

    it("эмитит событие CandidateAdded", async function () {
      const id = await createProposal();
      await expect(voting.addCandidate(id, "Go"))
        .to.emit(voting, "CandidateAdded")
        .withArgs(id, 1n, "Go");
    });

  });

  // ─── vote ─────────────────────────────────────────────────────────
  describe("vote", function () {

    let proposalId;

    beforeEach(async function () {
      proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");
      // Сдвигаем время — голосование началось
      await increaseTime(62);
    });

    it("записывает голос корректно", async function () {
      await voting.connect(voter1).vote(proposalId, 1);
      const candidate = await voting.candidates(proposalId, 1);
      expect(candidate.voteCount).to.equal(1n);
    });

    it("увеличивает totalVotes", async function () {
      await voting.connect(voter1).vote(proposalId, 1);
      const proposal = await voting.proposals(proposalId);
      expect(proposal.totalVotes).to.equal(1n);
    });

    it("отклоняет двойное голосование", async function () {
      await voting.connect(voter1).vote(proposalId, 1);
      await expect(
        voting.connect(voter1).vote(proposalId, 1)
      ).to.be.revertedWith("Already voted");
    });

    it("отклоняет несуществующего кандидата", async function () {
      await expect(
        voting.connect(voter1).vote(proposalId, 999)
      ).to.be.revertedWith("Bad candidate");
    });

    it("отклоняет голос до старта", async function () {
      const newId = await createProposal();
      await voting.addCandidate(newId, "Go");
      await expect(
        voting.connect(voter1).vote(newId, 1)
      ).to.be.revertedWith("Not started yet");
    });

    it("отклоняет голос после окончания", async function () {
      await increaseTime(3601);
      await expect(
        voting.connect(voter1).vote(proposalId, 1)
      ).to.be.revertedWith("Voting ended");
    });

    it("эмитит событие VoteCast", async function () {
      await expect(voting.connect(voter1).vote(proposalId, 1))
        .to.emit(voting, "VoteCast")
        .withArgs(proposalId, voter1.address, 1n);
    });

  });

  // ─── finalizeProposal ─────────────────────────────────────────────
  describe("finalizeProposal", function () {

    let proposalId;

    beforeEach(async function () {
      proposalId = await createProposal();
      await voting.addCandidate(proposalId, "Go");
      await voting.addCandidate(proposalId, "Rust");
      await increaseTime(62);
      await voting.connect(voter1).vote(proposalId, 1);
      await voting.connect(voter2).vote(proposalId, 1);
      await voting.connect(voter3).vote(proposalId, 2);
      // Сдвигаем время — голосование закончилось
      await increaseTime(3601);
    });

    it("финализирует голосование", async function () {
      await voting.finalizeProposal(proposalId);
      const proposal = await voting.proposals(proposalId);
      expect(proposal.finalized).to.equal(true);
    });

    it("любой может финализировать", async function () {
      await expect(
        voting.connect(voter3).finalizeProposal(proposalId)
      ).to.not.be.reverted;
    });

    it("отклоняет повторную финализацию", async function () {
      await voting.finalizeProposal(proposalId);
      await expect(
        voting.finalizeProposal(proposalId)
      ).to.be.revertedWith("Already done");
    });

    it("отклоняет финализацию до окончания", async function () {
      const newId = await createProposal();
      await voting.addCandidate(newId, "Go");
      await increaseTime(62);
      await expect(
        voting.finalizeProposal(newId)
      ).to.be.revertedWith("Not ended yet");
    });

    it("эмитит событие ProposalFinalized с правильным победителем", async function () {
      await expect(voting.finalizeProposal(proposalId))
        .to.emit(voting, "ProposalFinalized")
        .withArgs(proposalId, 1n); // кандидат 1 (Go) набрал 2 голоса
    });

  });

  // ─── getResults ───────────────────────────────────────────────────
  describe("getResults", function () {

    it("возвращает корректные результаты после голосования", async function () {
      const id = await createProposal();
      await voting.addCandidate(id, "Go");
      await voting.addCandidate(id, "Rust");
      await voting.addCandidate(id, "Python");
      await increaseTime(62);

      await voting.connect(voter1).vote(id, 1);
      await voting.connect(voter2).vote(id, 1);
      await voting.connect(voter3).vote(id, 3);

      const result = await voting.getResults(id);
      expect(result.ids).to.deep.equal([1n, 2n, 3n]);
      expect(result.votes).to.deep.equal([2n, 0n, 1n]);
    });

    it("возвращает пустые массивы если нет кандидатов", async function () {
      const id = await createProposal();
      const result = await voting.getResults(id);
      expect(result.ids).to.deep.equal([]);
      expect(result.votes).to.deep.equal([]);
    });

  });

});