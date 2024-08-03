package dao

import (
	"context"
	"database/sql"

	"github.com/o-ga09/note-app-backendapi/db/db"
	"github.com/o-ga09/note-app-backendapi/domain"
)

type noteDao struct {
	query *db.Queries
}

func NewNoteDao(d *sql.DB) *noteDao {
	q := db.New(d)
	return &noteDao{query: q}
}

func (dao *noteDao) GetNoteById(ctx context.Context, id string) (*domain.Note, error) {
	res, err := dao.query.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	note := domain.Note{
		NoteID:    res.NoteID,
		Title:     res.Title,
		Content:   res.Content,
		CreatedAt: res.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: res.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
	return &note, nil
}

func (dao *noteDao) GetNotes(ctx context.Context) ([]domain.Note, error) {
	res, err := dao.query.GetNotes(ctx, "")
	if err != nil {
		return nil, err
	}
	notes := make([]domain.Note, len(res))
	for _, r := range res {
		note := domain.Note{
			NoteID:    r.NoteID,
			Title:     r.Title,
			Content:   r.Content,
			CreatedAt: r.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt: r.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (dao *noteDao) CreateNote(ctx context.Context, note domain.Note) error {
	_, err := dao.query.CreateNote(ctx, db.CreateNoteParams{
		NoteID:  note.NoteID,
		Title:   note.Title,
		Content: note.Content,
	})
	if err != nil {
		return err
	}

	return nil
}

func (dao *noteDao) UpdateNote(ctx context.Context, note domain.Note) error {
	err := dao.query.UpdateNote(ctx, db.UpdateNoteParams{
		NoteID:  note.NoteID,
		Title:   note.Title,
		Content: note.Content,
	})
	if err != nil {
		return err
	}

	return nil
}

func (dao *noteDao) DeleteNoteById(ctx context.Context, id string) error {
	err := dao.query.DeleteNote(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
