package note

import (
	"context"

	"github.com/o-ga09/note-app-backendapi/domain"
)

type NoteService struct {
	noteRepo domain.INoteRepository
}

func NewNoteService(noteRepo domain.INoteRepository) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (s *NoteService) FetchNoteById(ctx context.Context, id string) (*domain.Note, error) {
	res, err := s.noteRepo.GetNoteById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *NoteService) FetchNotes(ctx context.Context) ([]*domain.Note, error) {
	res, err := s.noteRepo.GetNotes(ctx)
	if err != nil {
		return nil, err
	}
	notes := make([]*domain.Note, 0)
	for _, note := range res {
		note := domain.Note{
			NoteID:    note.NoteID,
			Title:     note.Title,
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		}
		notes = append(notes, &note)
	}
	return notes, nil
}

func (s *NoteService) CreateNote(ctx context.Context, title, content string) error {
	note := domain.Note{
		Title:   title,
		Content: content,
	}
	err := s.noteRepo.CreateNote(ctx, note)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) UpdateNote(ctx context.Context, id, title, content string) error {
	note := domain.Note{
		NoteID:  id,
		Title:   title,
		Content: content,
	}
	err := s.noteRepo.UpdateNote(ctx, note)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) DeleteNoteById(ctx context.Context, id string) error {
	err := s.noteRepo.DeleteNoteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
