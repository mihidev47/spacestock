// Package logger provides functionality for logging
package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"

	"../config"
)

type AppLogger struct {
	Hostname string
	*logrus.Logger
}

// Init get logger configuration and initiate logger instance
func Init() {
	ls := config.MustGetString("logger.level")
	var lv logrus.Level
	// Get log level
	switch ls {
	case "info":
		lv = logrus.InfoLevel
	default:
		lv = logrus.DebugLevel
	}
	// Set log level
	logrus.SetLevel(lv)
	// If server mode is production, set output to json
	if config.MustGetString("logger.level") == "production" {
		// Get path from configuration
		path := config.MustGetString("log.output_file")
		// Open file
		file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			fmt.Printf("Unable to open log output file. Error: %s\n", err.Error())
			os.Exit(4)
		}
		// Set formatter to json
		logrus.SetFormatter(&logrus.JSONFormatter{})
		// Set output file
		logrus.SetOutput(file)
	}
}

// Get returns logger instance. App will exit if an error occurred while getting logger
func Get() AppLogger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return AppLogger{hostname, logrus.StandardLogger()}
}

func (l AppLogger) Info(args ...interface{}) {
	l.addFields().Info(args...)
}

func (l AppLogger) Infof(format string, args ...interface{}) {
	l.addFields().Infof(format, args...)
}

func (l AppLogger) Debug(args ...interface{}) {
	l.addFields().Debug(args...)
}

func (l AppLogger) Debugf(format string, args ...interface{}) {
	l.addFields().Debugf(format, args...)
}

// DebugCtxf is error logging with context
func (l AppLogger) DebugCtxf(context string, format string, args ...interface{}) {
	l.addContextFields(context).Debugf(format, args...)
}

func (l AppLogger) Error(args ...interface{}) {
	l.addFields().Error(args...)
}

func (l AppLogger) Errorf(format string, args ...interface{}) {
	l.addFields().Errorf(format, args...)
}

// ErrorCtx is error logging with context
func (l AppLogger) ErrorCtx(context string, args ...interface{}) {
	l.addContextFields(context).Error(args...)
}

// ErrorCtxf is formatted error logging with context
func (l AppLogger) ErrorCtxf(context string, format string, args ...interface{}) {
	l.addContextFields(context).Errorf(format, args...)
}

// addFields add additional info to log
func (l AppLogger) addFields() *logrus.Entry {
	file, line := getCaller()
	return l.Logger.WithField("hostname", l.Hostname).
		WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func (l AppLogger) addContextFields(context string) *logrus.Entry {
	file, line := getCaller()
	return l.Logger.WithField("hostname", l.Hostname).
		WithField("source", fmt.Sprintf("%s:%d", file, line)).
		WithField("context", context)
}

// getCaller returns where a code is executed
func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		s := strings.LastIndex(file, "/")
		file = file[s+1:]
	}
	return file, line
}
