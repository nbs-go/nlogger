package nlogger

import (
	"github.com/nbs-go/nlogger/v2/level"
	logOption "github.com/nbs-go/nlogger/v2/option"
)

// Printer defines interface that are able to print a log message
type Printer interface {
	Print(outLevel level.LogLevel, msg string, options *logOption.Options)
}
