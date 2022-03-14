package logContext

import "context"

type ContextKey string

const (
	RequestIdKey ContextKey = "requestId"
)

// SetRequestId is helper function to set request id value to context
func SetRequestId(ctx context.Context, reqId string) context.Context {
	if ctx == nil || reqId == "" {
		return ctx
	}
	return context.WithValue(ctx, RequestIdKey, reqId)
}

// GetRequestId is helper function to retrieve request id value in context
func GetRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// Get request id
	v := ctx.Value(RequestIdKey)
	switch str := v.(type) {
	case string:
		return str
	default:
		return ""
	}
}
