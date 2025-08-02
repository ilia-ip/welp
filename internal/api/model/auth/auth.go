package auth

import (
	"github.com/google/uuid"
)

type AuthorizationMethod interface {
	VerifyToken([]byte) (uuid.UUID, bool)

	CreateUser(string, string) ([]byte, bool)
}
