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

	api.Register(mux)
	frontend.Register(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	logger.Info("Started server at :8080")
	server.ListenAndServe()
}
