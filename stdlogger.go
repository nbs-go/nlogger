package nlogger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	stdLog "log"
	"os"
)

type StdLogger struct {
	level       LogLevel
	levelPrefix map[LogLevel]string
	skipTrace   int
	writer      *stdLog.Logger
}

func (l StdLogger) Fatal(msg string, args ...interface{}) {
	l.print(LevelFatal, msg, evaluateOptions(args))
}

func (l StdLogger) Fatalf(format string, args ...interface{}) {
	l.print(LevelFatal, format, &Options{
		fmtArgs: args,
	})
}

func (l StdLogger) Error(msg string, args ...interface{}) {
	l.print(LevelError, msg, evaluateOptions(args))
}

func (l StdLogger) Errorf(format string, args ...interface{}) {
	l.print(LevelError, format, &Options{
		fmtArgs: args,
	})
}

func (l StdLogger) Warn(msg string, args ...interface{}) {
	l.print(LevelWarn, msg, evaluateOptions(args))
}

func (l StdLogger) Warnf(format string, args ...interface{}) {
	l.print(LevelWarn, format, &Options{
		fmtArgs: args,
	})
}

func (l StdLogger) Info(msg string, args ...interface{}) {
	l.print(LevelInfo, msg, evaluateOptions(args))
}

func (l StdLogger) Infof(format string, args ...interface{}) {
	l.print(LevelInfo, format, &Options{
		fmtArgs: args,
	})
}

func (l *StdLogger) Debug(msg string, args ...interface{}) {
	l.print(LevelDebug, msg, evaluateOptions(args))
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.print(LevelDebug, format, &Options{
		fmtArgs: args,
	})
}

func NewStdLogger(level LogLevel, w io.Writer, prefix string, flags int) Logger {
	// If writer is nil, set default writer to Stdout
	if w == nil {
		w = os.Stdout
	}

	if prefix != "" {
		prefix = fmt.Sprintf("(%s) ", prefix)
	}

	// Init standard logger instance
	l := StdLogger{
		level: level,
		levelPrefix: map[LogLevel]string{
			LevelFatal: "[FATAL] ",
			LevelError: "[ERROR] ",
			LevelWarn:  " [WARN] ",
			LevelInfo:  " [INFO] ",
			LevelDebug: "[DEBUG] ",
		},
		skipTrace: 2,
		writer:    stdLog.New(w, prefix, flags),
	}
	return &l
}

func (l *StdLogger) print(outLevel LogLevel, msg string, options *Options) {
	// if output level is greater than log level, don't print
	if outLevel > l.level {
		return
	}

	// Generate prefix
	prefix := l.levelPrefix[outLevel]

	// If options is existed
	if options != nil {
		// If formatted arguments is available, then print as formatted
		if len(options.fmtArgs) > 0 {
			l.writer.Printf(prefix+msg+"\n", options.fmtArgs...)
		} else {
			l.writer.Printf("%s%s\n", prefix, msg)
		}

		// If error exists, then print error and its trace
		if options.err != nil && outLevel <= LevelError {
			filePath, line := Trace(l.skipTrace)
			l.writer.Printf("  > Error: %s\n", options.err)
			l.writer.Printf("  > Trace: %s:%d\n", filePath, line)
			// Print cause
			if unErr := errors.Unwrap(options.err); unErr != nil {
				l.writer.Printf("  > ErrorCause: %s\n", unErr)
			}
		}

		if options.metadata != nil && len(options.metadata) > 0 {
			// Serialize to json
			metadata, err := json.Marshal(options.metadata)
			// If not error, then print
			if err == nil {
				l.writer.Printf("  > metadata: %s\n", metadata)
			}
		}

		return
	}

	l.writer.Printf("%s%s\n", prefix, msg)
}
