package logrus

import (
	"io"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// NewLogger create a logrus logger, add hook to it and return it.
func NewLogger(zapLogger *zap.Logger) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	logger.AddHook(newHook(zapLogger))

	return logger
}
