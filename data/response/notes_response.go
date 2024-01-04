package response

import "time"

type NoteResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedByID int      `json:"created_by_id"`
}

type SharedNoteResponse struct {
	ID        int       `json:"id"`
	NoteID    int       `json:"note_id"`
	UserID    int       `json:"user_id"`
	SharedByID int      `json:"shared_by_id"`
	CreatedAt time.Time `json:"created_at"`
}
