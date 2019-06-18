package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logrusLogger *logrus.Logger

var level logrus.Level

var AppName string
var HostName string
var Environment string

func init() {

	fmt.Println("Initializing logger")

	level = DEFAULT_LEVEL
	AppName = DEFAULT_APPNAME
	HostName, _ = os.Hostname()
	Environment = DEFAULT_ENVIRONMENT

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(DEFAULT_LEVEL)

	logrusLogger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}
