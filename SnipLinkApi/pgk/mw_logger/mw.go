package mwlogger

import (
	"api/internal/logger"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func New(logs *logger.Logs) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger := logs.MW.With(
				slog.String("Method", r.Method),
				slog.String("Path", r.URL.Path),
				slog.String("IP", r.RemoteAddr),
				slog.String("RequestID", middleware.GetReqID(r.Context())),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			ct := time.Now()

			next.ServeHTTP(ww, r)
			logger.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"Request",
				slog.Int("Status", ww.Status()),
				slog.Int("Byte", ww.BytesWritten()),
				slog.String("TimeRequest", time.Since(ct).String()),
			)
		}

		return http.HandlerFunc(fn)
	}
}
