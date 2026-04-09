package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
	"voting-dapp/internal/blockchain"
	"voting-dapp/internal/voting"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	client *blockchain.Client
}

type VoteRequest struct {
	PrivateKey  string `json:"private_key"`
	CandidateID uint64 `json:"candidate_id"`
}

// NewHandler создает новый экземпляр Handler с переданным блокчейн-клиентом
func NewHandler(client *blockchain.Client) *Handler {
	return &Handler{client: client}
}

// ctxWithTimeout - вспомогательная функция для создания контекста с таймаутом 5 секунд
func ctxWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// parseProposalID - парсит :id из URL, возвращает ошибку если не число
func parseProposalID(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
}

// writeJSON - вспомогательная функция записи JSON-ответа
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// Health - проверка доступности сервера
// GET /health
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GetProposal - возвращает детали голосования ID
// GET /proposals/{id}
func (h *Handler) GetProposal(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	info, err := h.client.GetProposalInfo(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if info["id"] == uint64(0) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "proposal not found"})
		return
	}

	writeJSON(w, http.StatusOK, info)
}

// GetResults - возвращает результаты голосования: Id кандидатов и количество голосов
// GET /api/proposals/{id}/results
func (h *Handler) GetResults(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	ids, votes, err := h.client.GetResults(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	results := make([]map[string]uint64, len(ids))
	for i := range ids {
		results[i] = map[string]uint64{
			"candidate_id": ids[i],
			"vote_count":   votes[i],
		}
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"proposal_id": id,
		"results":     results,
	})
}

// CheckVoted - проверяет, голосовал ли адрес в данном голосовании
// GET /api/proposals/{id}/votes/{addr}
func (h *Handler) CheckVoted(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
		return
	}

	// валидируем Ethereum-адрес
	addrStr := chi.URLParam(r, "addr")
	if !common.IsHexAddress(addrStr) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid Ethereum address"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	voted, err := h.client.HasVoted(ctx, id, common.HexToAddress(addrStr))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"proposal_id": id,
		"address":     addrStr,
		"voted":       voted,
	})
}

// Vote - принимает голос, подписывает транзакцию и отправляет в контракт
// POST /api/proposals/{id}/vote
// func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
// 	id, err := parseProposalID(r)
// 	if err != nil {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
// 		return
// 	}

// 	// Десериализуем тело запроса
// 	var req VoteRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
// 		return
// 	}

// 	// Валидация полей
// 	if req.PrivateKey == "" {
// 		log.Printf("request: %+v", req)
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Private key is required"})
// 		return
// 	}
// 	if req.CandidateID == 0 {
// 		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Candidate ID is required"})
// 		return
// 	}

// 	ctx, cancel := ctxWithTimeout()
// 	defer cancel()

// 	txHash, err := h.client.Vote(ctx, req.PrivateKey, id, req.CandidateID)
// 	if err != nil {
// 		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, map[string]any{
// 		"tx_hash": txHash,
// 		"status":  "submitted",
// 	})
// }

// GetAllProposals — возвращает список всех голосований
// GET /api/proposals
func (h *Handler) GetAllProposals(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	proposals, err := h.client.GetAllProposals(ctx)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"count":     len(proposals),
		"proposals": proposals,
	})
}

// GetCandidates — возвращает список кандидатов голосования
// GET /api/proposals/{id}/candidates
func (h *Handler) GetCandidates(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	candidates, err := h.client.GetCandidates(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"proposal_id": id,
		"candidates":  candidates,
	})
}

// FinalizeRequest — тело запроса для финализации
type FinalizeRequest struct {
	PrivateKey string `json:"private_key"`
}

// FinalizeProposal — финализирует голосование и определяет победителя
// POST /api/proposals/{id}/finalize
func (h *Handler) FinalizeProposal(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	var req FinalizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.client.FinalizeProposal(ctx, req.PrivateKey, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"tx_hash": txHash,
		"status":  "finalized",
	})
}


// CommitRequest — тело запроса для фазы commit
type CommitRequest struct {
	PrivateKey string `json:"private_key"`
	CommitHash string `json:"commit_hash"`
	DepositWei string `json:"deposit_wei"`
}

// Commit — отправляет commit хеш с депозитом
// POST /api/proposals/{id}/commit
func (h *Handler) Commit(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	var req CommitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}
	if req.CommitHash == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "commit_hash is required"})
		return
	}
	if req.DepositWei == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "deposit_wei is required"})
		return
	}

	// Декодируем commit hash из hex строки
	hashBytes, err := hexToBytes32(req.CommitHash)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid commit_hash format"})
		return
	}

	// Парсим депозит
	deposit, ok := new(big.Int).SetString(req.DepositWei, 10)
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid deposit_wei"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.client.Commit(ctx, req.PrivateKey, id, hashBytes, deposit)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"tx_hash": txHash,
		"status":  "committed",
	})
}

// RevealRequest — тело запроса для фазы reveal
type RevealRequest struct {
	PrivateKey  string `json:"private_key"`
	CandidateID uint64 `json:"candidate_id"`
	Salt        string `json:"salt"`
}

// Reveal — раскрывает голос (candidateId + salt)
// POST /api/proposals/{id}/reveal
func (h *Handler) Reveal(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	var req RevealRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}
	if req.CandidateID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "candidate_id is required"})
		return
	}
	if req.Salt == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "salt is required"})
		return
	}

	saltBytes, err := hexToBytes32(req.Salt)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid salt format"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.client.Reveal(ctx, req.PrivateKey, id, req.CandidateID, saltBytes)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"tx_hash": txHash,
		"status":  "revealed",
	})
}

// GetPhase — возвращает текущую фазу голосования
// GET /api/proposals/{id}/phase
func (h *Handler) GetPhase(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	info, err := h.client.GetProposalInfo(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"proposal_id": id,
		"phase":       info["phase"],
	})
}

// GenerateCommitHash — утилита: генерирует CommitData (hash + salt) для фазы commit
// POST /api/tools/commit-hash
func (h *Handler) GenerateCommitHash(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CandidateID uint64 `json:"candidate_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if req.CandidateID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "candidate_id is required"})
		return
	}

	commitData, err := voting.NewCommit(new(big.Int).SetUint64(req.CandidateID))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"candidate_id": req.CandidateID,
		"commit_hash":  fmt.Sprintf("0x%x", commitData.Hash),
		"salt":         fmt.Sprintf("0x%x", commitData.Salt),
		"warning":      "Save the salt! It is required for reveal phase. Backend does NOT store it.",
	})
}


// AdvancePhaseRequest — тело запроса для смены фазы
type AdvancePhaseRequest struct {
	PrivateKey string `json:"private_key"`
}

// AdvancePhase — переводит голосование в следующую фазу
// POST /api/proposals/{id}/advance-phase
func (h *Handler) AdvancePhase(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}

	var req AdvancePhaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.client.AdvancePhase(ctx, req.PrivateKey, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"tx_hash": txHash,
		"status":  "phase advanced",
	})
}

// hexToBytes32 — конвертирует hex строку (с 0x или без) в [32]byte
func hexToBytes32(hexStr string) ([32]byte, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	var result [32]byte
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		return result, err
	}
	if len(b) > 32 {
		return result, fmt.Errorf("hex value exceeds 32 bytes")
	}
	copy(result[32-len(b):], b)
	return result, nil
}