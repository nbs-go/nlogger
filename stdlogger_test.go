package nlogger_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nbs-go/nlogger/v2"
	"github.com/nbs-go/nlogger/v2/level"
	logOption "github.com/nbs-go/nlogger/v2/option"
	stdLog "log"
	"os"
	"testing"
	"time"
)

// Init sample variables
var metadata = map[string]interface{}{
	"string":  "string",
	"integer": 0,
	"boolean": true,
	"array":   []int{1, 2, 3, 4, 5},
	"struct": struct {
		Text    string
		Boolean bool
	}{
		Text:    "text",
		Boolean: false,
	},
}

func TestMain(m *testing.M) {
	// Run Test
	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

func TestFatal(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "test", stdLog.LstdFlags)
	testLogger.Fatal("Testing FATAL with message only")
	testLogger.Fatalf("Testing FATAL with formatted message: %s %s", "arg1", "arg2")
	testLogger.Fatal("Testing FATAL with options. Formatted Message: %s %s %s",
		logOption.Error(fmt.Errorf("a fatal error occurred. CausedBy => %w", errors.New("source of error"))),
		logOption.Metadata(metadata),
		logOption.AddMetadata("key", "value"),
		logOption.Format("arg1", "arg2", "arg3"),
	)
}

func TestError(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "test", stdLog.LstdFlags)
	testLogger.Error("Testing ERROR with message only")
	testLogger.Errorf("Testing ERROR with formatted message: %s %s", "arg1", "arg2")
	testLogger.Error("Testing ERROR with options. Formatted Message: %s %s %s",
		logOption.Error(errors.New("a error occurred")),
		logOption.Metadata(metadata),
		logOption.AddMetadata("key", "value"),
		logOption.Format("arg1", "arg2", "arg3"),
	)
}

func TestWarn(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "test", stdLog.LstdFlags)
	testLogger.Warn("Testing WARN with message only")
	testLogger.Warnf("Testing WARN with formatted message: %s %s", "arg1", "arg2")
	testLogger.Warn("Testing WARN with options. Formatted Message: %s %s %s",
		logOption.Metadata(metadata),
		logOption.AddMetadata("key", "value"),
		logOption.Format("arg1", "arg2", "arg3"))
}

func TestInfo(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "test", stdLog.LstdFlags)
	testLogger.Info("Testing INFO with message only")
	testLogger.Infof("Testing INFO with formatted message: %s %s", "arg1", "arg2")
	testLogger.Info("Testing INFO with options. Formatted Message: %s %s %s",
		logOption.Metadata(metadata),
		logOption.AddMetadata("key", "value"),
		logOption.Format("arg1", "arg2", "arg3"),
	)
}

func TestDebug(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "test", stdLog.LstdFlags)
	testLogger.Debug("Testing DEBUG with message only")
	testLogger.Debugf("Testing DEBUG with formatted message: %s %s", "arg1", "arg2")
	testLogger.Debug("Testing DEBUG with options. Formatted Message: %s %s %s",
		logOption.AddMetadata("key", "value"),
		logOption.Format("arg1", "arg2", "arg3"),
	)
	testLogger.Debug("Testing DEBUG with error only",
		logOption.Error(errors.New("this is error")))
}

func TestParseLevel(t *testing.T) {
	testParseLevel(t, "0", level.Fatal)
	testParseLevel(t, "panic", level.Fatal)
	testParseLevel(t, "PANIC", level.Fatal)

	testParseLevel(t, "1", level.Fatal)
	testParseLevel(t, "fatal", level.Fatal)
	testParseLevel(t, "FATAL", level.Fatal)

	testParseLevel(t, "3", level.Error)
	testParseLevel(t, "error", level.Error)
	testParseLevel(t, "ERROR", level.Error)

	testParseLevel(t, "4", level.Warn)
	testParseLevel(t, "warn", level.Warn)
	testParseLevel(t, "WARN", level.Warn)

	testParseLevel(t, "6", level.Info)
	testParseLevel(t, "info", level.Info)
	testParseLevel(t, "INFO", level.Info)

	testParseLevel(t, "7", level.Debug)
	testParseLevel(t, "debug", level.Debug)
	testParseLevel(t, "DEBUG", level.Debug)
}

func testParseLevel(t *testing.T, levelStr string, expectedLevel level.LogLevel) {
	lv := level.Parse(levelStr)
	if lv != expectedLevel {
		t.Errorf("failed parsing level. Input: %s, Expected: %d, Actual: %d", levelStr, expectedLevel, lv)
	}
}

