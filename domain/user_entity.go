package domain

import "github.com/google/uuid"

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	UserEmail string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created"`
	UpdatedAt string    `json:"updated"`
}
