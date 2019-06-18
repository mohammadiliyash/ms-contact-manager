package logging

type Logger interface {
	WithError(err error) Logger
	Error(args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(map[string]interface{}) Logger
	Info(args ...interface{})
}
