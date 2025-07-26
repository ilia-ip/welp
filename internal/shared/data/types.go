package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string
	Password []byte
}

type Msg struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Replied  bool
	RepledTo uint // original message id

	AuthorID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Content  string    // will use markdown
}

type Post struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	Tags []string `gorm:"type:text[]"`

	AuthorID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title    string
	Content  string // will use markdown
}
