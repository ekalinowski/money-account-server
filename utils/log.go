package utils

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		log.Formatter = new(logrus.JSONFormatter)
		log.Formatter = new(logrus.TextFormatter) // default
		log.Level = logrus.DebugLevel
		log.Info("Logger client initiated.")
	}

	return log
}
