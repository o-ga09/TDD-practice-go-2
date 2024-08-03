package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/o-ga09/note-app-backendapi/api"
	"github.com/o-ga09/note-app-backendapi/pkg/uuid"
	"github.com/ogen-go/ogen/middleware"
)

type RequestId string

func AddID() api.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		ctx := context.WithValue(req.Context, RequestId("requestId"), uuid.GenerateID())
		req.SetContext(ctx)
		res, err := next(req)
		if err != nil {
			slog.InfoContext(ctx, fmt.Sprintf("error: %v", err))
		} else {
			slog.InfoContext(ctx, fmt.Sprintf("response: %T", res.Type))
		}
		return res, nil
	}
}

func WithTimeout() api.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		ctx, cancel := context.WithTimeout(req.Context, 5*time.Second)
		defer cancel()
		req.SetContext(ctx)
		res, err := next(req)
		if err != nil {
			slog.InfoContext(ctx, fmt.Sprintf("error: %v", err))
		} else {
			slog.InfoContext(ctx, fmt.Sprintf("response: %T", res.Type))
		}
		return res, nil
	}
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(RequestId("requestId")).(string)
}
