package nlogger

import (
	"context"
	"time"
)

type Options struct {
	KV       map[string]interface{}
	Metadata map[string]interface{}
}

type SetOptionFn = func(*Options)

// Constructors

func NewOptions() *Options {
	return &Options{KV: make(map[string]interface{})}
}

func NewFormatOptions(args ...interface{}) *Options {
	return &Options{
		KV: map[string]interface{}{
			FormatArgsKey: args,
		},
	}
}

// Options instance methods

func (o *Options) GetString(k string) (string, bool) {
	v, ok := o.KV[k]
	if !ok {
		return "", false
	}

	s, ok := v.(string)
	return s, ok
}

func (o *Options) GetInt64(k string) (int64, bool) {
	v, ok := o.KV[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int64)
	return i, ok
}

func (o *Options) GetInt(k string) (int, bool) {
	v, ok := o.KV[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int)
	return i, ok
}

func (o *Options) GetTime(k string) (time.Time, bool) {
	v, ok := o.KV[k]
	if !ok {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}

func (o *Options) GetContext() context.Context {
	v, ok := o.KV[ContextKey]
	if !ok {
		return nil
	}
	return v.(context.Context)
}

func (o *Options) GetFmtArgs() []interface{} {
	v, ok := o.KV[FormatArgsKey]
	if !ok {
		return nil
	}

	return v.([]interface{})
}

func (o *Options) GetError() error {
	v, ok := o.KV[ErrorKey]
	if !ok {
		return nil
	}

	return v.(error)
}

// Utils

func EvaluateOptions(args []interface{}) *Options {
	optCopy := NewOptions()
	for _, v := range args {
		fn, ok := v.(SetOptionFn)
		if !ok {
			// Skipping
			continue
		}
		fn(optCopy)
	}
	return optCopy
}

func AddMetadata(key string, val interface{}) SetOptionFn {
	return func(o *Options) {
		if o.Metadata == nil {
			o.Metadata = make(map[string]interface{})
		}
		o.Metadata[key] = val
	}
}

func Metadata(m map[string]interface{}) SetOptionFn {
	return func(o *Options) {
		o.Metadata = m
	}
}

func Format(args ...interface{}) SetOptionFn {
	return func(o *Options) {
		o.KV[FormatArgsKey] = args
	}
}

func Error(err error) SetOptionFn {
	return func(o *Options) {
		o.KV[ErrorKey] = err
	}
}

func WithNamespace(n string) SetOptionFn {
	return func(o *Options) {
		o.KV[NamespaceKey] = n
	}
}

func Context(ctx context.Context) SetOptionFn {
	return func(o *Options) {
		o.KV[ContextKey] = ctx
	}
}
