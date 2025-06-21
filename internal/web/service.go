package web

import (
	"log"
	"net/http"
	"os"

	root "github.com/NotCoffee418/home-control-center"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartWebServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get the frontend from root package
	frontendFS, err := root.GetFrontendFS()
	if err != nil {
		log.Fatal("Failed to get frontend:", err)
	}

	// Create chi router
	r := chi.NewRouter()

	// Production vs development middleware
	if os.Getenv("GO_ENV") == "production" {
		// Production: minimal middleware
		r.Use(middleware.Recoverer)
		r.Use(middleware.Compress(5)) // gzip compression
	} else {
		// Development: more verbose
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.RequestID)
	}

	// API routes
	setupAPIRoutes(r)

	// Serve React app - handle everything else
	r.Handle("/*", http.FileServer(http.FS(frontendFS)))

	// Bind to localhost in development, all interfaces in production
	var addr string
	if os.Getenv("GO_ENV") == "production" {
		addr = ":" + port
		log.Printf("Server starting on all interfaces, port %s", port)
	} else {
		addr = "127.0.0.1:" + port
		log.Printf("Server starting on localhost:%s (development mode)", port)
	}

	log.Fatal(http.ListenAndServe(addr, r))
}
