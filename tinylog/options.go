package tinylog

import "io"

// Option defines a functional option to our logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output for logs.
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}

func WithTimestamp() Option {
	return func(lgr *Logger) {
		lgr.includeTimestamp = true
	}
}

func WithLogLevel() Option {
	return func(lgr *Logger) {
		lgr.includeLogLevel = true
	}
}