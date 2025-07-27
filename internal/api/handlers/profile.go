package handlers

import (
	"encoding/json"
	"net/http"
	"welp/internal/api/middleware"
	"welp/internal/shared"
	"welp/internal/shared/data"

	"gorm.io/gorm"
)

func Profile(users_db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.USER_ID_KEY).(string)
		if !ok {
			shared.Error(w, "Internal Error", http.StatusInternalServerError)
		}

		var user data.User
		users_db.First(&user, "username = ?", username)

		json, err := json.Marshal(user)
		if err != nil {
			shared.Error(w, "Internal error", http.StatusInternalServerError)
		}

		w.Write(json)
	}
}
