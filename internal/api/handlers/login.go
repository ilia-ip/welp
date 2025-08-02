package handlers

import (
	"io"
	"net/http"
	"welp/internal/api/model/auth"
	"welp/internal/api/model/werror"
)

func Login(auth auth.AuthorizationMethod) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			Error(w, werror.BAD_REQ)
		}

		token, ok := auth.CreateUser(string(body))
		if !ok {
			Error(w, werror.INT_ERR)
		}

		w.Write(token)
	}
}
