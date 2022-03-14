package nlogger_test

import (
	"github.com/nbs-go/nlogger/v2/level"
	"testing"
)

func TestString_Fatal(t *testing.T) {
	exp := "Fatal"
	str := level.String(level.Fatal)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}

func TestString_Error(t *testing.T) {
	exp := "Error"
	str := level.String(level.Error)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}

func TestString_Warn(t *testing.T) {
	exp := "Warn"
	str := level.String(level.Warn)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}

func TestString_Info(t *testing.T) {
	exp := "Info"
	str := level.String(level.Info)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}

func TestString_Debug(t *testing.T) {
	exp := "Debug"
	str := level.String(level.Debug)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}

func TestString_Unknown(t *testing.T) {
	exp := "Unknown"
	str := level.String(-1)
	if str != exp {
		t.Errorf("unexpected %s string value = %s", exp, str)
	}
}
