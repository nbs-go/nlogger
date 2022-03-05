package nlogger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	stdLog "log"
	"os"
)

var stdLevelPrefix = map[LogLevel]string{
	LevelFatal: "[FATAL] ",
	LevelError: "[ERROR] ",
	LevelWarn:  " [WARN] ",
	LevelInfo:  " [INFO] ",
	LevelDebug: "[DEBUG] ",
}

type StdPrinterFunc func(writer *stdLog.Logger, level LogLevel, msg string, options *Options, skipTrace int)

type StdLogger struct {
	level       LogLevel
	levelPrefix map[LogLevel]string
	skipTrace   int
	writer      *stdLog.Logger
	ioWriter    io.Writer
	namespace   string
	flags       int
	ctx         context.Context
	printerFn   StdPrinterFunc
}

func (l StdLogger) Fatal(msg string, args ...OptionSetterFunc) {
	l.print(LevelFatal, msg, EvaluateOptions(args))
}

func (l StdLogger) Fatalf(format string, args ...interface{}) {
	l.print(LevelFatal, format, NewFormatOptions(args...))
}

func (l StdLogger) Error(msg string, args ...OptionSetterFunc) {
	l.print(LevelError, msg, EvaluateOptions(args))
}

func (l StdLogger) Errorf(format string, args ...interface{}) {
	l.print(LevelError, format, NewFormatOptions(args...))
}

func (l StdLogger) Warn(msg string, args ...OptionSetterFunc) {
	l.print(LevelWarn, msg, EvaluateOptions(args))
}

func (l StdLogger) Warnf(format string, args ...interface{}) {
	l.print(LevelWarn, format, NewFormatOptions(args...))
}

func (l StdLogger) Info(msg string, args ...OptionSetterFunc) {
	l.print(LevelInfo, msg, EvaluateOptions(args))
}

func (l StdLogger) Infof(format string, args ...interface{}) {
	l.print(LevelInfo, format, NewFormatOptions(args...))
}

func (l *StdLogger) Debug(msg string, args ...OptionSetterFunc) {
	l.print(LevelDebug, msg, EvaluateOptions(args))
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.print(LevelDebug, format, NewFormatOptions(args...))
}

func (l *StdLogger) NewChild(args ...OptionSetterFunc) Logger {
	options := EvaluateOptions(args)

	// Override namespace if option is set
	n, _ := options.GetString(NamespaceKey)
	if n == "" {
		n = l.namespace
	}

	// Init logger
	cl := NewStdLogger(l.level, l.ioWriter, n, l.flags)

	// Set context if available
	ctx := options.GetContext()
	if ctx != nil {
		cl.ctx = ctx
	}

	return cl
}

func (l *StdLogger) print(outLevel LogLevel, msg string, options *Options) {
	// if output level is greater than log level, don't print
	if outLevel > l.level {
		return
	}

	var fn StdPrinterFunc
	if l.printerFn != nil {
		fn = l.printerFn
	} else {
		fn = stdPrint
	}

	// Inject context if not set
	if l.ctx != nil && !options.HasContext() {
		options.Values[ContextKey] = l.ctx
	}

	fn(l.writer, outLevel, msg, options, l.skipTrace)
}

func NewStdLogger(level LogLevel, w io.Writer, namespace string, flags int, args ...interface{}) *StdLogger {
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
		level:     level,
		skipTrace: 2,
		writer:    stdLog.New(w, prefix, flags),
		ioWriter:  w,
		namespace: namespace,
	}

	// Set optional arguments
	if len(args) > 0 {
		fn := args[0].(StdPrinterFunc)
		l.printerFn = fn
	}

	return &l
}

func stdPrint(writer *stdLog.Logger, level LogLevel, msg string, options *Options, skipTrace int) {
	// Generate prefix
	prefix := stdLevelPrefix[level]

	// If options is existed
	// If formatted arguments is available, then print as formatted
	fmtArgs := options.FmtArgs
	if len(fmtArgs) > 0 {
		writer.Printf(prefix+msg+"\n", fmtArgs...)
	} else {
		writer.Printf("%s%s\n", prefix, msg)
	}

	// Get context
	ctx := options.GetContext()
	if ctx != nil {
		// Get request id
		v := ctx.Value(RequestIdKey)
		reqId, ok := v.(string)
		if ok {
			writer.Printf("  > Request ID: %s\n", reqId)
		}
	}

	// If error exists, then print error and its trace
	logErr := options.GetError()
	if logErr != nil && level <= LevelError {
		filePath, line := Trace(skipTrace)
		writer.Printf("  > Error: %s\n", logErr)
		writer.Printf("  > Trace: %s:%d\n", filePath, line)
		// Print cause
		if unErr := errors.Unwrap(logErr); unErr != nil {
			writer.Printf("  > ErrorCause: %s\n", unErr)
		}
	}

	meta := options.Metadata
	if meta != nil && len(meta) > 0 {
		// Serialize to json
		metadata, err := json.Marshal(meta)
		// If not error, then print
		if err == nil {
			writer.Printf("  > Metadata: %s\n", metadata)
		}
	}
}
