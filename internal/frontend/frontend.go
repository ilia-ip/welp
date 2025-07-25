package frontend

import (
	"net/http"
	"welp/internal/frontend/pages"

	"github.com/a-h/templ"
)

func Register(mux *http.ServeMux) {
	mux.Handle("GET /", templ.Handler(pages.Home()))
}
