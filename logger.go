package nlogger

import (
	"fmt"
	stdLog "log"
	"os"
	"runtime"
	"strings"
	"sync"
)

// Logger contract defines methods that must be available for a Logger.
type Logger interface {
	// Fatal must write an error, message that explaining the error and where it's occurred in FATAL level.
	Fatal(msg string, options ...OptionSetterFunc)

	// Fatalf must write a formatted message and where it's occurred in FATAL level.
	Fatalf(format string, args ...interface{})

	// Error must write an error, message that explaining the error and where it's occurred in ERROR level.
	Error(msg string, options ...OptionSetterFunc)

	// Errorf must write a formatted message and where it's occurred in ERROR level.
	Errorf(format string, args ...interface{})

	// Warn must write a message in WARN level.
	Warn(msg string, options ...OptionSetterFunc)

	// Warnf must write a formatted message in WARN level.
	Warnf(format string, args ...interface{})

	// Info must write a message in INFO level.
	Info(msg string, options ...OptionSetterFunc)

	// Infof must write a formatted message in INFO level.
	Infof(format string, args ...interface{})

	// Debug must write a message in DEBUG level.
	Debug(msg string, options ...OptionSetterFunc)

	// Debugf must write a formatted message in DEBUG level.
	Debugf(format string, args ...interface{})

	// NewChild must create a child logger and inherit level, writer and other flags
	// only option such as namespace could be overridden
	NewChild(args ...OptionSetterFunc) Logger
}

// log is a singleton logger instance
var log Logger
var logMutex sync.RWMutex

// Get retrieve logger instance and will fallback to StdLogger if no logger registered
func Get() Logger {
	// If log is nil, initiate standard logger
	if log == nil {
		// Get logger from env
		logLevelStr, _ := os.LookupEnv(EnvLogLevel)
		logLevel := ParseLevel(logLevelStr)

		// Get logger prefix
		logPrefix, _ := os.LookupEnv(EnvLogPrefix)

		// Init standard logger
		l := NewStdLogger(logLevel, os.Stderr, logPrefix, stdLog.LstdFlags)

		// Register logger
		Register(l)
		log.Debug("No logger found. StdLogger initiated")
	}
	return log
}

func NewChild(args ...OptionSetterFunc) Logger {
	// Get parent logger
	logger := Get()
	return logger.NewChild(args...)
}

// Register a logger implementation instance
func Register(l Logger) {
	// If logger is nil, return error
	if l == nil {
		panic(fmt.Errorf("%s: logger to be registered is nil", namespace))
	}

	// Set logger
	logMutex.Lock()
	defer logMutex.Unlock()
	log = l
}

// Clear logger implementation instance
func Clear() {
	// Set logger
	logMutex.Lock()
	defer logMutex.Unlock()
	log = nil
}

// Trace retrieve where the code is being called and returns full path of file where the error occurred
func Trace(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		file = "<???>"
		line = 0
	}
	return file, line
}

// ParseLevel parse level from string to Log Level enum
func ParseLevel(level string) LogLevel {
	switch strings.ToLower(level) {
	case "panic", "0", "fatal", "1":
		return LevelFatal
	case "warn", "4":
		return LevelWarn
	case "info", "6":
		return LevelInfo
	case "debug", "7":
		return LevelDebug
	default:
		return DefaultLevel
	}
}
