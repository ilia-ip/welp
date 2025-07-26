package middleware

import (
	"context"
	"net/http"
	"strings"
	"welp/internal/shared"

	"github.com/google/uuid"
)

var USER_ID_KEY string = "user_id"

func Auth(auth_method func(string) (uuid.UUID, bool), next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			shared.Error(w, "Invalid auth method", http.StatusBadRequest)
		}

		user_id, exists := auth_method(strings.TrimPrefix(header, "Bearer "))
		if !exists {
			shared.Error(w, "Invalid jwt token", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), "user_id", user_id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
