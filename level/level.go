package level

import "strings"

// LogLevel constants as defined in RFC5424.
type LogLevel = int8

const (
	_ LogLevel = iota // LevelPanic
	Fatal
	_ // LevelCritical
	Error
	Warn
	_ // LevelNotice
	Info
	Debug
	Trace // LevelTrace
)

const (
	Default = Error
)

// Parse parse string value to level.LogLevel
func Parse(level string) LogLevel {
	switch strings.ToLower(level) {
	case "panic", "0", "fatal", "1":
		return Fatal
	case "warn", "4":
		return Warn
	case "info", "6":
		return Info
	case "debug", "7":
		return Debug
	case "trace", "8":
		return Trace
	default:
		return Default
	}
}

func String(l LogLevel) string {
	switch l {
	case Fatal:
		return "Fatal"
	case Error:
		return "Error"
	case Warn:
		return "Warn"
	case Info:
		return "Info"
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	}
	return "Unknown"
}
