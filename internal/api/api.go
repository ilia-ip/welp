package api

import (
	"net/http"
	"welp/internal/api/handlers"
)

func Register(mux *http.ServeMux) {

	mux.HandleFunc("GET /api", handlers.Api)
}
