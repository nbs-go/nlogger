package nlogger_test

import (
	"context"
	logContext "github.com/nbs-go/nlogger/v2/context"
	"testing"
)

func TestSetRequestId(t *testing.T) {
	ctx := logContext.SetRequestId(nil, "1")
	if ctx != nil {
		t.Errorf("unexpected context not nil")
		return
	}

	ctx = logContext.SetRequestId(context.Background(), "")
	if reqId := logContext.GetRequestId(ctx); reqId != "" {
		t.Errorf("unexpected request id is not empty = %s", reqId)
	}
}
