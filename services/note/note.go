package note

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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
	notes := []*domain.Note{}
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
	fmt.Println(len(notes))
	return notes, nil
}

func (s *NoteService) CreateNote(ctx context.Context, title, content string) (domain.Note, error) {
	noteId := uuid.New()
	note := domain.Note{
		NoteID:  noteId,
		Title:   title,
		Content: content,
	}
	err := s.noteRepo.CreateNote(ctx, note)
	if err != nil {
		return domain.Note{}, err
	}
	createdNote, err := s.noteRepo.GetNoteById(ctx, noteId.String())
	if err != nil {
		return domain.Note{}, err
	}
	return *createdNote, nil
}

func (s *NoteService) UpdateNote(ctx context.Context, id, title, content string) (domain.Note, error) {
	noteId, err := uuid.Parse(id)
	if err != nil {
		return domain.Note{}, err
	}
	note := domain.Note{
		NoteID:  noteId,
		Title:   title,
		Content: content,
	}
	err = s.noteRepo.UpdateNote(ctx, note)
	if err != nil {
		return domain.Note{}, err
	}
	updatedNote, err := s.noteRepo.GetNoteById(ctx, noteId.String())
	if err != nil {
		return domain.Note{}, err
	}
	return *updatedNote, nil
}

func (s *NoteService) DeleteNoteById(ctx context.Context, id string) error {
	err := s.noteRepo.DeleteNoteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
