package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	ID       uuid.UUID `gorm:"type:uuid"`
	Username string
	Password []byte

	DisplayName string
	Description string
	Status      string
}

type Msg struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Replied  bool
	RepledTo uuid.UUID `gorm:"type:uuid"` // original message id

	MessageID uuid.UUID `gorm:"type:uuid"`
	AuthorID  uuid.UUID `gorm:"type:uuid"`
	Content   string    // will use markdown
}

type Post struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Tags []string `gorm:"type:text[]"`

	PostID   uuid.UUID `gorm:"type:uuid"`
	AuthorID uuid.UUID `gorm:"type:uuid"`
	Title    string
	Content  string // will use markdown
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
