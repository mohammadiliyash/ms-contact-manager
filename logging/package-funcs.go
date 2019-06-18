package logging

// Error logs an error message
func Error(args ...interface{}) {
	New().Error(args...)
}

// WithError returns a reusable logger interface decorated with an "error=<err>" tag
func WithError(err error) Logger {
	return New().WithError(err)
}

// WithField returns a reusable logger interface decorated with "key=value" tag
func WithField(key string, value interface{}) Logger {
	return New().WithField(key, value)
}

// WithFields returns a reusable logger interface decorated with any number of "key=value" tags
func WithFields(fields map[string]interface{}) Logger {
	return New().WithFields(fields)
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	New().Debug(args...)
}

// Info logs an info message
func Info(args ...interface{}) {
	New().Info(args...)
}
