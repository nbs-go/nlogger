package nlogger

import (
	"context"
	"errors"
	"fmt"
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
	testLogger := NewStdLogger(LevelDebug, nil, "test", stdLog.LstdFlags)
	testLogger.Fatal("Testing FATAL with message only")
	testLogger.Fatalf("Testing FATAL with formatted message: %s %s", "arg1", "arg2")
	testLogger.Fatal("Testing FATAL with options. Formatted Message: %s %s %s",
		Error(fmt.Errorf("a fatal error occurred. %w", errors.New("source of error"))),
		Metadata(metadata),
		AddMetadata("key", "value"),
		Format("arg1", "arg2", "arg3"),
	)
}

func TestError(t *testing.T) {
	testLogger := NewStdLogger(LevelDebug, nil, "test", stdLog.LstdFlags)
	testLogger.Error("Testing ERROR with message only")
	testLogger.Errorf("Testing ERROR with formatted message: %s %s", "arg1", "arg2")
	testLogger.Error("Testing ERROR with options. Formatted Message: %s %s %s",
		Error(errors.New("a error occurred")),
		Metadata(metadata),
		AddMetadata("key", "value"),
		Format("arg1", "arg2", "arg3"),
	)
}

func TestWarn(t *testing.T) {
	testLogger := NewStdLogger(LevelDebug, nil, "test", stdLog.LstdFlags)
	testLogger.Warn("Testing WARN with message only")
	testLogger.Warnf("Testing WARN with formatted message: %s %s", "arg1", "arg2")
	testLogger.Warn("Testing WARN with options. Formatted Message: %s %s %s",
		Metadata(metadata),
		AddMetadata("key", "value"),
		Format("arg1", "arg2", "arg3"))
}

func TestInfo(t *testing.T) {
	testLogger := NewStdLogger(LevelDebug, nil, "test", stdLog.LstdFlags)
	testLogger.Info("Testing INFO with message only")
	testLogger.Infof("Testing INFO with formatted message: %s %s", "arg1", "arg2")
	testLogger.Info("Testing INFO with options. Formatted Message: %s %s %s",
		Metadata(metadata),
		AddMetadata("key", "value"),
		Format("arg1", "arg2", "arg3"),
	)
}

func TestDebug(t *testing.T) {
	testLogger := NewStdLogger(LevelDebug, nil, "test", stdLog.LstdFlags)
	testLogger.Debug("Testing DEBUG with message only")
	testLogger.Debugf("Testing DEBUG with formatted message: %s %s", "arg1", "arg2")
	testLogger.Debug("Testing DEBUG with options. Formatted Message: %s %s %s",
		AddMetadata("key", "value"),
		Format("arg1", "arg2", "arg3"),
		"expected to skip",
	)
	testLogger.Debug("Testing DEBUG with error only", Error(errors.New("this is error")))
}

func TestParseLevel(t *testing.T) {
	testParseLevel(t, "0", LevelFatal)
	testParseLevel(t, "panic", LevelFatal)
	testParseLevel(t, "PANIC", LevelFatal)

	testParseLevel(t, "1", LevelFatal)
	testParseLevel(t, "fatal", LevelFatal)
	testParseLevel(t, "FATAL", LevelFatal)

	testParseLevel(t, "3", LevelError)
	testParseLevel(t, "error", LevelError)
	testParseLevel(t, "ERROR", LevelError)

	testParseLevel(t, "4", LevelWarn)
	testParseLevel(t, "warn", LevelWarn)
	testParseLevel(t, "WARN", LevelWarn)

	testParseLevel(t, "6", LevelInfo)
	testParseLevel(t, "info", LevelInfo)
	testParseLevel(t, "INFO", LevelInfo)

	testParseLevel(t, "7", LevelDebug)
	testParseLevel(t, "debug", LevelDebug)
	testParseLevel(t, "DEBUG", LevelDebug)
}

func TestTraceNotFound(t *testing.T) {
	file, line := Trace(3)
	if file != "<???>" {
		t.Errorf("unexpected traced file. Expected: <???>, Actual: %s", file)
	}
	if line != 0 {
		t.Errorf("unexpected traced line. Expected: 0, Actual: %d", line)
	}
}

func testParseLevel(t *testing.T, levelStr string, expectedLevel LogLevel) {
	level := ParseLevel(levelStr)
	if level != expectedLevel {
		t.Errorf("failed parsing level. Input: %s, Expected: %d, Actual: %d", levelStr, expectedLevel, level)
	}
}

func TestDefault(t *testing.T) {
	l := Get()
	l.Error("This is called from StdLogger")
	l.Debugf("This should not appear")
}

func TestRegisterEmptyLogger(t *testing.T) {
	// Clear existing logger
	Clear()

	// Ensure test will recover when panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Logf("Registering nil Logger is expected to panic. Error = %+v", r)
		}
	}()

	// The following is the code under test
	Register(nil)
}

func TestChildLogger(t *testing.T) {
	testLogger := NewStdLogger(LevelDebug, nil, "parent", stdLog.LstdFlags)
	testLogger.Debug("this is called from parent logger")

	childLogger1 := testLogger.NewChild(WithNamespace("child"))
	childLogger1.Debug("this is called from child logger with namespace")

	childLogger2 := NewChild()
	childLogger2.Debug("this is called from child logger without namespace")
}

func TestEvaluateOptions(t *testing.T) {
	// Evaluate context argument
	args := []interface{}{
		Context(context.Background()),
		AddMetadata("testInt64", 1),
	}
	o := EvaluateOptions(args)

	// Get context
	ctx := o.GetContext()
	if ctx == nil {
		t.Error("Context is not evaluated")
	} else {
		t.Log("Context is evaluated and set to options")
	}
}

func TestCustomOptions(t *testing.T) {
	dt := time.Now()
	o := &Options{
		KV: map[string]interface{}{
			"testInt64":  int64(99),
			"testInt":    98,
			"testString": "Hello",
			"testTime":   dt,
		},
	}

	// Check string
	str, ok := o.GetString("testString")
	if !ok || str != "Hello" {
		t.Errorf("Unexpected str value from option. Got value = %s", str)
	}

	// Check int
	i, ok := o.GetInt("testInt")
	if !ok || i != 98 {
		t.Errorf("Unexpected int64 value from option. Got value = %d", i)
	}

	// Check int64
	i64, ok := o.GetInt64("testInt64")
	if !ok || i64 != 99 {
		t.Errorf("Unexpected int64 value from option. Got value = %d", i)
	}

	// Check time
	dtt, ok := o.GetTime("testTime")
	if !ok || dtt.Unix() != dt.Unix() {
		t.Errorf("Unexpected int64 value from option. Got value = %s", dtt)
	}

	// Get empty int
	_, ok = o.GetInt("noInt")
	if ok {
		t.Errorf("Unexpected value. int Value is supposed not to be set")
	}

	// Get empty int64
	_, ok = o.GetInt64("noInt64")
	if ok {
		t.Errorf("Unexpected value. int64 Value is supposed not to be set")
	}

	// Get empty string
	_, ok = o.GetString("noString")
	if ok {
		t.Errorf("Unexpected value. string Value is supposed not to be set")
	}

	// Get empty time
	_, ok = o.GetTime("noTime")
	if ok {
		t.Errorf("Unexpected value. time.Time Value is supposed not to be set")
	}

	// Get empty context
	ctx := o.GetContext()
	if ctx != nil {
		t.Errorf("Unexpected value. context Value is supposed not to be set")
	}
}
