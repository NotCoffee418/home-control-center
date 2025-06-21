package main

import (
	"log"
	"net/http"
	"os"

	root "github.com/NotCoffee418/home-control-center"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get the frontend from root package
	frontendFS, err := root.GetFrontendFS()
	if err != nil {
		log.Fatal("Failed to get frontend:", err)
	}

	// Serve React app
	http.Handle("/", http.FileServer(http.FS(frontendFS)))

	// API routes
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Bind to localhost in development, all interfaces in production
	var addr string
	if os.Getenv("GO_ENV") == "production" {
		addr = ":" + port
		log.Printf("Server starting on all interfaces, port %s", port)
	} else {
		addr = "127.0.0.1:" + port
		log.Printf("Server starting on localhost:%s (development mode)", port)
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}
