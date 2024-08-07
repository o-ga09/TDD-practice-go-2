package domain

import "github.com/google/uuid"

type Note struct {
	NoteID    uuid.UUID `json:"note_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created"`
	UpdatedAt string    `json:"updated"`
}
