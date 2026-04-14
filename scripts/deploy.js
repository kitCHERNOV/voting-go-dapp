const { ethers } = require("hardhat");
const fs = require("fs");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying from:", deployer.address);

  // 1. Деплоим VoterRegistry — deployer получает REGISTRAR_ROLE
  const VoterRegistry = await ethers.getContractFactory("VoterRegistry");
  const registry = await VoterRegistry.deploy([deployer.address]);
  await registry.waitForDeployment();
  const registryAddress = await registry.getAddress();
  console.log("VoterRegistry deployed to:", registryAddress);

  // 2. Деплоим Voting с адресом реестра
  const Voting = await ethers.getContractFactory("Voting");
  const voting = await Voting.deploy(registryAddress);
  await voting.waitForDeployment();
  const votingAddress = await voting.getAddress();
  console.log("Voting deployed to:", votingAddress);

  // 3. Сохраняем оба адреса
  fs.writeFileSync("deploy.json", JSON.stringify({
    address: votingAddress,
    registryAddress: registryAddress,
  }, null, 2));

  console.log("Saved to deploy.json");
}

main().catch((e) => { console.error(e); process.exit(1); });