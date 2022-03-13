package logOption

import (
	"context"
)

func AddMetadata(key string, val interface{}) SetterFunc {
	return func(o *Options) {
		if o.Metadata == nil {
			o.Metadata = make(map[string]interface{})
		}
		o.Metadata[key] = val
	}
}

func Metadata(m map[string]interface{}) SetterFunc {
	return func(o *Options) {
		o.Metadata = m
	}
}

func Format(args ...interface{}) SetterFunc {
	return func(o *Options) {
		o.FmtArgs = args
	}
}

func Error(err error) SetterFunc {
	return func(o *Options) {
		o.Values[ErrorKey] = err
	}
}

func WithNamespace(n string) SetterFunc {
	return func(o *Options) {
		o.Values[NamespaceKey] = n
	}
}

func Context(ctx context.Context) SetterFunc {
	return func(o *Options) {
		o.Context = ctx
	}
}
