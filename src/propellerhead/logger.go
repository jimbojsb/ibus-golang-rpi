package propellerhead

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = nil

func Logger() (*logrus.Logger) {
	if (logger == nil) {
		logger = logrus.New()
		logger.Level = logrus.DebugLevel
	}
	return logger
}
