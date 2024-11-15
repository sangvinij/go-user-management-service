package logger

import (
	"github.com/sirupsen/logrus"
)

type LogHookParams struct {
	file string
	line int
}

type LogHook struct {
	LogHookParams
}

func (h *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogHook) Fire(entry *logrus.Entry) error {
	entry.Data["file"] = h.file
	entry.Data["line"] = h.line
	return nil
}

func GetLogHook(file string, line int) *LogHook {
	return &LogHook{LogHookParams: LogHookParams{file: file, line: line}}
}
