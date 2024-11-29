package utils

import (
	"log"
	"os"
)

// LogLevel represents the severity of the log message
type LogLevel string

const (
	Info     LogLevel = "INFO"
	Warning  LogLevel = "WARNING"
	Error    LogLevel = "ERROR"
	Critical LogLevel = "CRITICAL"
)

// Logger is a custom logger structure
type Logger struct {
	logger *log.Logger
}

// NewLogger initializes and returns a new logger instance
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Log logs a message with the specified log level
func (l *Logger) Log(level LogLevel, message string) {
	l.logger.Printf("[%s] %s\n", level, message)
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.Log(Info, message)
}

// Warning logs a warning message
func (l *Logger) Warning(message string) {
	l.Log(Warning, message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.Log(Error, message)
}

// Critical logs a critical error message
func (l *Logger) Critical(message string) {
	l.Log(Critical, message)
}
