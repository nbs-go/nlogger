package nlogger

const pkgNamespace = "nlogger"

// Configuration constants.
const (
	EnvLogLevel  = "LOG_LEVEL"
	EnvLogPrefix = "LOG_PREFIX"
)

// Context Key

type ContextKey string

const (
	RequestIdKey ContextKey = "requestId"
)
