package logger

import (
	"base-monitor-service/pkg/config"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// SetupLogger initializes the logger with configurations from cfg
// SetupLogger sets up the logger configuration based on the provided config.
// It adjusts the log level and format according to the configuration.
// The log level can be set to "debug", "info", "warn", or "error".
// If an invalid log level is provided, it defaults to "info".
// The log format includes colors, full timestamp, and a custom timestamp format.
// The log output is directed to os.Stdout.
func SetupLogger(cfg *config.Config) {
	// This is a simplified example. Adjust log level and format as needed.
	switch cfg.LogLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: cfg.LogTimeFormat,
	})
	log.SetOutput(os.Stdout)
}

// LogField wraps logrus.Fields to avoid logrus import in other packages
type LogField map[string]interface{}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// WithFields returns a new log entry with additional fields.
// It takes a slice of LogField, which is a map[string]interface{},
// and adds each key-value pair to the log entry's fields.
// The function returns a pointer to the new log entry.
func WithFields(fields []LogField) *logrus.Entry {
	logFields := make(logrus.Fields)
	for _, field := range fields {
		for key, value := range field {
			logFields[key] = value
		}
	}
	return log.WithFields(logFields)
}

// Info logs the provided arguments as an information message.
// It uses the underlying log package's Info function to log the message.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Debug logs debug-level messages.
// It takes variadic arguments of type interface{} and logs them using the underlying logger's Debug method.
// Example usage: Debug("This is a debug message")
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Warn logs a warning message.
// It accepts variadic arguments of type interface{}.
// The warning message is logged using the underlying log package's Warn function.
// Example usage: Warn("This is a warning message")
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logs an error message.
// It accepts variadic arguments of type interface{}.
// The error message is logged using the log package's Error function.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fields is a function that returns the provided logrus.Fields.
// It is used to pass additional fields to the logger.
func Fields(fields logrus.Fields) logrus.Fields {
	return fields
}
