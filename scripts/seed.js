const { ethers } = require("hardhat");
const { address } = require("../deploy.json");

async function main() {
  const [owner] = await ethers.getSigners();
  const voting = await ethers.getContractAt("Voting", address);

  // Создаём голосование: старт через 61 сек, длительность 2 часа
  const tx = await voting.createProposal(
    "Лучший язык программирования",
    "Голосуем за лучший язык 2024",
    61,
    7200
  );
  const receipt = await tx.wait();

  // Получаем ID из события
//   const event = receipt.logs
//     .map(log => { try { return voting.interface.parseLog(log) } catch { return null } })
//     .find(e => e && e.name === "ProposalCreated");

  const proposalId = await voting.proposalCount();
  console.log("Proposal ID:", proposalId.toString());

  // Добавляем кандидатов
  await (await voting.addCandidate(proposalId, "Go")).wait();
  console.log("Added: Go");

  await (await voting.addCandidate(proposalId, "Rust")).wait();
  console.log("Added: Rust");

  await (await voting.addCandidate(proposalId, "Python")).wait();
  console.log("Added: Python");

  console.log("Done! Proposal ID:", proposalId.toString());
}

main().catch((e) => { console.error(e); process.exit(1); });