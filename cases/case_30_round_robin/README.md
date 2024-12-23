# Round-Robin Load Balancer

## Task Description

Implement a **round-robin** load balancer with a list of server endpoints. Each time `Next()` is called,
return the **next** endpoint in the list, cycling back to the first endpoint after reaching the end.

```go
package main

type RoundRobin struct {
	// Your code here (e.g. slice of endpoints, current index, etc.)
}

func NewRoundRobin(endpoints []string) *RoundRobin {
	// Your code here
	return nil
}

func (r *RoundRobin) Next() string {
	// Your code here
	return ""
}
```

### Requirements

1. **Endpoints**
    - A slice of server endpoints, e.g. `[]string{"server1", "server2", "server3"}`.
    - If the slice is empty, decide how to handle `Next()` (return an empty string, or handle differently).

2. **Round-Robin Behavior**
    - Each call to `Next()` should return the next endpoint in a cycle.
    - After the last endpoint, wrap around to the first.

3. **Thread Safety**
    - (Optional) If concurrency is needed, you might protect the current index with a mutex.
    - The tests here do not specifically test concurrency, so a basic approach is fine.

4. **Edge Cases**
    - **Single Endpoint**: Repeatedly return the same endpoint.
    - **Empty Endpoints**: `Next()` should return `""`.

### Example Usage

```go
package main

import "fmt"

func main() {
	lb := NewRoundRobin([]string{"s1", "s2", "s3"})
	for i := 0; i < 5; i++ {
		fmt.Println(lb.Next())
		// s1, s2, s3, s1, s2 ...
	}
}
```

### Hints

- Keep a **current index**. Each call to `Next()` returns the endpoint at `current index`, then increments the index
  modulo the length of `endpoints`.
