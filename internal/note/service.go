package note

type Service interface {
	CreateNote(input CreateNoteInput) (Note, error)
	GetNotes() ([]Note, error)
	GetNoteById(id string) (Note, error)
	UpdateNote(id string, input UpdateNoteInput) (Note, error)
	DeleteNote(id string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateNote(input CreateNoteInput) (Note, error) {
	return Note{}, nil
}

func (s *service) GetNotes() ([]Note, error) {
	return []Note{}, nil
}

func (s *service) GetNoteById(id string) (Note, error) {
	return Note{}, nil
}

func (s *service) UpdateNote(id string, input UpdateNoteInput) (Note, error) {
	return Note{}, nil
}

func (s *service) DeleteNote(id string) error {
	return nil
}
