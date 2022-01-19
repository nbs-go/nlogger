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
	ioWriter    io.Writer
	namespace   string
	flags       int
}

func (l StdLogger) Fatal(msg string, args ...interface{}) {
	l.print(LevelFatal, msg, EvaluateOptions(args))
}

func (l StdLogger) Fatalf(format string, args ...interface{}) {
	l.print(LevelFatal, format, NewFormatOptions(args))
}

func (l StdLogger) Error(msg string, args ...interface{}) {
	l.print(LevelError, msg, EvaluateOptions(args))
}

func (l StdLogger) Errorf(format string, args ...interface{}) {
	l.print(LevelError, format, NewFormatOptions(args))
}

func (l StdLogger) Warn(msg string, args ...interface{}) {
	l.print(LevelWarn, msg, EvaluateOptions(args))
}

func (l StdLogger) Warnf(format string, args ...interface{}) {
	l.print(LevelWarn, format, NewFormatOptions(args))
}

func (l StdLogger) Info(msg string, args ...interface{}) {
	l.print(LevelInfo, msg, EvaluateOptions(args))
}

func (l StdLogger) Infof(format string, args ...interface{}) {
	l.print(LevelInfo, format, NewFormatOptions(args))
}

func (l *StdLogger) Debug(msg string, args ...interface{}) {
	l.print(LevelDebug, msg, EvaluateOptions(args))
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.print(LevelDebug, format, NewFormatOptions(args))
}

func (l *StdLogger) NewChild(args ...interface{}) Logger {
	options := EvaluateOptions(args)

	// Override namespace if option is set
	n, _ := options.GetString(NamespaceKey)
	if n == "" {
		n = l.namespace
	}

	return NewStdLogger(l.level, l.ioWriter, n, l.flags)
}

func NewStdLogger(level LogLevel, w io.Writer, namespace string, flags int) Logger {
	// If writer is nil, set default writer to Stdout
	if w == nil {
		w = os.Stdout
	}

	var prefix string
	if namespace != "" {
		prefix = fmt.Sprintf("(%s) ", namespace)
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
		ioWriter:  w,
		namespace: namespace,
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
	// If formatted arguments is available, then print as formatted
	fmtArgs := options.GetFmtArgs()
	if len(fmtArgs) > 0 {
		l.writer.Printf(prefix+msg+"\n", fmtArgs...)
	} else {
		l.writer.Printf("%s%s\n", prefix, msg)
	}

	// If error exists, then print error and its trace
	logErr := options.GetError()
	if logErr != nil && outLevel <= LevelError {
		filePath, line := Trace(l.skipTrace)
		l.writer.Printf("  > Error: %s\n", logErr)
		l.writer.Printf("  > Trace: %s:%d\n", filePath, line)
		// Print caus
		if unErr := errors.Unwrap(logErr); unErr != nil {
			l.writer.Printf("  > ErrorCause: %s\n", unErr)
		}
	}

	meta := options.Metadata
	if meta != nil && len(meta) > 0 {
		// Serialize to json
		metadata, err := json.Marshal(meta)
		// If not error, then print
		if err == nil {
			l.writer.Printf("  > metadata: %s\n", metadata)
		}
	}
}
