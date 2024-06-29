package telegram

import (
	"fmt"
	"time"

	"github.com/Impisigmatus/service_core/log"
	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source=telegram.go -package mocks -destination ../autogen/mocks/telegram.go
type IReciever interface {
	Handle(payload *tg_bot.Message) (*tg_bot.MessageConfig, error)
	GetOffset() int
	GetTimeout() time.Duration
}

type _ITelegramAPI interface {
	Send(c tg_bot.Chattable) (tg_bot.Message, error)
	GetUpdatesChan(cfg tg_bot.UpdateConfig) tg_bot.UpdatesChannel
}

type Telegram struct {
	api      _ITelegramAPI
	receiver IReciever
}

func New(token string, receiver IReciever) *Telegram {
	if receiver == nil {
		log.Panicf("Invalid init: receiver is nil")
	}
	api, err := tg_bot.NewBotAPI(token)
	if err != nil {
		log.Panicf("Invalid telegram api: %s", err)
	}

	tg := newTelegram(api, receiver)
	tg.consume()
	return tg
}

func newTelegram(api _ITelegramAPI, receiver IReciever) *Telegram {
	return &Telegram{api: api, receiver: receiver}
}

func (tg *Telegram) Send(chatID uint64, data string) error {
	msg := tg_bot.NewMessage(int64(chatID), data)
	if _, err := tg.api.Send(msg); err != nil {
		return fmt.Errorf("Invalid send: %s", err)
	}

	return nil
}

func (tg *Telegram) consume() {
	go func() {
		updater := tg_bot.NewUpdate(tg.receiver.GetOffset())
		updater.Timeout = int(tg.receiver.GetTimeout().Seconds())
		updates := tg.api.GetUpdatesChan(updater)

		for update := range updates {
			if update.Message != nil {
				msg, err := tg.receiver.Handle(update.Message)
				if err != nil {
					log.Error("Invalid handle msg", err)
					continue
				}
				if _, err := tg.api.Send(msg); err != nil {
					log.Error("Invalid send", err)
					continue
				}
			}
		}
	}()
}
