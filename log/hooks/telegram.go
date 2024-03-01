package hooks

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/Impisigmatus/service_core/log"
	"github.com/rs/zerolog"
	slogtelegram "github.com/samber/slog-telegram"
)

var toSlog = map[zerolog.Level]slog.Level{
	zerolog.TraceLevel: slog.LevelDebug,
	zerolog.DebugLevel: slog.LevelDebug,
	zerolog.InfoLevel:  slog.LevelInfo,
	zerolog.WarnLevel:  slog.LevelWarn,
	zerolog.ErrorLevel: slog.LevelError,
	zerolog.FatalLevel: slog.LevelError,
	zerolog.PanicLevel: slog.LevelError,
	zerolog.NoLevel:    slog.LevelError,
	zerolog.Disabled:   slog.LevelError,
}

type telegramHook struct {
	lvl    zerolog.Level
	logger *slog.Logger
}

func NewTelegramHook(lvl log.Level, token string, chatID string) zerolog.Hook {
	logger := slog.New(slogtelegram.Option{
		Level:    toSlog[lvl.Zerolog()],
		Token:    token,
		Username: chatID,
	}.NewTelegramHandler())
	return &telegramHook{
		lvl:    lvl.Zerolog(),
		logger: logger,
	}
}

func (hook *telegramHook) Run(event *zerolog.Event, lvl zerolog.Level, message string) {
	if lvl >= hook.lvl {
		timestamp := time.Now()
		hook.logger.
			With("Level", strings.ToUpper(lvl.String())).
			With("Time", timestamp.Format(time.TimeOnly)).
			With("Date", timestamp.Format(time.DateOnly)).
			Log(context.Background(), toSlog[lvl], message)
	}
}
