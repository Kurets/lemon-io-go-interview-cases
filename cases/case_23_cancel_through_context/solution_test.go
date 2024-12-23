package case_23_cancel_through_context

import (
	"context"
	"testing"
	"time"
)

const maxWaitTime = 2*time.Second + 5*time.Millisecond

func TestLongOperationCompletion(t *testing.T) {
	ctx := context.Background()

	done := make(chan struct{})
	go func() {
		longOperation(ctx)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(maxWaitTime):
		t.Error("longOperation did not complete in the expected time")
	}
}

func TestLongOperationCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		longOperation(ctx)
		close(done)
	}()

	time.Sleep(500 * time.Millisecond)
	cancel()

	select {
	case <-done:
	case <-time.After(maxWaitTime):
		t.Error("longOperation did not cancel after context cancellation")
	}
}

func TestLongOperationTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		longOperation(ctx)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(maxWaitTime):
		t.Error("longOperation did not timeout as expected")
	}
}
