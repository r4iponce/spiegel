package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func (config Config) Init() {
	if config.File != "" {
		file, err := os.OpenFile(config.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o640) //nolint:mnd
		if err != nil {
			logrus.Fatal("Cannot open log file: ", err)
		}
		logrus.SetOutput(file)
	}

	if config.Level == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if config.Level == "INFO" {
		logrus.SetLevel(logrus.InfoLevel)
	}
	if config.Level == "WARN" {
		logrus.SetLevel(logrus.WarnLevel)
	}
	if config.Level == "ERROR" {
		logrus.SetLevel(logrus.ErrorLevel)
	}
	if config.Level == "FATAL" {
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
