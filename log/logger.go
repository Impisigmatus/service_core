package log

import (
	"os"

	"github.com/Impisigmatus/service_core/log/hooks"
	"github.com/rs/zerolog"
)

type Level zerolog.Level

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var logger zerolog.Logger

func Init(lvl Level, path string) {
	zerolog.CallerSkipFrameCount = 3
	zerolog.LevelFieldName = "lvl"
	zerolog.ErrorFieldName = "err"
	zerolog.TimestampFieldName = "dt"
	zerolog.CallerFieldName = "call"
	zerolog.MessageFieldName = "msg"

	logger = zerolog.New(os.Stdout).
		Hook(hooks.NewFileHook(path)).
		Level(zerolog.Level(lvl)).
		With().Timestamp().Caller().
		Logger()
}

func Get() zerolog.Logger {
	return logger
}
