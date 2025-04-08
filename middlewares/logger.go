package middlewares

import (
	"net/http"
	"time"

	"github.com/Impisigmatus/service_core/log"
	chi "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

type contextLogger struct {
	next http.Handler
}

func ContextLogger() Middleware {
	return func(next http.Handler) http.Handler {
		return &contextLogger{next: next}
	}
}

func (m *contextLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware := chi.RequestLogger(&requestLogger{})
	middleware(m.next).ServeHTTP(w, r)
}

type requestLogger struct{}

func (*requestLogger) NewLogEntry(r *http.Request) chi.LogEntry {
	logger, _ := r.Context().Value(log.CtxKey).(zerolog.Logger)
	return &entry{
		logger:   logger,
		Method:   r.Method,
		Hostname: r.RemoteAddr,
		Path:     r.URL.Path,
	}
}

type entry struct {
	logger   zerolog.Logger
	Method   string
	Hostname string
	Path     string
}

func (e *entry) Write(status int, _ int, _ http.Header, duration time.Duration, extra interface{}) {
	e.logger.Info().Msgf("%s %s | %s | %s | %d %s",
		e.Method,
		e.Path,
		e.Hostname,
		duration,
		status,
		http.StatusText(status),
	)
}

func (e *entry) Panic(v interface{}, stack []byte) {
	e.logger.Error().Msgf("panic occurred panic: %+v stack:%s, ", v, stack)
}
