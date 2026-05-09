const { ethers, network } = require("hardhat");

async function main() {
  const proposalId = 1;

  const deploy = require("../deploy.json");
  const votingAddress = deploy.address;

  const Voting = await ethers.getContractFactory("Voting");
  const voting = Voting.attach(votingAddress);

  const info = await voting.getProposalInfo(proposalId);

  const commitDeadline = info.commitDeadline;

  console.log("Commit deadline:", commitDeadline.toString());

  await network.provider.send("evm_setNextBlockTimestamp", [
    Number(commitDeadline) + 1,
  ]);

  await network.provider.send("evm_mine");

  console.log("Time moved after commitDeadline");

//   const tx = await voting.advancePhase(proposalId);
//   await tx.wait();

//   console.log("Phase advanced to reveal");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});