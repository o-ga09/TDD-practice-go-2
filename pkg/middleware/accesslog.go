package middleware

import (
	"log/slog"
	"time"

	"github.com/o-ga09/note-app-backendapi/api"
	"github.com/o-ga09/note-app-backendapi/pkg/logger"
	"github.com/ogen-go/ogen/middleware"
)

type RequestInfo struct {
	status                                    int
	contents_length                           int64
	method, path, sourceIP, query, user_agent string
	elapsed                                   time.Duration
}

func RequestLogger() api.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		start := time.Now()
		slog.Log(req.Context, logger.SeverityInfo, "処理開始", "request Id", GetRequestID(req.Context))
		res, err := next(req)
		if err != nil {
			slog.InfoContext(req.Context, "処理エラー", "error", err)
		} else {
			slog.InfoContext(req.Context, "処理終了", "response", res)
		}

		r := &RequestInfo{
			status:          req.Raw.Response.StatusCode,
			contents_length: req.Raw.ContentLength,
			path:            req.Raw.URL.Path,
			sourceIP:        req.Raw.RemoteAddr,
			query:           req.Raw.URL.RawQuery,
			user_agent:      req.Raw.UserAgent(),
			elapsed:         time.Since(start),
		}
		slog.Log(req.Context, logger.SeverityInfo, "処理終了", "Request", r.LogValue(), "requestId", GetRequestID(req.Context))
		return res, nil
	}
}

func (r *RequestInfo) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("status", r.status),
		slog.Int64("Content-length", r.contents_length),
		slog.String("method", r.method),
		slog.String("path", r.path),
		slog.String("sourceIP", r.sourceIP),
		slog.String("query", r.query),
		slog.String("user_agent", r.user_agent),
		slog.String("elapsed", r.elapsed.String()),
	)
}
