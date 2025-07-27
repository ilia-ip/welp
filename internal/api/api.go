package api

import (
	"net/http"
	"welp/internal/api/handlers"
	"welp/internal/api/middleware"
	"welp/internal/shared/auth"

	"gorm.io/gorm"
)

func Init(users_db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handlers.Register(users_db))
	// mux.HandleFunc("POST /login", nil)

	authed_mux := http.NewServeMux()
	authed_mux.HandleFunc("GET /profile", handlers.Profile(users_db))
	// mux.HandleFunc("GET /users/{user_id}/profile", nil)

	// mux.HandleFunc("GET /posts", nil)
	// mux.HandleFunc("GET /posts/{post_id}", nil)
	// mux.HandleFunc("POST /posts", nil)
	// mux.HandleFunc("PUT /posts/{post_id}", nil)
	// mux.HandleFunc("DELETE /posts/{post_id}", nil)

	// mux.HandleFunc("GET /users/{user_id}/posts", nil)

	// mux.HandleFunc("GET /chats", nil)
	// mux.HandleFunc("GET /chats/{user_id}", nil)
	// mux.HandleFunc("POST /chats/{user_id}", nil)
	// mux.HandleFunc("PUT /messages/{message_id}", nil)
	// mux.HandleFunc("DELETE /messages/{message_id}", nil)

	mux.Handle("/", middleware.Auth(auth.ParseJWT, authed_mux))

	return mux
}
