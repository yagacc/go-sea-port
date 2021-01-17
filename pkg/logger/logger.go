package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// common interface for logging
type Logger interface {
	GetLevel() logrus.Level

	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

const osKey = "LOG_LEVEL"

var (
	logstashFieldMap = logrus.FieldMap{
		logrus.FieldKeyTime:  "timestamp",
		logrus.FieldKeyMsg:   "message",
		logrus.FieldKeyLevel: "level",
		logrus.FieldKeyFile:  "logger_name",
		logrus.FieldKeyFunc:  "method",
	}
)

func NewLogrusLogger(log *logrus.Logger) Logger {
	lvl, osKeyAvailable := os.LookupEnv(osKey)
	if !osKeyAvailable {
		lvl = "debug"
	}
	ll, parseErr := logrus.ParseLevel(lvl)
	if parseErr != nil {
		ll = logrus.DebugLevel
	}

	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{FieldMap: logstashFieldMap}
	log.Level = ll

	logger := LogrusLogger{log}

	if !osKeyAvailable {
		logger.Warnf("level key [%s] missing from environment", osKey)
	}

	if parseErr != nil {
		logger.Warnf("cannot parse log level [%s], using [%s]", lvl, "DebugLevel")
	}

	return logger
}

type LogrusLogger struct {
	*logrus.Logger
}

func (l LogrusLogger) GetLevel() logrus.Level {
	return l.Logger.Level
}

func (l LogrusLogger) Errorf(format string, v ...interface{}) {
	l.Logger.Errorf(format, v...)
}

func (l LogrusLogger) Debugf(format string, v ...interface{}) {
	l.Logger.Debugf(format, v...)
}

func (l LogrusLogger) Infof(format string, v ...interface{}) {
	l.Logger.Infof(format, v...)
}

func (l LogrusLogger) Warnf(format string, v ...interface{}) {
	l.Logger.Warnf(format, v...)
}

func (l LogrusLogger) Fatalf(format string, v ...interface{}) {
	l.Errorf(format, v...)
	os.Exit(1)
}
