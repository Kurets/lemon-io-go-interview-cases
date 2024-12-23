# Token Bucket Rate Limiter

## Task Description

Implement a **token bucket rate limiter** that allows a fixed number of requests per second. The rate limiter should:

1. Allow a specified number of requests (`rate`) per second.
2. Provide a method `Allow() bool` that:
    - Returns `true` if a request is allowed (a token is available).
    - Returns `false` if no tokens are available.

### Function Signatures

```go
package main

type RateLimiter struct {
   // Your code here
}

func NewRateLimiter(rate int) *RateLimiter {
   // Your code here
   return nil
}

func (r *RateLimiter) Allow() bool {
   // Your code here
   return false
}

```

### Requirements

1. **Token Bucket Behavior**:
    - The rate limiter should start with a full bucket of tokens (equal to `rate`).
    - Each second, the bucket should be refilled with tokens up to the maximum capacity (`rate`).
    - Requests can only proceed if tokens are available; each request consumes one token.

2. **Concurrency**:
    - Ensure thread-safe operation for concurrent calls to `Allow()`.

3. **Edge Cases**:
    - Handle scenarios where no requests are made for an extended period (tokens should still refill).
    - Ensure tokens do not exceed the maximum capacity (`rate`).

### Example Usage

```go
package main

import (
   "fmt"
   "time"
)

func main() {
   rl := NewRateLimiter(2)

   // Simulate requests
   for i := 0; i < 5; i++ {
      fmt.Println(rl.Allow()) // Prints true for the first two requests, then false
      time.Sleep(300 * time.Millisecond)
   }

   time.Sleep(1 * time.Second) // Wait for tokens to refill

   fmt.Println(rl.Allow()) // Prints true after refill
   fmt.Println(rl.Allow()) // Prints true
   fmt.Println(rl.Allow()) // Prints false (exhausted tokens)

   close(rl.stopChan) // Stop the refiller goroutine
}
```

### Notes

- Use a **ticker** to periodically refill the tokens.
- Protect shared data (e.g., token count) with a synchronization mechanism like `sync.Mutex`.
- Ensure proper cleanup of any goroutines when the rate limiter is no longer in use.

### Hints

- Use a `sync.Mutex` to synchronize access to the token count.
- Create a separate goroutine to handle token refills at fixed intervals.
- Use a `chan struct{}` or a similar mechanism to stop the refill goroutine gracefully.
