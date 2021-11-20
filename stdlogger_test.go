package nlogger

import (
	"errors"
	"fmt"
	stdLog "log"
	"os"
	"testing"
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
			t.Log("Registering nil Logger is expected to panic")
		}
	}()

	// The following is the code under test
	Register(nil)
}
