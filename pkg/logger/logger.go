package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	LogrusLogger *logrus.Logger
)

func Init() {
	LogrusLogger = InitLogrusLogger()
	if LogrusLogger == nil {
		panic("Failed to initialize logger")
	}
}
