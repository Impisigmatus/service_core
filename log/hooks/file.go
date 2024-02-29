package hooks

import (
	"os"

	"github.com/rs/zerolog"
)

type FileHook struct {
	path string
}

func NewFileHook(path string) zerolog.Hook {
	return &FileHook{path: path}
}

func (hook *FileHook) Run(event *zerolog.Event, lvl zerolog.Level, message string) {
	if lvl == zerolog.DebugLevel {
		file, err := os.OpenFile(hook.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {
			return
		}
		defer file.Close()

		if _, err := file.WriteString(message + "\n"); err != nil {
			return
		}
	}
}
