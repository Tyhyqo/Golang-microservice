package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
	log := logrus.New()

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.Panic("Invalid log level")
	}
	log.SetLevel(lvl)

	return log
}
