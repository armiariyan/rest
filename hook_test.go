package rest

import (
	"context"
	"testing"
)

// this test will do nothing except just call, since no operation happens inside the function
func TestNoopHookBeforeRequest(t *testing.T) {
	req := new(NoopHook)
	req.BeforeRequest(context.Background(), HookData{})
}

func TestNoopHookAfterRequest(t *testing.T) {
	req := new(NoopHook)
	req.AfterRequest(context.Background(), HookData{})
}
