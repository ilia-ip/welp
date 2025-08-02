package api

import (
	"net/http"
)

func Register() http.Handler {

	mux := http.NewServeMux()

	return mux
}
