package frontend

import (
	"net/http"
	"welp/internal/frontend/pages"

	"github.com/a-h/templ"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", templ.Handler(pages.Home()))
	return mux
}
