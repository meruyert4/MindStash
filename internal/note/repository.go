package note

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, input CreateNoteInput) (Note, error)
	GetAll(ctx context.Context) ([]Note, error)
	GetById(ctx context.Context, id string) (Note, error)
	Update(ctx context.Context, id string, input UpdateNoteInput) (Note, error)
	Delete(ctx context.Context, id string) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, input CreateNoteInput) (Note, error) {
	return Note{}, nil
}

func (r *repo) GetAll(ctx context.Context) ([]Note, error) {
	return []Note{}, nil
}

func (r *repo) GetById(ctx context.Context, id string) (Note, error) {
	return Note{}, nil
}

func (r *repo) Update(ctx context.Context, id string, input UpdateNoteInput) (Note, error) {
	return Note{}, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	return nil
}
