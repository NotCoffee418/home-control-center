package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

	// Serve static assets first
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(frontendFS))))

	// Catch-all: serve index.html for all other routes (React Router)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Try to open the requested file
		file, err := frontendFS.Open(strings.TrimPrefix(r.URL.Path, "/"))
		if err == nil {
			file.Close()
			// File exists, serve it normally
			http.FileServer(http.FS(frontendFS)).ServeHTTP(w, r)
			return
		}

		// File doesn't exist, serve index.html (React Router will handle it)
		indexFile, err := frontendFS.Open("index.html")
		if err != nil {
			http.Error(w, "Frontend not found", 404)
			return
		}
		defer indexFile.Close()

		w.Header().Set("Content-Type", "text/html")
		io.Copy(w, indexFile)
	})

	// for now bind to 9040 (needs to be changed to port from config)
	log.Printf("Server starting on all interfaces, port %s", port)
	log.Fatal(http.ListenAndServe(":9040", r))
}
