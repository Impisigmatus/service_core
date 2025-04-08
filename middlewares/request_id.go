package middlewares

import (
	"context"
	"net/http"

	"github.com/Impisigmatus/service_core/log"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type requestID struct {
	next       http.Handler
	baseLogger zerolog.Logger
}

func RequestID(logger zerolog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return &requestID{
			next:       next,
			baseLogger: logger,
		}
	}
}

func (m *requestID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Request-ID")
	if len(id) == 0 {
		id = uuid.New().String()
	}

	ctx := context.WithValue(r.Context(), log.CtxKey, m.baseLogger.With().Str("request_id", id).Logger())
	m.next.ServeHTTP(w, r.WithContext(ctx))
}
