package request

type CreateNoteRequest struct {
	Title   string `validate:"required,min=2,max=255" json:"title"`
	Content string `validate:"max=65535" json:"content"`
}

type UpdateNoteRequest struct {
	Title   string `validate:"min=2,max=255" json:"title"`
	Content string `validate:"max=65535" json:"content"`
}

type ShareNoteRequest struct {
	SharedWith int `validate:"required" json:"shared_with"`
	Username string `validate:"required,min=2,max=100" json:"username"`
}