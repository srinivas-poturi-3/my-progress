package logs

import (
	"fmt"
)

// LogLevel represents different log levels.
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var logLevelNames = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

// Logger represents a custom logger.
type Logger struct {
	level LogLevel
}

// NewLogger creates a new logger with the specified log level.
func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

// Log prints the log message with the appropriate log level.
func (l *Logger) Log(level LogLevel, message string) {
	if level >= l.level {
		fmt.Printf("[%s] %s\n", logLevelNames[level], message)
	}
}

// Logf prints the log message with the appropriate log level.
func (l *Logger) Logf(level LogLevel, message string, a ...any) {
	if level >= l.level {
		msg := fmt.Sprintf(message, a)
		fmt.Printf("[%s] %s\n", logLevelNames[level], msg)
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(message string) {
	l.Log(Debug, message)
}

// Info logs an info message.
func (l *Logger) Info(message string) {
	l.Log(Info, message)
}

func (l *Logger) Infof(message string, a any) {
	l.Logf(Info, message, a)
}

// Warn logs a warning message.
func (l *Logger) Warn(message string) {
	l.Log(Warn, message)
}

// Warnf logs a warning message.
func (l *Logger) Warnf(message string, r any) {
	l.Logf(Warn, message, r)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.Log(Error, message)
}

// Errorf logs an error message.
func (l *Logger) Errorf(message string, a any) {
	l.Logf(Error, message, a)
}
