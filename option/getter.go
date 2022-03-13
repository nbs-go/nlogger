package logOption

import "time"

// GetString is helper to retrieve string value in Values by key
// Value type must be exact, as it use casting instead of converting to target value
func GetString(o *Options, k string) (string, bool) {
	v, ok := o.Values[k]
	if !ok {
		return "", false
	}

	s, ok := v.(string)
	return s, ok
}

// GetInt64 is helper to retrieve int64 value in Values by key.
// Value type must be exact, as it use casting instead of converting to target value
func GetInt64(o *Options, k string) (int64, bool) {
	v, ok := o.Values[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int64)
	return i, ok
}

// GetInt is helper to retrieve int value in Values by key
// Value type must be exact, as it use casting instead of converting to target value
func GetInt(o *Options, k string) (int, bool) {
	v, ok := o.Values[k]
	if !ok {
		return 0, false
	}
	i, ok := v.(int)
	return i, ok
}

// GetTime is helper to retrieve time.Time value in Values by key
// Value type must be exact, as it use casting instead of converting to target value
func GetTime(o *Options, k string) (time.Time, bool) {
	v, ok := o.Values[k]
	if !ok {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}

// GetError is helper to retrieve error value in Values by key
// Value type must be exact, as it use casting instead of converting to target value
func GetError(o *Options, k string) error {
	v, ok := o.Values[k]
	if !ok {
		return nil
	}
	return v.(error)
}
