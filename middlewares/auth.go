package middlewares

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Impisigmatus/service_core/log"
	"github.com/Impisigmatus/service_core/utils"
	"github.com/rs/zerolog"
)

type authorization struct {
	next    http.Handler
	secrets []string
}

func Authorization(secrets []string) Middleware {
	return func(next http.Handler) http.Handler {
		return &authorization{
			next:    next,
			secrets: secrets,
		}
	}
}

func (m *authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const header = "Authorization"
	log, ok := r.Context().Value(log.CtxKey).(zerolog.Logger)
	if !ok {
		utils.WriteString(zerolog.Logger{}, w, http.StatusInternalServerError, errors.New("Invalid logger"), "Невалидный логгер")
		return
	}

	authorization := r.Header.Get(header)
	if !strings.HasPrefix(authorization, "Basic") {
		utils.WriteString(log, w, http.StatusUnauthorized, errors.New("Invalid type"), "Неверный тип авторизации")
		return
	}

	if err := m.basic(authorization); err != nil {
		utils.WriteString(log, w, http.StatusUnauthorized, err, "Неверные логин или пароль")
		return
	}

	m.next.ServeHTTP(w, r)
}

func (m *authorization) basic(header string) error {
	const prefix = "Basic "
	authorization := header[len(prefix):]
	data, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return fmt.Errorf("Invalid decode basic authorization: %w", err)
	}
	decoded := string(data)

	for _, secret := range m.secrets {
		if decoded == secret {
			return nil
		}
	}

	return errors.New("Invalid basic authorization")
}
