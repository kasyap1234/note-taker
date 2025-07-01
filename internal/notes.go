package internal

import (
	"context"
	"errors"
	"note-taker/repository"
)

// NoteService contains the business logic for note operations.
// It's decoupled from the web layer (e.g., Echo) and focuses on core logic.
type NoteService struct {
	repo *repository.Queries
}

// NewNoteService creates a new NoteService.
func NewNoteService(repo *repository.Queries) *NoteService {
	return &NoteService{repo: repo}
}

// CreateNoteParams defines the parameters for creating a note in the service layer.
type CreateNoteParams struct {
	Title   string
	Content string
}

// CreateNote handles the business logic for creating a new note.
func (s *NoteService) CreateNote(ctx context.Context, params CreateNoteParams) (repository.Note, error) {
	if params.Title == "" {
		return repository.Note{}, errors.New("title cannot be empty")
	}
	if params.Content == "" {
		return repository.Note{}, errors.New("content cannot be empty")
	}

	dbParams := repository.CreateNoteParams{
		Title:   params.Title,
		Content: params.Content,
	}

	return s.repo.CreateNote(ctx, dbParams)
}
