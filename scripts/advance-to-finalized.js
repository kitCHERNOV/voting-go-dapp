const { ethers, network } = require("hardhat");

async function main() {
  const proposalId = 1;

  const deploy = require("../deploy.json");
  const votingAddress = deploy.address;

  const Voting = await ethers.getContractFactory("Voting");
  const voting = Voting.attach(votingAddress);

  const info = await voting.getProposalInfo(proposalId);

  const revealDeadline = info.revealDeadline;

  console.log("Reveal deadline:", revealDeadline.toString());

  await network.provider.send("evm_setNextBlockTimestamp", [
    Number(revealDeadline) + 1,
  ]);

  await network.provider.send("evm_mine");

  console.log("Time moved after revealDeadline");

//   const tx = await voting.advancePhase(proposalId);
//   await tx.wait();

//   console.log("Phase advanced to finalized");

//   const phase = await voting.getPhase(proposalId);
//   console.log("Current phase:", phase.toString());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});