package shared

import "net/http"

func Error(w http.ResponseWriter, msg string, code int) {
	http.Error(w, msg, code)
}
