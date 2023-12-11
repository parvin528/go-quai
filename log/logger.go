package log

import (
	"github.com/adrg/xdg"
	"github.com/sirupsen/logrus"
)

const (
	// default log level
	defaultLogLevel = "info"

	// log file name
	logFileName = "go-quai.log"
	// default log directory
	logDir = "nodelogs"
	// default log file params
	defaultLogMaxSize    = 100  // maximum file size before rotation, in MB
	defaultLogMaxBackups = 3    // maximum number of old log files to keep
	defaultLogMaxAge     = 28   // maximum number of days to retain old log files
	defaultLogCompress   = true // whether to compress the rotated log files using gzip
)

var (
	// logger instance used by the application
	logger Logger

	// TODO: consider refactoring to dinamically read the app name (i.e. "go-quai") ?
	// default logfile path
	defaultLogFilePath = xdg.DataHome + "/" + "go-quai" + "/" + logDir + "/" + logFileName
)

func init() {
	entry := logrus.NewEntry(logrus.StandardLogger())
	logger = &LogWrapper{
		entry: entry,
	}
	ConfigureLogger(
		WithLevel(defaultLogLevel),
		WithOutput(ToStdOut(), ToLogFile(defaultLogFilePath)),
	)
	logger.Infof("Logger started. Writing logs to: %s", defaultLogFilePath)
}

func ConfigureLogger(opts ...Options) {
	for _, opt := range opts {
		opt(logger.(*LogWrapper))
	}
}

func Trace(keyvals ...interface{}) {
	logger.Trace(keyvals...)
}

func Tracef(msg string, args ...interface{}) {
	logger.Tracef(msg, args...)
}

func Debug(keyvals ...interface{}) {
	logger.Debug(keyvals...)
}

func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func Info(keyvals ...interface{}) {
	logger.Info(keyvals...)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Warn(keyvals ...interface{}) {
	logger.Warn(keyvals...)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func Error(keyvals ...interface{}) {
	logger.Error(keyvals...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func Fatal(keyvals ...interface{}) {
	logger.Fatal(keyvals...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

func WithField(key string, val interface{}) Logger {
	return logger.WithField(key, val)
}