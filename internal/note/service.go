package note

import (
	"context"
	"errors"
	"mindstash/pkg/logger"
)

type Service interface {
	CreateNote(ctx context.Context, input CreateNoteInput) (Note, error)
	GetNotes(ctx context.Context) ([]Note, error)
	GetNoteById(ctx context.Context, id string) (Note, error)
	UpdateNote(ctx context.Context, id string, input UpdateNoteInput) (Note, error)
	DeleteNote(ctx context.Context, id string) error
}

type service struct {
	repo   Repository
	logger logger.Logger
}

func NewService(r Repository, log logger.Logger) Service {
	return &service{repo: r, logger: log}
}

func (s *service) CreateNote(ctx context.Context, input CreateNoteInput) (Note, error) {
	if err := validateNoteInput(input.Title, input.Content); err != nil {
		s.logger.Warn(ctx, "note validation failed", "error", err)
		return Note{}, err
	}

	sanitizedTitle, sanitizedContent := sanitizeNoteInput(input.Title, input.Content)
	input.Title = sanitizedTitle
	input.Content = sanitizedContent

	note, err := s.repo.Create(ctx, input)
	if err != nil {
		s.logger.Error(ctx, "service: note creation failed", "error", err)
		return Note{}, err
	}

	s.logger.Info(ctx, "service: note created", "id", note.Id)
	return note, nil
}

func (s *service) GetNotes(ctx context.Context) ([]Note, error) {
	notes, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error(ctx, "service: failed to get notes")
		return []Note{}, err
	}
	s.logger.Info(ctx, "service: notes successfully fetched")
	return notes, nil
}

func (s *service) GetNoteById(ctx context.Context, id string) (Note, error) {
	note, err := s.repo.GetById(ctx, id)
	if err != nil {
		s.logger.Error(ctx, "service: failed to get note", "error", err)
		return Note{}, err
	}

	s.logger.Info(ctx, "service: note successfully fetched", "id", note.Id)
	return note, nil
}

func (s *service) UpdateNote(ctx context.Context, id string, input UpdateNoteInput) (Note, error) {
	if err := validateNoteInput(*input.Title, *input.Content); err != nil {
		return Note{}, err
	}
	note, err := s.repo.Update(ctx, id, input)
	if err != nil {
		s.logger.Error(ctx, "service: failed to update note")
		return Note{}, err
	}
	if isEmptyUpdate(input) {
		s.logger.Warn(ctx, "service: empty update input")
		return Note{}, errors.New("no fields to update")
	}
	s.logger.Info(ctx, "service: note updated successsfully", "id", note.Id)
	return note, nil
}

func (s *service) DeleteNote(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.Error(ctx, "service: failed to delete note", "error", err)
		return err
	}
	s.logger.Info(ctx, "service: note successfully deleted", "id", id)
	return nil
}
