# Concurrency-Safe Memoization

## Task Description

Implement a **concurrency-safe memoization function**. Given a function `f(key string) string`, return a new function
that:

1. Caches results for each key, so subsequent calls with the same key return the cached result.
2. Ensures thread safety for concurrent calls to the memoized function.

### Function Signature

```go
package main

func Memoize(f func(string) string) func(string) string {
   // Your code here
   return nil
}
```

### Requirements

1. **Caching Behavior**:
    - The returned function should store results for previously computed keys.
    - If a key is already cached, return the cached result instead of recomputing.

2. **Concurrency**:
    - Ensure the memoization is thread-safe, allowing multiple goroutines to access the memoized function simultaneously
      without issues.

3. **Correctness**:
    - Handle multiple unique keys, ensuring each key is computed and cached correctly.
    - Avoid redundant calls to the underlying function `f` for cached keys.

### Example Usage

```go
package main

import (
   "fmt"
)

func main() {
   slowFunc := func(s string) string {
      // Simulate slow computation
      return s + "-computed"
   }

   memoized := Memoize(slowFunc)

   // First call computes and caches the result
   fmt.Println(memoized("test")) // Output: "test-computed"

   // Second call retrieves the cached result
   fmt.Println(memoized("test")) // Output: "test-computed"
}
```

### Notes

- Use a **mutex** (e.g., `sync.Mutex`) to synchronize access to the cache.
- Store cached results in a `map[string]string`.
- Ensure that the function behaves correctly for both single-threaded and multi-threaded scenarios.

### Hints

- Before computing the result for a key, check if it exists in the cache.
- Lock the cache for reading and writing to ensure thread safety.
- Avoid holding the lock while executing the underlying function `f` to prevent deadlocks or reduced concurrency.
