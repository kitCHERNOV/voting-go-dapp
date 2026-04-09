package api

import (
	"voting-dapp/internal/blockchain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter - регистрирует все маршруты и возвращает настроенный роутер
func SetupRouter(client *blockchain.Client) chi.Router {
	r := chi.NewRouter()

	// Middleware для логирования и восстановления после паники
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	h := NewHandler(client)

	// Health check endpoint
	r.Get("/health", h.Health)

	r.Route("/api/proposals", func(r chi.Router) {
		// Список всех голосований
		r.Get("/", h.GetAllProposals)

		// Детали голосования
		r.Get("/{id}", h.GetProposal)

		// Список кандидатов
		r.Get("/{id}/candidates", h.GetCandidates)

		// Результаты голосования
		r.Get("/{id}/results", h.GetResults)

		// Проверка голосовал ли адрес
		r.Get("/{id}/votes/{addr}", h.CheckVoted)

		// Голосование
		// r.Post("/{id}/vote", h.Vote)

		// Финализация
		r.Post("/{id}/finalize", h.FinalizeProposal)

		// Stage 2: Commit-Reveal
		r.Post("/{id}/commit", h.Commit)
		r.Post("/{id}/reveal", h.Reveal)
		r.Get("/{id}/phase", h.GetPhase)
		r.Post("/{id}/advance-phase", h.AdvancePhase)
	})

	// Утилиты
	r.Post("/api/tools/commit-hash", h.GenerateCommitHash)

	return r
}