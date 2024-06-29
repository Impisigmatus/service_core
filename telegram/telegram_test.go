package telegram

import (
	"testing"

	"github.com/Impisigmatus/service_core/autogen/mocks"
	"github.com/golang/mock/gomock"
)

func Test_Telegram(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := mocks.NewMock_ITelegramAPI(ctrl)
	receiver := mocks.NewMockIReciever(ctrl)
	tg := newTelegram(api, receiver)
	_ = tg

	t.Run("Send()", func(t *testing.T) {
		t.Run("SuccessTestCases", func(t *testing.T) {
			t.Run("TODO", func(t *testing.T) {
				// tg.Send()
			})
		})
		t.Run("FailTestCases", func(t *testing.T) {
			t.Run("TODO", func(t *testing.T) {
				// tg.Send()
			})
		})
	})

	t.Run("consume()", func(t *testing.T) {
		t.Run("SuccessTestCases", func(t *testing.T) {
			t.Run("TODO", func(t *testing.T) {
				// tg.consume()
			})
		})
		t.Run("FailTestCases", func(t *testing.T) {
			t.Run("TODO", func(t *testing.T) {
				// tg.consume()
			})
		})
	})
}
