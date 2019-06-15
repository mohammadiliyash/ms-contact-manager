package logging

import logrus "github.com/sirupsen/logrus"

// logWriter wraps a logrus Entry
type logWriter struct {
	entry *logrus.Entry
}

func New() *logWriter {
	lw := logWriter{}
	return &lw
}

func (lw *logWriter) getEntry() *logrus.Entry {
	if lw.entry == nil {
		lw.entry = lw.createEntry()
	}
	return lw.entry
}

func (lw *logWriter) setEntry(entry *logrus.Entry) *logWriter {
	lw.entry = entry
	return lw
}

func (lw *logWriter) createEntry() *logrus.Entry {
	return logrus.WithFields(
		logrus.Fields{
			FIELDNAME_APPNAME:     AppName,
			FIELDNAME_HOSTNAME:    HostName,
			FIELDNAME_ENVIRONMENT: Environment,
		})
}

func (lw *logWriter) WithError(err error) Logger {
	ll := lw.getEntry().WithError(err)
	lw.setEntry(ll)
	return lw
}

// Error logs an error message
func Error(args ...interface{}) {
	New().Error(args...)
}

func (lw *logWriter) Error(args ...interface{}) {
	if level < logrus.ErrorLevel {
		return
	}
	lw.getEntry().Error(args...)
}
