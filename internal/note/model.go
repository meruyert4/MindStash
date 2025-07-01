package note

type Note struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Color     string
}

type CreateNoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteInput struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}
