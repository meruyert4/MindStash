package note

import (
	"context"
	"database/sql"
	"mindstash/pkg/logger"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, input CreateNoteInput) (Note, error)
	GetAll(ctx context.Context) ([]Note, error)
	GetById(ctx context.Context, id string) (Note, error)
	Update(ctx context.Context, id string, input UpdateNoteInput) (Note, error)
	Delete(ctx context.Context, id string) error
}

type repo struct {
	db     *sql.DB
	logger logger.Logger
}

func NewRepository(db *sql.DB, log logger.Logger) Repository {
	return &repo{db: db, logger: log}
}

func (r *repo) Create(ctx context.Context, input CreateNoteInput) (Note, error) {
	id := uuid.New().String()
	now := time.Now().Format(time.RFC3339)

	query := `
		INSERT INTO notes (id, title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	r.logger.Debug(ctx, "Inserting note into notes", query, "id", id)

	_, err := r.db.ExecContext(ctx, query, id, input.Title, input.Content, now, now)
	if err != nil {
		r.logger.Error(ctx, "failed to insert note", "error", err, "id", id)
		return Note{}, err
	}

	r.logger.Info(ctx, "note inserted succeddfully", "id", id)
	return Note{
		Id:        id,
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (r *repo) GetAll(ctx context.Context) ([]Note, error) {
	query := `
	SELECT id, title, content, created_at, updated_at 
	FROM notes
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		r.logger.Error(ctx, "failed to fetch notes", "getall")
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.Id, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			r.logger.Error(ctx, "failed to scan notes through all", "error", err)
			return nil, err
		}
		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error(ctx, "something went wrong", "error", err)
		return nil, err
	}

	r.logger.Info(ctx, "notes successfully fetched")

	return notes, nil
}

func (r *repo) GetById(ctx context.Context, id string) (Note, error) {
	query := `
	SELECT id, title, content, created_at, updated_at FROM notes WHERE id=$1 
	`

	var note Note
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		r.logger.Error(ctx, "failed to scan note", "error", err, "id", id)
		if err == sql.ErrNoRows {
			r.logger.Error(ctx, "not found", "error", err)
			return Note{}, err
		}
		return Note{}, err
	}

	r.logger.Info(ctx, "note successfully fetched", "id", id)
	return note, nil
}

func (r *repo) Update(ctx context.Context, id string, input UpdateNoteInput) (Note, error) {
	now := time.Now().Format(time.RFC3339)

	query := `
		UPDATE notes SET title = $1, content = $2, updated_at=$3 WHERE id=$4
	`

	existing, err := r.GetById(ctx, id)
	if err != nil {
		r.logger.Error(ctx, "failed to fetch note", "error", err, "id", id)
		return Note{}, err
	}
	existing.Title = *input.Title
	existing.Content = *input.Content
	existing.UpdatedAt = now
	_, err = r.db.ExecContext(ctx, query, existing.Title, existing.Content, existing.UpdatedAt, existing.Id)
	if err != nil {
		r.logger.Error(ctx, "failed to update note", "failed to execute query", "id", id, "error", err)
		return Note{}, err
	}
	r.logger.Info(ctx, "note successfully updated", "id", id)
	return existing, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM notes WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error(ctx, "failed to delete note", "failed to execute query", "error", err, "id", id)
		return err
	}

	r.logger.Info(ctx, "note successfully deleted")
	return nil
}
