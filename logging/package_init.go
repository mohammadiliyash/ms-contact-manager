package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logrusLogger *logrus.Logger

var level logrus.Level

// AppName stores the application name for automatically added "app-name=<AppName>" tags
// If not set at runtime, it will display as app-name=unconfigured-app
var AppName string

// HostName stores the hostname for automatically added "hostname=<hostname>" tags
// By default this is set to os.HostName(), which is probably what you want
var HostName string

// Environment stores the environment (QA, Prod, etc.) for automatically-added "environment=<environment>" tags
// If not set at runtime, it will display as environment=unconfigured
var Environment string

func init() {
	fmt.Println("Initializing logger")

	// Initialize with default values, just to get us rolling
	// These MUST be overwritten by the consuming app
	// After the app's full config has been loaded and parsed
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

const DEFAULT_APPNAME = "unconfigured-app"
const DEFAULT_ENVIRONMENT = "unconfigured"
