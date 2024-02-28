package telegram

import (
	"fmt"

	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	api      *tg_bot.BotAPI
	reciever IReciever
}

func New(token string, reciever IReciever) *Telegram {
	if reciever == nil {
		logrus.Panicf("Invalid init: reciever is nil")
	}
	api, err := tg_bot.NewBotAPI(token)
	if err != nil {
		logrus.Panicf("Invalid telegram api: %s", err)
	}
	return newTelegram(api, reciever)
}

func newTelegram(api *tg_bot.BotAPI, reciever IReciever) *Telegram {
	tg := Telegram{api: api, reciever: reciever}
	tg.consume()
	return &tg
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
		updater := tg_bot.NewUpdate(tg.reciever.GetOffset())
		updater.Timeout = int(tg.reciever.GetTimeout().Seconds())
		updates := tg.api.GetUpdatesChan(updater)

		for update := range updates {
			if update.Message != nil {
				msg, err := tg.reciever.Handle(update.Message)
				if err != nil {
					logrus.Errorf("Invalid handle msg: %s", err)
					continue
				}
				if _, err := tg.api.Send(msg); err != nil {
					logrus.Errorf("Invalid send: %s", err)
					continue
				}
			}
		}
	}()
}
