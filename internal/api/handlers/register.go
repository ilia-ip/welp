package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"welp/internal/shared"
	"welp/internal/shared/data"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(users_db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			shared.Error(w, "Invalid body", http.StatusBadRequest)
		}

		var req data.RegisterRequest
		if err := json.Unmarshal(body, &req); err != nil {
			shared.Error(w, "Invalid credentials", http.StatusBadRequest)
		}

		h_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			shared.Error(w, "Internal error", http.StatusInternalServerError)
		}

		users_db.Create(&data.User{
			Username: req.Username,
			Password: h_password,
		})
	}
}
