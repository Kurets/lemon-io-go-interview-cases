# LRU Cache

## Task Description

Implement an **LRU (Least Recently Used) cache** that is safe for concurrent use. The cache should support the following
operations:

1. **Get(key)**:
    - Retrieve the value associated with the key.
    - Return `false` if the key is not present in the cache.

2. **Put(key, value)**:
    - Insert a key-value pair into the cache.
    - If the cache reaches its capacity, evict the least recently used item.

### Requirements

- The cache should be **thread-safe** and allow concurrent access.
- Handle scenarios where the cache capacity is exceeded by evicting the least recently used item.

### Function Signatures

```go
package main

type LRUCache struct {
   // Your code here
}

func NewLRUCache(cap int) *LRUCache {
   // Your code here
   return nil
}

func (c *LRUCache) Get(key string) (string, bool) {
   // Your code here
   return "", false
}

func (c *LRUCache) Put(key, value string) {
   // Your code here
}

```

### Example Usage

```go
package main

import "fmt"

func main() {
cache := NewLRUCache(2)

    cache.Put("a", "1")
    cache.Put("b", "2")

    fmt.Println(cache.Get("a")) // Output: "1", true

    cache.Put("c", "3")         // Evicts key "b"

    fmt.Println(cache.Get("b")) // Output: "", false
    fmt.Println(cache.Get("c")) // Output: "3", true

    cache.Put("d", "4")         // Evicts key "a"

    fmt.Println(cache.Get("a")) // Output: "", false
    fmt.Println(cache.Get("d")) // Output: "4", true

}
```

### Notes

- Use a **doubly linked list** to maintain the order of usage.
- Use a **hash map** for fast access to nodes in the linked list.
- Ensure thread safety by protecting shared data structures with a synchronization mechanism like `sync.Mutex`.
