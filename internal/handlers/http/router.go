package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/1tsandre/mini-go-backend/internal/handlers/http/health"
	"github.com/1tsandre/mini-go-backend/internal/handlers/http/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recovery)
	r.Use(middleware.Logging)

	r.Get("/health", health.HealthHandler)
	r.Get("/ready", health.ReadinessHandler)

	return r
}
