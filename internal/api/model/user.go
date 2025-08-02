package model

import (
	"crypto/rsa"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Password []byte

	UserName    string
	Description string
	Status      string

	// For encrypting messages
	Key rsa.PrivateKey
}
