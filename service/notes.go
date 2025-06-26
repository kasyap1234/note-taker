package service

import (
	"context"
	"errors"
	"note-taker/repository"
)

type NoteService struct {
	repo *repository.Queries
}

func NewNoteService(repo *repository.Queries) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

type CreateNoteParams struct {
	Title   string
	Content string
}

func (s *NoteService) CreateNote(ctx context.Context, params CreateNoteParams) (repository.Note, error) {
	if params.Title == "" {
		return repository.Note{}, errors.New("title cannot be empty")
	}
	if params.Content == "" {
		return repository.Note{}, errors.New("content cannot be empty")
	}
	dbParams := repository.CreateNoteParams{
		Title:   params.Title,
		Content: params.Title,
	}
	return s.repo.CreateNote(ctx, dbParams)
}