func TestDefault(t *testing.T) {
	l := nlogger.Get()
	l.Error("This is called from StdLogger")
	l.Debugf("This should not appear")
}

func TestRegisterEmptyLogger(t *testing.T) {
	// Clear existing logger
	nlogger.Clear()

	// Ensure test will recover when panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Logf("Registering nil Logger is expected to panic. Error = %+v", r)
		}
	}()

	// The following is the code under test
	nlogger.Register(nil)
}

func TestChildLogger(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "parent", stdLog.LstdFlags)
	testLogger.Debug("this is called from parent logger")

	childLogger1 := testLogger.NewChild(logOption.WithNamespace("child"))
	childLogger1.Debug("this is called from child logger with namespace")

	childLogger2 := nlogger.NewChild()
	childLogger2.Debug("this is called from child logger without namespace")

	ctx := context.Background()
	ctx = context.WithValue(ctx, nlogger.RequestIdKey, "b0a495f4-f919-4fc0-b3e2-95f83d0c4a04")
	ctxLogger := testLogger.NewChild(logOption.Context(ctx))
	ctxLogger.Debugf("this log must contains request id")
}

func TestEvaluateOptions(t *testing.T) {
	// Evaluate context argument
	args := []logOption.SetterFunc{
		logOption.Context(context.Background()),
		logOption.AddMetadata("testInt64", 1),
	}
	o := logOption.Evaluate(args)

	// Get context
	ctx := o.Context
	if ctx == nil {
		t.Error("Context is not evaluated")
	} else {
		t.Log("Context is evaluated and set to options")
	}
}

func TestCustomOptions(t *testing.T) {
	dt := time.Now()
	o := &logOption.Options{
		Values: map[string]interface{}{
			"testInt64":  int64(99),
			"testInt":    98,
			"testString": "Hello",
			"testTime":   dt,
		},
	}

	// Check string
	str, ok := logOption.GetString(o, "testString")
	if !ok || str != "Hello" {
		t.Errorf("Unexpected str value from option. Got value = %s", str)
	}

	// Check int
	i, ok := logOption.GetInt(o, "testInt")
	if !ok || i != 98 {
		t.Errorf("Unexpected int64 value from option. Got value = %d", i)
	}

	// Check int64
	i64, ok := logOption.GetInt64(o, "testInt64")
	if !ok || i64 != 99 {
		t.Errorf("Unexpected int64 value from option. Got value = %d", i)
	}

	// Check time
	dtt, ok := logOption.GetTime(o, "testTime")
	if !ok || dtt.Unix() != dt.Unix() {
		t.Errorf("Unexpected int64 value from option. Got value = %s", dtt)
	}

	// Get empty int
	_, ok = logOption.GetInt(o, "noInt")
	if ok {
		t.Errorf("Unexpected value. int Value is supposed not to be set")
	}

	// Get empty int64
	_, ok = logOption.GetInt64(o, "noInt64")
	if ok {
		t.Errorf("Unexpected value. int64 Value is supposed not to be set")
	}

	// Get empty string
	_, ok = logOption.GetString(o, "noString")
	if ok {
		t.Errorf("Unexpected value. string Value is supposed not to be set")
	}

	// Get empty time
	_, ok = logOption.GetTime(o, "noTime")
	if ok {
		t.Errorf("Unexpected value. time.Time Value is supposed not to be set")
	}

	// Get empty context
	ctx := o.Context
	if ctx != nil {
		t.Errorf("Unexpected value. context Value is supposed not to be set")
	}
}

func TestCustomPrinter(t *testing.T) {
	testLogger := nlogger.NewStdLogger(level.Debug, nil, "", 0, newCustomPrinter())
	testLogger.Debug("This message is printed in json")
}

func newCustomPrinter() nlogger.StdPrinterFunc {
	return func(writer *stdLog.Logger, outLevel level.LogLevel, msg string, options *logOption.Options) {
		// Init json body
		jsonBody := map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     outLevel,
		}

		// Compose message
		if len(options.FmtArgs) > 0 {
			jsonBody["message"] = fmt.Sprintf(msg, options.FmtArgs...)
		} else {
			jsonBody["message"] = msg
		}

		// Compose json string
		jsonStr, _ := json.Marshal(jsonBody)
		writer.Printf("%s\n", jsonStr)
	}
}
