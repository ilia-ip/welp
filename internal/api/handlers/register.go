package handlers

import (
	"encoding/json"
	"net/http"
	"welp/internal/shared"
	"welp/internal/shared/data"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var req data.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	_, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		shared.Error(w, "Internal error", http.StatusInternalServerError)
	}

	// .Create(&data.User{
	// Username: req.Username,
	// Password: h_password,
	// })
}
