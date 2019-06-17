package logging

import "context"

type Logger interface {
	WithError(err error) Logger
	CError(c context.Context, args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) Logger
}

// WithError returns a reusable logger interface decorated with an "error=<err>" tag
func WithError(err error) Logger {
	return New().WithError(err)
}

// WithField returns a reusable logger interface decorated with "key=value" tag
func WithField(key string, value interface{}) Logger {
	return New().WithField(key, value)
}
