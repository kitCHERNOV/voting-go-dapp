const { ethers } = require("hardhat");
const { address, registryAddress } = require("../deploy.json");

async function main() {
  const [owner, voter1, voter2, voter3] = await ethers.getSigners();

  const voting = await ethers.getContractAt("Voting", address);
  const registry = await ethers.getContractAt("VoterRegistry", registryAddress);

  // Регистрируем тестовых участников
  await (await registry.register(voter1.address)).wait();
  console.log("Registered:", voter1.address);

  await (await registry.register(voter2.address)).wait();
  console.log("Registered:", voter2.address);

  await (await registry.register(voter3.address)).wait();
  console.log("Registered:", voter3.address);

  // Создаём тестовое голосование
  const depositRequired = ethers.parseEther("0.001");
  const tx = await voting.createProposal(
    "Лучший язык программирования",
    "Голосуем за лучший язык 2024",
    61,    // startDelay (секунд)
    7200,  // commitDuration (2 часа)
    3600,  // revealDuration (1 час)
    depositRequired
  );
  await tx.wait();

  const proposalId = await voting.proposalCount();
  console.log("Proposal ID:", proposalId.toString());

  // Добавляем кандидатов
  await (await voting.addCandidate(proposalId, "Go")).wait();
  console.log("Added: Go");

  await (await voting.addCandidate(proposalId, "Rust")).wait();
  console.log("Added: Rust");

  await (await voting.addCandidate(proposalId, "Python")).wait();
  console.log("Added: Python");

  console.log("\nDone!");
  console.log("Proposal ID:", proposalId.toString());
  console.log("Deposit required:", ethers.formatEther(depositRequired), "ETH");
  console.log("Registered voters:", voter1.address, voter2.address, voter3.address);
}

main().catch((e) => { console.error(e); process.exit(1); });