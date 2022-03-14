package logOption

import (
	"context"
	"github.com/nbs-go/nlogger/v2/level"
)

type Options struct {
	Values   map[string]interface{}
	Metadata map[string]interface{}
	FmtArgs  []interface{}
	Context  context.Context
	Level    level.LogLevel
}

type SetterFunc = func(*Options)

// NewOptions construct options
func NewOptions() *Options {
	return &Options{
		Values: make(map[string]interface{}),
		Level:  level.Default,
	}
}

// NewFormatOptions construct options for formatting
func NewFormatOptions(args ...interface{}) *Options {
	return &Options{
		Values:  make(map[string]interface{}),
		FmtArgs: args,
	}
}

// Evaluate initiate given option setter that is set in args parameter and returns Options
func Evaluate(args []SetterFunc) *Options {
	optCopy := NewOptions()
	for _, fn := range args {
		fn(optCopy)
	}
	return optCopy
}
