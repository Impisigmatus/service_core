package log

import (
	"os"

	"github.com/rs/zerolog"
)

type ctxKey struct{}

var CtxKey = ctxKey{}

func New(lvl Level) zerolog.Logger {
	zerolog.CallerSkipFrameCount = 7
	zerolog.LevelFieldName = "lvl"
	zerolog.TimestampFieldName = "dt"
	zerolog.CallerFieldName = "call"
	zerolog.MessageFieldName = "msg"

	return zerolog.New(os.Stdout).
		Level(lvl.Zerolog()).
		With().Timestamp().Caller().
		Logger()
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelPanic
)

type Level zerolog.Level

func (lvl *Level) Zerolog() zerolog.Level {
	return zerolog.Level(*lvl)
}
