package case_26_fifo_queue

import (
	"sync"
)

type SafeQueue struct {
	mu sync.Mutex
	q  []interface{}
}

func (q *SafeQueue) Enqueue(val interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = append(q.q, val)
}

func (q *SafeQueue) Dequeue() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) == 0 {
		return nil, false
	}
	val := q.q[0]
	q.q = q.q[1:]
	return val, true
}
