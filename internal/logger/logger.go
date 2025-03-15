package logger

import (
	"github.com/sirupsen/logrus"
)

// Global logger variable
var log *logrus.Logger

// Initialize logger
func InitLogger() {
	// Create a new logger instance
	log = logrus.New()

	// Set the log format (e.g., JSON or Text)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Add timestamps
	})

	// Optionally, set log level (use logrus.DebugLevel, logrus.InfoLevel, etc.)
	log.SetLevel(logrus.InfoLevel)
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger() // Ensure the logger is initialized
	}
	return log
}

func LogError(err error, msg string) {
	GetLogger().WithField("error", err.Error()).Error(msg)
}
