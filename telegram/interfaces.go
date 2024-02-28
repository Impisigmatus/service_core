package telegram

import (
	"time"

	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IReciever interface {
	Handle(payload *tg_bot.Message) (*tg_bot.MessageConfig, error)
	GetOffset() int
	GetTimeout() time.Duration
}
