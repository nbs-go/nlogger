package nlogger

const pkgNamespace = "nlogger"

// LogLevel constants as defined in RFC5424.
type LogLevel = int8

const (
	_ LogLevel = iota // LevelPanic
	LevelFatal
	_ // LevelCritical
	LevelError
	LevelWarn
	_ // LevelNotice
	LevelInfo
	LevelDebug
	_ // LevelTrace
)

// Configuration constants.
const (
	EnvLogLevel  = "LOG_LEVEL"
	EnvLogPrefix = "LOG_PREFIX"

	DefaultLevel = LevelError
)

// Option keys constants
const (
	ErrorKey     = "error"
	NamespaceKey = "namespace"
)

// Context Key

type ContextKey string

const (
	RequestIdKey ContextKey = "requestId"
)
