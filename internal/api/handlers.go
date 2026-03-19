package api

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
	"voting-dapp/internal/blockchain"

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

	proposal, err := h.client.GetProposal(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if proposal.Id.Cmp(big.NewInt(0)) == 0 {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "proposal not found"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"id":          proposal.Id.Uint64(),
		"title":       proposal.Title,
		"description": proposal.Description,
		"creator":     proposal.Creator.Hex(),
		"start_time":  proposal.StartTime.Uint64(),
		"end_time":    proposal.EndTime.Uint64(),
		"finalized":   proposal.Finalized,
		"total_votes": proposal.TotalVotes.Uint64(),
	})
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
func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
		return
	}

	// Десериализуем тело запроса
	var req VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Валидация полей
	if req.PrivateKey == "" {
		log.Printf("request: %+v", req)
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Private key is required"})
		return
	}
	if req.CandidateID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Candidate ID is required"})
		return
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.client.Vote(ctx, req.PrivateKey, id, req.CandidateID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"tx_hash": txHash,
		"status":  "submitted",
	})
}

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

