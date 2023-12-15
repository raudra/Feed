package models

import "github.com/google/uuid"

type Feed struct {
	uuid        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CreatedAt   int       `json:"created_at"`
}
