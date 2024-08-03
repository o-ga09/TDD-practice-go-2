// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createNote = `-- name: CreateNote :execresult
INSERT INTO notes (
    note_id,
    title,
    content
) VALUES (?, ?, ?)
`

type CreateNoteParams struct {
	NoteID  string
	Title   string
	Content string
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createNote, arg.NoteID, arg.Title, arg.Content)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    user_id,
    name,
    email,
    password
) VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes
WHERE note_id = ?
`

func (q *Queries) DeleteNote(ctx context.Context, noteID string) error {
	_, err := q.db.ExecContext(ctx, deleteNote, noteID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getNote = `-- name: GetNote :one
SELECT id, note_id, title, content, created_at, updated_at FROM notes
WHERE note_id = ? LIMIT 1
`

func (q *Queries) GetNote(ctx context.Context, noteID string) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNote, noteID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNotes = `-- name: GetNotes :many
SELECT id, note_id, title, content, created_at, updated_at FROM notes
WHERE note_id = ?
ORDER BY created_at DESC
`

func (q *Queries) GetNotes(ctx context.Context, noteID string) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getNotes, noteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.NoteID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, user_id, name, email, password, created_at, updated_at FROM users
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, user_id, name, email, password, created_at, updated_at FROM users
ORDER BY created_at DESC
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNote = `-- name: UpdateNote :exec
UPDATE notes
SET title = ?,
    content = ?
WHERE note_id = ?
`

type UpdateNoteParams struct {
	Title   string
	Content string
	NoteID  string
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) error {
	_, err := q.db.ExecContext(ctx, updateNote, arg.Title, arg.Content, arg.NoteID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = ?,
    email = ?,
    password = ?
WHERE user_id = ?
`

type UpdateUserParams struct {
	Name     string
	Email    string
	Password string
	UserID   string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.UserID,
	)
	return err
}
