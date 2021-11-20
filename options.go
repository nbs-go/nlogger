package nlogger

type Options struct {
	metadata map[string]interface{}
	err      error
	fmtArgs  []interface{}
}

var defaultOptions = &Options{
	metadata: nil,
	err:      nil,
	fmtArgs:  nil,
}

type SetOptionFn = func(*Options)

func evaluateOptions(opts []interface{}) *Options {
	if len(opts) == 0 {
		return nil
	}

	optCopy := &Options{}
	*optCopy = *defaultOptions
	for _, v := range opts {
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
		if o.metadata == nil {
			o.metadata = make(map[string]interface{})
		}
		o.metadata[key] = val
	}
}

func Metadata(m map[string]interface{}) SetOptionFn {
	return func(o *Options) {
		o.metadata = m
	}
}

func Format(args ...interface{}) SetOptionFn {
	return func(o *Options) {
		o.fmtArgs = args
	}
}

func Error(err error) SetOptionFn {
	return func(o *Options) {
		o.err = err
	}
}
