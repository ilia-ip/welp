package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"welp/internal/shared"
	"welp/internal/shared/data"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(users_db *gorm.DB, token_gen func(uuid.UUID) (string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			shared.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		var req data.RegisterRequest
		if err := json.Unmarshal(body, &req); err != nil {
			shared.Error(w, "Invalid credentials", http.StatusBadRequest)
			return
		}

		h_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			shared.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		uuid := uuid.New()

		users_db.Create(&data.User{
			Username: req.Username,
			Password: h_password,
			ID:       uuid,
		})

		token, err := token_gen(uuid)
		if err != nil {
			shared.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(token))
	}
}
