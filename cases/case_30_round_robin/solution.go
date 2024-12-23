package case_30_round_robin

import (
	"fmt"
	"sync"
)

// The approach:
// 1. Keep a slice of endpoints and an index.
// 2. Each call to Next() returns endpoints[index], then increment index mod length.
// 3. Protect with a mutex for concurrency safety.

type RoundRobin struct {
	mu        sync.Mutex
	endpoints []string
	idx       int
}

func NewRoundRobin(endpoints []string) *RoundRobin {
	return &RoundRobin{endpoints: endpoints}
}

func (r *RoundRobin) Next() string {
	if len(r.endpoints) == 0 {
		return ""
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	e := r.endpoints[r.idx]
	r.idx = (r.idx + 1) % len(r.endpoints)
	return e
}

func main() {
	rr := NewRoundRobin([]string{"server1", "server2", "server3"})
	for i := 0; i < 6; i++ {
		fmt.Println(rr.Next())
	}
}
