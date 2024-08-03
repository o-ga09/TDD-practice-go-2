package middleware

// import (
// 	"context"
// 	"errors"

// 	"github.com/o-ga09/note-app-backendapi/api"
// )

// var _ api.SecurityHandler = (*SecurityHandler)(nil)

// type SecurityHandler struct{}

// func (*SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
// 	if t.Token == "" {
// 		return ctx, errors.New("token is empty")
// 	}

// 	return ctx, nil
// }

// func (*SecurityHandler) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
// 	if t.APIKey == "" {
// 		return ctx, errors.New("api key is empty")
// 	}

// 	return ctx, nil
// }
