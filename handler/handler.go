package handler

import (
	"context"

	"github.com/o-ga09/note-app-backendapi/api"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) CreateNote(ctx context.Context, req *api.Note) (api.CreateNoteRes, error) {
	return &api.Note{}, nil
}
func (h *handler) DeleteNote(ctx context.Context, params api.DeleteNoteParams) (api.DeleteNoteRes, error) {
	return &api.DeleteNote{}, nil
}
func (h *handler) GetNote(ctx context.Context, params api.GetNoteParams) (api.GetNoteRes, error) {
	return &api.Note{}, nil
}
func (h *handler) GetNotes(ctx context.Context) (api.GetNotesRes, error) {
	return &api.Notes{}, nil
}
func (h *handler) UpdateNote(ctx context.Context, req *api.Note, params api.UpdateNoteParams) (api.UpdateNoteRes, error) {
	return &api.UpdateNote{}, nil
}

func (h *handler) CreateUser(ctx context.Context, req *api.User) (api.CreateUserRes, error) {
	return &api.User{}, nil
}
func (h *handler) DeleteUser(ctx context.Context, params api.DeleteUserParams) (api.DeleteUserRes, error) {
	return &api.DeleteUser{}, nil
}
func (h *handler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {
	return &api.User{}, nil
}
func (h *handler) GetUsers(ctx context.Context) (api.GetUsersRes, error) {
	return &api.Users{}, nil
}
func (h *handler) UpdateUser(ctx context.Context, req *api.UpdateUser, params api.UpdateUserParams) (api.UpdateUserRes, error) {
	return &api.User{}, nil
}

func (h *handler) NewError(ctx context.Context, err error) *api.CommonErrorStatusCode {
	return &api.CommonErrorStatusCode{}
}
