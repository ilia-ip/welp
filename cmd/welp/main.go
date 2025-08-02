package main

import (
	"log/slog"
	"net/http"
	"os"
	"welp/internal/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", api.Register()))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	slog.Info("Running server at :8080")
	server.ListenAndServe()
}
