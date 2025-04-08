package utils

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
)

func WriteNoContent(log zerolog.Logger, w http.ResponseWriter) {
	WriteString(log, w, http.StatusNoContent, nil, "")
}

func WriteString(log zerolog.Logger, w http.ResponseWriter, status int, err error, str string, args ...any) {
	const (
		header  = "Content-Type"
		content = "text/plain"
	)

	if err != nil {
		log.Error().Msg(err.Error())
	}

	w.Header().Set(header, content)
	w.WriteHeader(status)

	if _, err := w.Write([]byte(fmt.Sprintf(str, args...))); err != nil {
		log.Error().Msgf("Invalid write response body: %s", err)
	}
}

func WriteObject(log zerolog.Logger, w http.ResponseWriter, obj interface{}) {
	const (
		header  = "Content-Type"
		content = "application/json"
	)

	w.Header().Set(header, content)
	w.WriteHeader(http.StatusOK)

	data, err := jsoniter.Marshal(obj)
	if err != nil {
		WriteString(log, w, http.StatusInternalServerError, fmt.Errorf("Invalid parse body: %w", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if _, err := w.Write(data); err != nil {
		log.Error().Msgf("Invalid write response body: %s", err)
	}
}
