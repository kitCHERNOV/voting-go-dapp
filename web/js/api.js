// api.js — все запросы к Go-серверу (только чтение)
// Транзакции в блокчейн — только через web3.js

const API_BASE = '';  // пустая строка — запросы идут на тот же хост

// ─── Вспомогательная функция ──────────────────────────────────────────────

async function apiFetch(path, options = {}) {
  const res = await fetch(API_BASE + path, options);
  if (!res.ok) {
    const err = await res.json().catch(() => ({ error: res.statusText }));
    throw new Error(err.error || 'Ошибка запроса');
  }
  return res.json();
}

// ─── Голосования ──────────────────────────────────────────────────────────

// Получить список всех голосований
async function getProposals() {
  return apiFetch('/api/proposals');
}

// Получить детали голосования по ID
async function getProposal(id) {
  return apiFetch(`/api/proposals/${id}`);
}

// Получить список кандидатов голосования
async function getCandidates(id) {
  return apiFetch(`/api/proposals/${id}/candidates`);
}

// Получить результаты голосования
async function getResults(id) {
  return apiFetch(`/api/proposals/${id}/results`);
}

// Получить текущую фазу голосования
async function getPhase(id) {
  return apiFetch(`/api/proposals/${id}/phase`);
}

// Получить список адресов сделавших commit
async function getVoters(id) {
  return apiFetch(`/api/proposals/${id}/voters`);
}

// ─── Участники ────────────────────────────────────────────────────────────

// Проверить статус регистрации адреса
async function getVoterStatus(addr) {
  return apiFetch(`/api/voters/${addr}/status`);
}

// Проверить голосовал ли адрес в голосовании
async function checkVoted(id, addr) {
  return apiFetch(`/api/proposals/${id}/votes/${addr}`);
}

// ─── Верификация ──────────────────────────────────────────────────────────

// Независимая верификация результатов (Stage B)
async function verifyProposal(id) {
  const res = await fetch(`/api/proposals/${id}/verify`);
  // 200 = valid, 409 = расхождения — оба варианта парсим как JSON
  return res.json();
}

// ─── Утилиты ──────────────────────────────────────────────────────────────

// Сгенерировать commit hash для кандидата
// Возвращает { commit_hash, salt, candidate_id }
async function generateCommitHash(candidateId) {
  return apiFetch('/api/tools/commit-hash', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ candidate_id: candidateId }),
  });
}

// ─── Форматирование ───────────────────────────────────────────────────────

// Метка фазы для отображения
function phaseLabel(phase) {
  const labels = { commit: 'Commit', reveal: 'Reveal', finalized: 'Завершено' };
  return labels[phase] || phase;
}

// CSS-класс бейджа фазы
function phaseBadgeClass(phase) {
  const classes = {
    commit:    'badge badge-commit',
    reveal:    'badge badge-reveal',
    finalized: 'badge badge-finalized',
  };
  return classes[phase] || 'badge';
}

// Форматировать unix timestamp в читаемую дату
function formatDeadline(ts) {
  if (!ts) return '—';
  return new Date(ts * 1000).toLocaleString('ru-RU');
}

// Сократить адрес: 0x1234...abcd
function shortAddr(addr) {
  if (!addr) return '';
  return addr.slice(0, 6) + '...' + addr.slice(-4);
}

// Обратный отсчёт до timestamp
function countdown(ts) {
  const diff = ts - Math.floor(Date.now() / 1000);
  if (diff <= 0) return 'Истёк';
  const h = Math.floor(diff / 3600);
  const m = Math.floor((diff % 3600) / 60);
  const s = diff % 60;
  if (h > 0) return `${h}ч ${m}м`;
  if (m > 0) return `${m}м ${s}с`;
  return `${s}с`;
}