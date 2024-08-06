package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/o-ga09/note-app-backendapi/api"
	"github.com/o-ga09/note-app-backendapi/pkg/date"
	"github.com/o-ga09/note-app-backendapi/pkg/logger"
	"github.com/o-ga09/note-app-backendapi/pkg/middleware"
	"github.com/o-ga09/note-app-backendapi/services/note"
	"github.com/o-ga09/note-app-backendapi/services/user"
)

type handler struct {
	noteService note.NoteService
	userService user.UserService
}

func NewHandler(n note.NoteService, u user.UserService) *handler {
	return &handler{noteService: n, userService: u}
}

func (h *handler) CreateNote(ctx context.Context, req *api.Note) (api.CreateNoteRes, error) {
	title := req.Title
	content := req.Content
	err := h.noteService.CreateNote(ctx, title.Value, content.Value)
	if err != nil {
		return nil, err
	}
	note := &api.Note{
		Title:   title,
		Content: content,
	}
	return note, nil
}

func (h *handler) DeleteNote(ctx context.Context, params api.DeleteNoteParams) (api.DeleteNoteRes, error) {
	noteId := params.NoteId
	err := h.noteService.DeleteNoteById(ctx, noteId)
	if err != nil {
		return nil, err
	}
	deletedNoteId := &api.DeleteNote{}
	return deletedNoteId, nil
}

func (h *handler) GetNote(ctx context.Context, params api.GetNoteParams) (api.GetNoteRes, error) {
	noteId := params.NoteId
	res, err := h.noteService.FetchNoteById(ctx, noteId)
	if err != nil {
		return nil, err
	}

	createdAt, err := date.TimeToString(res.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := date.TimeToString(res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	note := &api.Note{
		Title:     api.NewOptString(res.Title),
		Content:   api.NewOptString(res.Content),
		CreatedAt: api.NewOptDateTime(createdAt),
		UpdatedAt: api.NewOptDateTime(updatedAt),
	}
	return note, nil
}

func (h *handler) GetNotes(ctx context.Context) (api.GetNotesRes, error) {
	res, err := h.noteService.FetchNotes(ctx)
	if err != nil {
		return nil, err
	}
	notes := &api.Notes{}
	for _, n := range res {
		createdAt, err := date.TimeToString(n.CreatedAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err := date.TimeToString(n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes.Notes = append(notes.Notes, api.Note{
			Title:     api.NewOptString(n.Title),
			Content:   api.NewOptString(n.Content),
			CreatedAt: api.NewOptDateTime(createdAt),
			UpdatedAt: api.NewOptDateTime(updatedAt),
		})
	}
	return notes, nil
}

func (h *handler) UpdateNote(ctx context.Context, req *api.Note, params api.UpdateNoteParams) (api.UpdateNoteRes, error) {
	noteId := params.NoteId
	title := req.Title
	content := req.Content
	err := h.noteService.UpdateNote(ctx, noteId, title.Value, content.Value)
	if err != nil {
		return nil, err
	}

	updatedNote := &api.UpdateNote{
		Title:   title,
		Content: content,
	}
	return updatedNote, nil
}

func (h *handler) CreateUser(ctx context.Context, req *api.User) (api.CreateUserRes, error) {
	name := req.Name
	email := req.Email
	password := req.Password
	res, err := h.userService.CreateUser(ctx, name.Value, email.Value, password.Value)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (h *handler) DeleteUser(ctx context.Context, params api.DeleteUserParams) (api.DeleteUserRes, error) {
	userId := params.UserId
	err := h.userService.DeleteUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &api.DeleteUser{}, nil
}
func (h *handler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {
	userId := params.UserId
	res, err := h.userService.FetchUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	createdAt, err := date.TimeToString(res.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := date.TimeToString(res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user := &api.User{
		ID:        api.NewOptUUID(res.UserID),
		Password:  api.NewOptString(res.Password),
		Name:      api.NewOptString(res.Username),
		Email:     api.NewOptString(res.UserEmail),
		CreatedAt: api.NewOptDateTime(createdAt),
		UpdatedAt: api.NewOptDateTime(updatedAt),
	}
	return user, nil
}

func (h *handler) GetUsers(ctx context.Context) (api.GetUsersRes, error) {
	res, err := h.userService.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := &api.Users{}
	for _, u := range res {
		createdAt, err := date.TimeToString(u.CreatedAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err := date.TimeToString(u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, api.User{
			ID:        api.NewOptUUID(u.UserID),
			Name:      api.NewOptString(u.Username),
			Email:     api.NewOptString(u.UserEmail),
			Password:  api.NewOptString(u.Password),
			CreatedAt: api.NewOptDateTime(createdAt),
			UpdatedAt: api.NewOptDateTime(updatedAt),
		})
	}
	users.TotalCount = api.NewOptInt32(int32(len(users.Users)))
	users.Count = api.NewOptInt32(int32(len(users.Users)))
	users.NextPagetoken = api.NewOptString("")
	slog.Log(ctx, logger.SeverityInfo, "@app", "request Id", middleware.GetRequestID(ctx), "response", users)
	return users, nil
}

func (h *handler) UpdateUser(ctx context.Context, req *api.UpdateUser, params api.UpdateUserParams) (api.UpdateUserRes, error) {
	userId, err := uuid.Parse(params.UserId)
	if err != nil {
		return nil, err
	}
	name := req.Name
	email := req.Email
	password := req.Password

	user, err := h.userService.FetchUserById(ctx, userId.String())
	if err != nil {
		return nil, err
	}
	if name.Value == "" {
		name.Value = user.Username
	}
	if email.Value == "" {
		email.Value = user.UserEmail
	}
	if password.Value == "" {
		password.Value = user.Password
	}

	err = h.userService.UpdateUser(ctx, userId, name.Value, email.Value)
	if err != nil {
		return nil, err
	}
	user, err = h.userService.FetchUserById(ctx, userId.String())
	if err != nil {
		return nil, err
	}
	createdAt, err := date.TimeToString(user.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := date.TimeToString(user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	updatedUser := &api.User{
		ID:        api.NewOptUUID(userId),
		Name:      name,
		Email:     email,
		Password:  api.NewOptString(password.Value),
		CreatedAt: api.NewOptDateTime(createdAt),
		UpdatedAt: api.NewOptDateTime(updatedAt),
	}
	return updatedUser, nil
}

func (h *handler) NewError(ctx context.Context, err error) *api.CommonErrorStatusCode {
	res := &api.CommonErrorStatusCode{}
	if errors.Is(err, fmt.Errorf("NOT FOUND")) {
		res.StatusCode = 404
		res.Response = api.CommonError{Code: api.NewOptInt32(404), Message: api.NewOptString("NOT FOUND")}
	}
	return res
}
