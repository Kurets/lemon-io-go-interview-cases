package case_23_cancel_through_context

import (
	"context"
	"fmt"
	"time"
)

func longOperation(ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("operation canceled")
			return
		default:
			// simulate work
			time.Sleep(200 * time.Millisecond)
			fmt.Println("step", i)
		}
	}
	fmt.Println("operation completed")
}
