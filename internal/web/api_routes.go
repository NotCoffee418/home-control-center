package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func setupAPIRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler)
	})
}

// Demo handler, health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
