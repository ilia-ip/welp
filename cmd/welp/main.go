package main

import (
	"log/slog"
	"net/http"
	"os"
	"welp/internal/api"
	"welp/internal/frontend"
)

func main() {
	mux := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux.Handle("/api/", http.StripPrefix("/api", api.Init()))

	mux.Handle("/", frontend.Init())

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	logger.Info("Started server at :8080")
	server.ListenAndServe()
}
