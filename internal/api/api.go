package api

import (
	"net/http"
	"welp/internal/api/handlers"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Api)
	return mux
}
