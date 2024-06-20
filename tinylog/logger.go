package tinylog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger is used to log application messages.
type Logger struct {
	threshold        Level
	output           io.Writer
	includeTimestamp bool
	includeLogLevel  bool
}

// New returns you a logger, ready to log at the required level.
// Provide a list of configuration functions to configure the logger.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a log message when the log level is Debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, LevelDebug, args...)
}

// Debug prints a log message when the log level is Debug or higher.
func (l *Logger) Debug(message string) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf("%s", LevelDebug, message)
}

// Infof formats and prints a log message when the log level is Info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(format, LevelInfo, args...)
}

// Info prints a log message when the log level is Info or higher.
func (l *Logger) Info(message string) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf("%s", LevelInfo, message)
}

// Errorf formats and prints a log message when the log level is Error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(format, LevelError, args...)
}

// Error prints a log message when the log level is Error or higher.
func (l *Logger) Error(message string) {
	if l.threshold > LevelError {
		return
	}

	l.logf("%s", LevelError, message)
}

// logf formats and prints the message to the output.
func (l *Logger) logf(format string, lvl Level, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	var newArgs []any
	var newFormat string

	if l.includeLogLevel {
		newArgs = append(newArgs, lvl.DisplayName())
		newFormat = newFormat + "[ %s ] "
	}

	if l.includeTimestamp {
		newArgs = append(newArgs, time.Now().Format(time.RFC3339))
		newFormat = newFormat + "%v - "
	}

	newArgs = append(newArgs, args...)
	newFormat = newFormat + format + "\n"

	_, _ = fmt.Fprintf(l.output, newFormat, newArgs...)
}
