package logging

import (
	logrus "github.com/sirupsen/logrus"
)

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
	return logrus.WithFields(logrus.Fields{})
}

func (lw *logWriter) WithError(err error) Logger {
	ll := lw.getEntry().WithError(err)
	lw.setEntry(ll)
	return lw
}

func (lw *logWriter) WithField(key string, value interface{}) Logger {
	ll := lw.getEntry().WithField(key, value)
	lw.setEntry(ll)
	return lw
}

func (lw *logWriter) Error(args ...interface{}) {
	if level < logrus.ErrorLevel {
		return
	}
	lw.getEntry().Error(args...)
}

func (lw *logWriter) WithFields(fields map[string]interface{}) Logger {
	ll := lw.getEntry().WithFields(fields)
	lw.setEntry(ll)
	return lw
}

func (lw *logWriter) Debug(args ...interface{}) {
	if level < logrus.DebugLevel {
		return
	}
	lw.getEntry().Debug(args...)
}

func (lw *logWriter) Info(args ...interface{}) {
	if level < logrus.InfoLevel {
		return
	}
	lw.getEntry().Info(args...)
}
