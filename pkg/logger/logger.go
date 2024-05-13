package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	LogrusLogger *logrus.Logger
)

func Init() {
	LogrusLogger = InitLogrusLogger()

	// ZapLogger = InitZapLogger()
	// ZapSugar = ZapLogger.Sugar()
}
