// web3.js — подключение MetaMask и отправка транзакций
// Чтение данных — только через api.js

// ─── Конфигурация ─────────────────────────────────────────────────────────

const NETWORK_CONFIG = {
  chainId: '0x7A69',         // 31337 — Hardhat localhost
  chainName: 'Hardhat Local',
  rpcUrl: 'http://127.0.0.1:8545',
};

// Адреса контрактов — заполняются при инициализации из /api/proposals
// (Go сервер знает адреса из deploy.json)
let CONTRACT_ADDRESS = '';
let REGISTRY_ADDRESS = '';

// ─── ABI контрактов (только нужные функции) ───────────────────────────────

const VOTING_ABI = [
  {
    name: 'commit',
    type: 'function',
    stateMutability: 'payable',
    inputs: [
      { name: '_proposalId', type: 'uint256' },
      { name: '_commitHash', type: 'bytes32' },
    ],
    outputs: [],
  },
  {
    name: 'reveal',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: '_proposalId', type: 'uint256' },
      { name: '_candidateId', type: 'uint256' },
      { name: '_salt', type: 'bytes32' },
    ],
    outputs: [],
  },
  {
    name: 'advancePhase',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [{ name: '_proposalId', type: 'uint256' }],
    outputs: [],
  },
  {
    name: 'createProposal',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: '_title', type: 'string' },
      { name: '_description', type: 'string' },
      { name: '_startDelay', type: 'uint256' },
      { name: '_commitDuration', type: 'uint256' },
      { name: '_revealDuration', type: 'uint256' },
      { name: '_depositRequired', type: 'uint256' },
    ],
    outputs: [{ name: 'proposalId', type: 'uint256' }],
  },
  {
    name: 'addCandidate',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: '_proposalId', type: 'uint256' },
      { name: '_name', type: 'string' },
    ],
    outputs: [],
  },
];

const REGISTRY_ABI = [
  {
    name: 'register',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [{ name: '_voter', type: 'address' }],
    outputs: [],
  },
  {
    name: 'selfRegister',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [],
    outputs: [],
  },
];

// ─── Состояние кошелька ───────────────────────────────────────────────────

let _provider = null;
let _signer   = null;
let _account  = null;

// ─── Проверка MetaMask ────────────────────────────────────────────────────

function isMetaMaskInstalled() {
  return typeof window.ethereum !== 'undefined';
}

function getAccount() {
  return _account;
}

// ─── Подключение кошелька ─────────────────────────────────────────────────

async function connectWallet() {
  if (!isMetaMaskInstalled()) {
    throw new Error('MetaMask не установлен');
  }

  _provider = new ethers.BrowserProvider(window.ethereum);
  const accounts = await _provider.send('eth_requestAccounts', []);
  _account = accounts[0];
  _signer = await _provider.getSigner();

  // Слушаем смену аккаунта
  window.ethereum.on('accountsChanged', (accounts) => {
    _account = accounts[0] || null;
    _signer = null;
    window.dispatchEvent(new CustomEvent('walletChanged', { detail: { account: _account } }));
  });

  return _account;
}

// ─── Получить провайдер и подписанта ─────────────────────────────────────

async function getSigner() {
  if (!_signer) {
    if (!isMetaMaskInstalled()) throw new Error('MetaMask не установлен');
    _provider = new ethers.BrowserProvider(window.ethereum);
    _signer = await _provider.getSigner();
    _account = await _signer.getAddress();
  }
  return _signer;
}

// ─── Инициализация адресов контрактов ────────────────────────────────────
// Получаем адрес реестра из первого proposal или из отдельного endpoint
// Пока задаём вручную через атрибут data-* в HTML или window

function setContractAddresses(votingAddr, registryAddr) {
  CONTRACT_ADDRESS = votingAddr;
  REGISTRY_ADDRESS = registryAddr;
}

// ─── Транзакции ───────────────────────────────────────────────────────────

// Commit голоса
async function commitVote(proposalId, commitHash, depositWei) {
  const signer = await getSigner();
  const contract = new ethers.Contract(CONTRACT_ADDRESS, VOTING_ABI, signer);

  const tx = await contract.commit(proposalId, commitHash, {
    value: BigInt(depositWei),
  });
  return tx.wait();
}

// Reveal голоса
async function revealVote(proposalId, candidateId, salt) {
  const signer = await getSigner();
  const contract = new ethers.Contract(CONTRACT_ADDRESS, VOTING_ABI, signer);

  const tx = await contract.reveal(proposalId, candidateId, salt);
  return tx.wait();
}

// Смена фазы голосования
async function advancePhase(proposalId) {
  const signer = await getSigner();
  const contract = new ethers.Contract(CONTRACT_ADDRESS, VOTING_ABI, signer);

  const tx = await contract.advancePhase(proposalId);
  return tx.wait();
}

// Создание голосования
async function createProposal(title, description, startDelay, commitDuration, revealDuration, depositWei) {
  const signer = await getSigner();
  const contract = new ethers.Contract(CONTRACT_ADDRESS, VOTING_ABI, signer);

  const tx = await contract.createProposal(
    title, description,
    BigInt(startDelay), BigInt(commitDuration),
    BigInt(revealDuration), BigInt(depositWei)
  );
  return tx.wait();
}

// Добавление кандидата
async function addCandidate(proposalId, name) {
  const signer = await getSigner();
  const contract = new ethers.Contract(CONTRACT_ADDRESS, VOTING_ABI, signer);

  const tx = await contract.addCandidate(proposalId, name);
  return tx.wait();
}

// Регистрация участника (только admin с REGISTRAR_ROLE)
async function registerVoter(voterAddress) {
  const signer = await getSigner();
  const contract = new ethers.Contract(REGISTRY_ADDRESS, REGISTRY_ABI, signer);

  const tx = await contract.register(voterAddress);
  return tx.wait();
}

// Самостоятельная регистрация
async function selfRegister() {
  const signer = await getSigner();
  const contract = new ethers.Contract(REGISTRY_ADDRESS, REGISTRY_ABI, signer);

  const tx = await contract.selfRegister();
  return tx.wait();
}

// ─── Обработка ошибок транзакций ─────────────────────────────────────────

function parseContractError(err) {
  // Пользователь отклонил транзакцию в MetaMask
  if (err.code === 4001 || err.code === 'ACTION_REJECTED') {
    return 'Транзакция отклонена пользователем';
  }
  // Revert с сообщением из контракта
  if (err.reason) return err.reason;
  if (err.message) {
    const match = err.message.match(/execution reverted: (.+?)"/);
    if (match) return match[1];
  }
  return 'Ошибка транзакции';
}