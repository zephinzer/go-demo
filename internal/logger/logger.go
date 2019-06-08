package logger

import (
	"github.com/sirupsen/logrus"
)

var FieldMap = logrus.FieldMap{
	logrus.FieldKeyFile: "@file",
	logrus.FieldKeyFunc: "@func",
	logrus.FieldKeyTime: "@timestamp",
	logrus.FieldKeyLevel: "@level",
	logrus.FieldKeyMsg: "@msg",
}

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableSorting: true,
		FieldMap: FieldMap,
		QuoteEmptyFields: true,
		TimestampFormat: "01/02-15:04:05",
		FullTimestamp: true,
	})
	return logger
}