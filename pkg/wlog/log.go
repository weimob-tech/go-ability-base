package wlog

import (
	"log"
	"os"
)

var (
	// Provide default logger for users to use
	logger FullLogger = &defaultLogger{
		stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		depth:  4,
	}
)

func GetLogger() FullLogger {
	return logger
}

func SetLogger(l FullLogger) {
	logger = l
}
