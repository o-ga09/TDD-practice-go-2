package domain

import "context"

type INoteRepository interface {
	GetNoteById(ctx context.Context, id string) (*Note, error)
	GetNotes(ctx context.Context) ([]Note, error)
	CreateNote(ctx context.Context, note Note) error
	UpdateNote(ctx context.Context, note Note) error
	DeleteNoteById(ctx context.Context, id string) error
}

type IUserRepository interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUserById(ctx context.Context, id string) error
}
