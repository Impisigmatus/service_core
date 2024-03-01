package hooks

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Impisigmatus/service_core/log"
	"github.com/rs/zerolog"
)

type fileHook struct {
	lvl  zerolog.Level
	file *os.File
}

func NewFileHook(lvl log.Level, path string) zerolog.Hook {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	return &fileHook{
		lvl:  lvl.Zerolog(),
		file: file,
	}
}

func (hook *fileHook) Run(event *zerolog.Event, lvl zerolog.Level, msg string) {
	if lvl >= hook.lvl {
		msg = "{\"lvl\":\"" + strings.ToUpper(lvl.String()) +
			"\",\"dt\":\"" + time.Now().Format(time.DateTime) +
			"\",\"msg\":\"" + msg + "\"}\n"
		n, err := hook.file.WriteString(msg)
		if err != nil {
			panic(err)
		}
		if length := len(msg); n != length {
			panic(strconv.Itoa(n) + "/" + strconv.Itoa(length) + " bytes writed")
		}
	}
}
