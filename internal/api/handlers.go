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

// Handler содержит зависимости для всех HTTP-хендлеров.
type Handler struct {
	client   *blockchain.Client
	registry *blockchain.RegistryClient
}

// NewHandler создаёт Handler с клиентами Voting и VoterRegistry.
func NewHandler(client *blockchain.Client, registry *blockchain.RegistryClient) *Handler {
	return &Handler{client: client, registry: registry}
}

// ctxWithTimeout создаёт контекст с таймаутом 5 секунд.
func ctxWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// parseProposalID парсит :id из URL.
func parseProposalID(r *http.Request) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
}

// writeJSON записывает JSON-ответ.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// ─── Health ───────────────────────────────────────────────────────────────────

// Health — GET /health
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// ─── Proposals ────────────────────────────────────────────────────────────────

// GetProposal — GET /api/proposals/{id}
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

// GetResults — GET /api/proposals/{id}/results
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
		results[i] = map[string]uint64{"candidate_id": ids[i], "vote_count": votes[i]}
	}
	writeJSON(w, http.StatusOK, map[string]any{"proposal_id": id, "results": results})
}

// CheckVoted — GET /api/proposals/{id}/votes/{addr}
func (h *Handler) CheckVoted(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid proposal ID"})
		return
	}
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
	writeJSON(w, http.StatusOK, map[string]any{"proposal_id": id, "address": addrStr, "voted": voted})
}

// GetAllProposals — GET /api/proposals
func (h *Handler) GetAllProposals(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	proposals, err := h.client.GetAllProposals(ctx)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"count": len(proposals), "proposals": proposals})
}

// GetCandidates — GET /api/proposals/{id}/candidates
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
	writeJSON(w, http.StatusOK, map[string]any{"proposal_id": id, "candidates": candidates})
}

// FinalizeRequest — тело запроса для финализации.
type FinalizeRequest struct {
	PrivateKey string `json:"private_key"`
}

// FinalizeProposal — POST /api/proposals/{id}/finalize
func (h *Handler) FinalizeProposal(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}
	var req FinalizeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.PrivateKey == "" {
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
	writeJSON(w, http.StatusOK, map[string]any{"tx_hash": txHash, "status": "finalized"})
}

// CommitRequest — тело запроса для фазы commit.
type CommitRequest struct {
	PrivateKey string `json:"private_key"`
	CommitHash string `json:"commit_hash"`
	DepositWei string `json:"deposit_wei"`
}

// Commit — POST /api/proposals/{id}/commit
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
	hashBytes, err := hexToBytes32(req.CommitHash)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid commit_hash format"})
		return
	}
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
	writeJSON(w, http.StatusOK, map[string]string{"tx_hash": txHash, "status": "committed"})
}

// RevealRequest — тело запроса для фазы reveal.
type RevealRequest struct {
	PrivateKey  string `json:"private_key"`
	CandidateID uint64 `json:"candidate_id"`
	Salt        string `json:"salt"`
}

// Reveal — POST /api/proposals/{id}/reveal
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
	writeJSON(w, http.StatusOK, map[string]string{"tx_hash": txHash, "status": "revealed"})
}

// GetPhase — GET /api/proposals/{id}/phase
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
	writeJSON(w, http.StatusOK, map[string]any{"proposal_id": id, "phase": info["phase"]})
}

// GenerateCommitHash — POST /api/tools/commit-hash
func (h *Handler) GenerateCommitHash(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CandidateID uint64 `json:"candidate_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.CandidateID == 0 {
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

// AdvancePhaseRequest — тело запроса для смены фазы.
type AdvancePhaseRequest struct {
	PrivateKey string `json:"private_key"`
}

// AdvancePhase — POST /api/proposals/{id}/advance-phase
func (h *Handler) AdvancePhase(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}
	
	var req AdvancePhaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.PrivateKey == "" {
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
	writeJSON(w, http.StatusOK, map[string]string{"tx_hash": txHash, "status": "phase advanced"})
}

// ─── Stage 3: Voter Registry ──────────────────────────────────────────────────

// GetVoterStatus — GET /api/voters/{addr}/status
func (h *Handler) GetVoterStatus(w http.ResponseWriter, r *http.Request) {
	addrStr := chi.URLParam(r, "addr")
	if !common.IsHexAddress(addrStr) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid Ethereum address"})
		return
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	registered, err := h.registry.IsRegistered(ctx, common.HexToAddress(addrStr))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"address":    addrStr,
		"registered": registered,
	})
}

// GetVoterCount — GET /api/voters/count
func (h *Handler) GetVoterCount(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	count, err := h.registry.GetRegisteredCount(ctx)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"count": count})
}

// RegisterVoterRequest — тело запроса для регистрации участника.
type RegisterVoterRequest struct {
	PrivateKey   string `json:"private_key"`
	VoterAddress string `json:"voter_address"`
}

// RegisterVoter — POST /api/admin/voters/register
func (h *Handler) RegisterVoter(w http.ResponseWriter, r *http.Request) {
	var req RegisterVoterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}
	if !common.IsHexAddress(req.VoterAddress) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid voter_address"})
		return
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.registry.Register(ctx, req.PrivateKey, common.HexToAddress(req.VoterAddress))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"tx_hash": txHash, "status": "registered"})
}

// RegisterBatchRequest — тело запроса для массовой регистрации.
type RegisterBatchRequest struct {
	PrivateKey string   `json:"private_key"`
	Addresses  []string `json:"addresses"`
}

// RegisterBatch — POST /api/admin/voters/register-batch
func (h *Handler) RegisterBatch(w http.ResponseWriter, r *http.Request) {
	var req RegisterBatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}
	if len(req.Addresses) == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "addresses list is empty"})
		return
	}
	voters := make([]common.Address, 0, len(req.Addresses))
	for _, addr := range req.Addresses {
		if !common.IsHexAddress(addr) {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("invalid address: %s", addr)})
			return
		}
		voters = append(voters, common.HexToAddress(addr))
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	lastTx, err := h.registry.RegisterBatch(ctx, req.PrivateKey, voters)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"last_tx_hash": lastTx,
		"registered":   len(voters),
		"status":       "batch registered",
	})
}

// RevokeVoterRequest — тело запроса для отзыва регистрации.
type RevokeVoterRequest struct {
	PrivateKey string `json:"private_key"`
}

// RevokeVoter — DELETE /api/admin/voters/{addr}
func (h *Handler) RevokeVoter(w http.ResponseWriter, r *http.Request) {
	addrStr := chi.URLParam(r, "addr")
	if !common.IsHexAddress(addrStr) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid Ethereum address"})
		return
	}
	var req RevokeVoterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.PrivateKey == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "private_key is required"})
		return
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	txHash, err := h.registry.Revoke(ctx, req.PrivateKey, common.HexToAddress(addrStr))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"tx_hash": txHash, "status": "revoked"})
}

// GetProposalVoters — GET /api/proposals/{id}/voters
func (h *Handler) GetProposalVoters(w http.ResponseWriter, r *http.Request) {
	id, err := parseProposalID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid proposal id"})
		return
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	voters, err := h.client.GetProposalCommitters(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"proposal_id": id,
		"voters":      voters,
		"count":       len(voters),
	})
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

// hexToBytes32 конвертирует hex строку (с 0x или без) в [32]byte.
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