// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateNote implements createNote operation.
//
// Create a note.
//
// POST /notes
func (UnimplementedHandler) CreateNote(ctx context.Context, req *Note) (r CreateNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// Create a user.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *User) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteNote implements deleteNote operation.
//
// Delete a note.
//
// DELETE /note/{noteId}
func (UnimplementedHandler) DeleteNote(ctx context.Context, params DeleteNoteParams) (r DeleteNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteUser implements deleteUser operation.
//
// Delete a user.
//
// DELETE /user/{userId}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNote implements getNote operation.
//
// Get a note.
//
// GET /note/{noteId}
func (UnimplementedHandler) GetNote(ctx context.Context, params GetNoteParams) (r GetNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNotes implements getNotes operation.
//
// Get all notes.
//
// GET /notes
func (UnimplementedHandler) GetNotes(ctx context.Context) (r GetNotesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUser implements getUser operation.
//
// Get a user.
//
// GET /user/{userId}
func (UnimplementedHandler) GetUser(ctx context.Context, params GetUserParams) (r GetUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUsers implements getUsers operation.
//
// Get all users.
//
// GET /users
func (UnimplementedHandler) GetUsers(ctx context.Context) (r GetUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateNote implements updateNote operation.
//
// Update a note.
//
// PUT /note/{noteId}
func (UnimplementedHandler) UpdateNote(ctx context.Context, req *Note, params UpdateNoteParams) (r UpdateNoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// Update a user.
//
// PUT /user/{userId}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UpdateUser, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *CommonErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *CommonErrorStatusCode) {
	r = new(CommonErrorStatusCode)
	return r
}
