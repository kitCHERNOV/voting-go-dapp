package api

import (
	"voting-dapp/internal/blockchain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter регистрирует все маршруты и возвращает настроенный роутер.
func SetupRouter(client *blockchain.Client, registry *blockchain.RegistryClient) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewHandler(client, registry)

	// Health check
	r.Get("/health", h.Health)

	r.Route("/api/proposals", func(r chi.Router) {
		r.Get("/", h.GetAllProposals)
		r.Get("/{id}", h.GetProposal)
		r.Get("/{id}/candidates", h.GetCandidates)
		r.Get("/{id}/results", h.GetResults)
		r.Get("/{id}/votes/{addr}", h.CheckVoted)
		r.Post("/{id}/finalize", h.FinalizeProposal)

		// Stage 2: Commit-Reveal
		r.Post("/{id}/commit", h.Commit)
		r.Post("/{id}/reveal", h.Reveal)
		r.Get("/{id}/phase", h.GetPhase)
		r.Post("/{id}/advance-phase", h.AdvancePhase)

		// Stage 3: список адресов сделавших commit
		r.Get("/{id}/voters", h.GetProposalVoters)
	})

	// Утилиты
	r.Post("/api/tools/commit-hash", h.GenerateCommitHash)

	// Stage 3: Voter Registry
	r.Get("/api/voters/count", h.GetVoterCount)
	r.Get("/api/voters/{addr}/status", h.GetVoterStatus)

	r.Route("/api/admin/voters", func(r chi.Router) {
		r.Post("/register", h.RegisterVoter)
		r.Post("/register-batch", h.RegisterBatch)
		r.Delete("/{addr}", h.RevokeVoter)
	})

	return r
}