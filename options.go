package nlogger

import (
	"context"
	"time"
)

type Options struct {
	Values   map[string]interface{}
	Metadata map[string]interface{}
	FmtArgs  []interface{}
	Context  context.Context
}

type OptionSetterFunc = func(*Options)

// Constructors

func NewOptions() *Options {
	return &Options{Values: make(map[string]interface{})}
}

func NewFormatOptions(args ...interface{}) *Options {
	return &Options{
		Values:  make(map[string]interface{}),
		FmtArgs: args,
	}
}

// Options instance methods

func (o *Options) GetString(k string) (string, bool) {
	v, ok := o.Values[k]
	if !ok {
		return "", false
	}

	s, ok := v.(string)
	return s, ok
}

func (o *Options) GetInt64(k string) (int64, bool) {
	v, ok := o.Values[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int64)
	return i, ok
}

func (o *Options) GetInt(k string) (int, bool) {
	v, ok := o.Values[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int)
	return i, ok
}

func (o *Options) GetTime(k string) (time.Time, bool) {
	v, ok := o.Values[k]
	if !ok {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}

func (o *Options) GetError() error {
	v, ok := o.Values[ErrorKey]
	if !ok {
		return nil
	}

	return v.(error)
}

// Utils

func EvaluateOptions(args []OptionSetterFunc) *Options {
	optCopy := NewOptions()
	for _, fn := range args {
		fn(optCopy)
	}
	return optCopy
}

func AddMetadata(key string, val interface{}) OptionSetterFunc {
	return func(o *Options) {
		if o.Metadata == nil {
			o.Metadata = make(map[string]interface{})
		}
		o.Metadata[key] = val
	}
}

func Metadata(m map[string]interface{}) OptionSetterFunc {
	return func(o *Options) {
		o.Metadata = m
	}
}

func Format(args ...interface{}) OptionSetterFunc {
	return func(o *Options) {
		o.FmtArgs = args
	}
}

func Error(err error) OptionSetterFunc {
	return func(o *Options) {
		o.Values[ErrorKey] = err
	}
}

func WithNamespace(n string) OptionSetterFunc {
	return func(o *Options) {
		o.Values[NamespaceKey] = n
	}
}

func Context(ctx context.Context) OptionSetterFunc {
	return func(o *Options) {
		o.Context = ctx
	}
}
