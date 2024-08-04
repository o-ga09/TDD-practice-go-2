package middleware

import (
	"log/slog"
	"time"

	"github.com/o-ga09/note-app-backendapi/api"
	"github.com/o-ga09/note-app-backendapi/pkg/logger"
	"github.com/ogen-go/ogen/middleware"
)

type RequestInfo struct {
	method, path, sourceIP, query, user_agent string
	elapsed                                   time.Duration
}

func RequestLogger() api.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		start := time.Now()
		slog.Log(req.Context, logger.SeverityInfo, "処理開始", "request Id", GetRequestID(req.Context))
		res, err := next(req)
		if err != nil {
			slog.Log(req.Context, logger.SeverityError, "処理エラー", "error", err)
		}
		r := &RequestInfo{
			path:       req.Raw.URL.Path,
			sourceIP:   req.Raw.RemoteAddr,
			query:      req.Raw.URL.RawQuery,
			user_agent: req.Raw.UserAgent(),
			elapsed:    time.Since(start),
		}
		slog.Log(req.Context, logger.SeverityInfo, "処理終了", "Request", r.LogValue(), "requestId", GetRequestID(req.Context))
		return res, nil
	}
}

func (r *RequestInfo) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("method", r.method),
		slog.String("path", r.path),
		slog.String("sourceIP", r.sourceIP),
		slog.String("query", r.query),
		slog.String("user_agent", r.user_agent),
		slog.String("elapsed", r.elapsed.String()),
	)
}
