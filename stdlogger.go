package nlogger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nbs-go/nlogger/v2/level"
	"github.com/nbs-go/nlogger/v2/option"
	"io"
	stdLog "log"
	"os"
)

var stdLevelPrefix = map[level.LogLevel]string{
	level.Fatal: "[FATAL] ",
	level.Error: "[ERROR] ",
	level.Warn:  " [WARN] ",
	level.Info:  " [INFO] ",
	level.Debug: "[DEBUG] ",
}

type StdPrinterFunc func(writer *stdLog.Logger, lv level.LogLevel, msg string, options *logOption.Options)

type StdLogger struct {
	level     level.LogLevel
	writer    *stdLog.Logger
	ioWriter  io.Writer
	namespace string
	flags     int
	ctx       context.Context
	printerFn StdPrinterFunc
}

func (l *StdLogger) Fatal(msg string, args ...logOption.SetterFunc) {
	l.print(level.Fatal, msg, logOption.Evaluate(args))
}

func (l *StdLogger) Fatalf(format string, args ...interface{}) {
	l.print(level.Fatal, format, logOption.NewFormatOptions(args...))
}

func (l *StdLogger) Error(msg string, args ...logOption.SetterFunc) {
	l.print(level.Error, msg, logOption.Evaluate(args))
}

func (l *StdLogger) Errorf(format string, args ...interface{}) {
	l.print(level.Error, format, logOption.NewFormatOptions(args...))
}

func (l *StdLogger) Warn(msg string, args ...logOption.SetterFunc) {
	l.print(level.Warn, msg, logOption.Evaluate(args))
}

func (l *StdLogger) Warnf(format string, args ...interface{}) {
	l.print(level.Warn, format, logOption.NewFormatOptions(args...))
}

func (l *StdLogger) Info(msg string, args ...logOption.SetterFunc) {
	l.print(level.Info, msg, logOption.Evaluate(args))
}

func (l *StdLogger) Infof(format string, args ...interface{}) {
	l.print(level.Info, format, logOption.NewFormatOptions(args...))
}

func (l *StdLogger) Debug(msg string, args ...logOption.SetterFunc) {
	l.print(level.Debug, msg, logOption.Evaluate(args))
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.print(level.Debug, format, logOption.NewFormatOptions(args...))
}

func (l *StdLogger) NewChild(args ...logOption.SetterFunc) Logger {
	options := logOption.Evaluate(args)

	// Override namespace if option is set
	n, _ := logOption.GetString(options, logOption.NamespaceKey)
	if n == "" {
		n = l.namespace
	}

	// Init logger
	cl := NewStdLogger(l.level, l.ioWriter, n, l.flags)

	// Set context if available
	if ctx := options.Context; ctx != nil {
		cl.ctx = ctx
	}

	return cl
}

func (l *StdLogger) print(outLevel level.LogLevel, msg string, options *logOption.Options) {
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
	if l.ctx != nil && options.Context == nil {
		options.Context = l.ctx
	}

	fn(l.writer, outLevel, msg, options)
}

func NewStdLogger(level level.LogLevel, w io.Writer, namespace string, flags int, args ...interface{}) *StdLogger {
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

func stdPrint(writer *stdLog.Logger, lv level.LogLevel, msg string, options *logOption.Options) {
	// Generate prefix
	prefix := stdLevelPrefix[lv]

	// If options is existed
	// If formatted arguments is available, then print as formatted
	fmtArgs := options.FmtArgs
	if len(fmtArgs) > 0 {
		writer.Printf(prefix+msg+"\n", fmtArgs...)
	} else {
		writer.Printf("%s%s\n", prefix, msg)
	}

	// Get context
	if ctx := options.Context; ctx != nil {
		// Get request id
		v := ctx.Value(RequestIdKey)
		reqId, ok := v.(string)
		if ok {
			writer.Printf("  > Request ID: %s\n", reqId)
		}
	}

	// If error exists, then print error and its trace
	logErr := logOption.GetError(options, logOption.ErrorKey)
	if logErr != nil && lv <= level.Error {
		writer.Printf("  > Error: %s\n", logErr)
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
