package logging

type Logger interface {
	WithError(err error) Logger
}

// WithError returns a reusable logger interface decorated with an "error=<err>" tag
func WithError(err error) Logger {
	return New().WithError(err)
}
