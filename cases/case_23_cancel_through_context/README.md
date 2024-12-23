# Cancel through Context

## Task Description

You need to write a function:

```go
package main

import "context"

func longOperation(ctx context.Context) {
	// Your code here
}
```

### Requirements

1. **Long Operation**
    - The function simulates a time-consuming (long-running) task.
    - It should run asynchronously if needed, but at least it must respond to context cancellation.

2. **Use `context.Context`**
    - Honor the `ctx.Done()` channel to detect cancellation or timeout.
    - Terminate the operation as soon as `ctx.Done()` is signaled, rather than continuing to completion.

3. **Behavior**
    - If the context is **not** canceled, the function should simulate completing its work.
    - If the context **is** canceled or times out, the function should stop early and return promptly.

4. **Edge Cases**
    - **Immediate cancellation**: If the context is canceled before starting, your function should return quickly.
    - **Long running**: The function might have loops or periodic checks for `ctx.Err()`.

### Example Usage

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Example: Start a long operation with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	longOperation(ctx)
	fmt.Println("Operation done (either completed or canceled)")
}
```

### Hints

- A typical pattern is something like:
```go
for {
	// do partial work
	select {
	case <-ctx.Done():
		// context canceled or timed out
		return
	default:
		// continue
	}
}
```
- If your operation does smaller sub-steps or sleeps, just ensure each iteration checks `ctx.Done()`.
- Use `context.WithCancel`, `context.WithTimeout`, or `context.WithDeadline` in tests to confirm cancellation/timeout.
- Make sure to exit cleanly and not block if the context is canceled.  
