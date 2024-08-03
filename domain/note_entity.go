package domain

type Note struct {
	NoteID    string `json:"note_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
}
