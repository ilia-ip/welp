package handlers

import "net/http"

func Api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welpi"))
}
