package case_26_fifo_queue

import (
	"sync"
	"testing"
)

func TestSafeQueueBasic(t *testing.T) {
	var q SafeQueue

	val, ok := q.Dequeue()
	if ok {
		t.Errorf("expected empty queue to return ok=false, got ok=true with val=%v", val)
	}

	q.Enqueue("first")
	q.Enqueue("second")

	val, ok = q.Dequeue()
	if !ok || val != "first" {
		t.Errorf("expected 'first', got %v (ok=%v)", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != "second" {
		t.Errorf("expected 'second', got %v (ok=%v)", val, ok)
	}

	val, ok = q.Dequeue()
	if ok {
		t.Errorf("expected empty queue after all dequeued, got val=%v (ok=true)", val)
	}
}

func TestSafeQueueMixed(t *testing.T) {
	var q SafeQueue

	q.Enqueue(1)
	q.Enqueue(2)
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v (ok=%v)", val, ok)
	}

	q.Enqueue(3)
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("expected 2, got %v (ok=%v)", val, ok)
	}

	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("expected 3, got %v (ok=%v)", val, ok)
	}
}

func TestSafeQueueConcurrent(t *testing.T) {
	var q SafeQueue
	const numWorkers = 10
	const numOpsPerWorker = 100

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		go func(offset int) {
			defer wg.Done()
			for i := 1; i <= numOpsPerWorker; i++ {
				q.Enqueue(offset*1000 + i)
			}
		}(w)
	}

	wg.Wait()

	expectedCount := numWorkers * numOpsPerWorker

	var mu sync.Mutex
	results := make([]int, 0, expectedCount)

	wg.Add(numWorkers)
	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			for {
				val, ok := q.Dequeue()
				if !ok {
					return
				}
				mu.Lock()
				results = append(results, val.(int))
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	if len(results) != expectedCount {
		t.Errorf("expected %d total dequeued items, got %d", expectedCount, len(results))
	}

	val, ok := q.Dequeue()
	if ok {
		t.Errorf("expected empty queue, got val=%v (ok=true)", val)
	}
}

func TestSafeQueueDataIntegrity(t *testing.T) {
	var q SafeQueue
	inputVals := []int{10, 20, 30, 10, 40, 10, 20}
	valueCounts := make(map[int]int)

	for _, v := range inputVals {
		q.Enqueue(v)
		valueCounts[v]++
	}

	for i := 0; i < len(inputVals); i++ {
		val, ok := q.Dequeue()
		if !ok {
			t.Errorf("expected more items, but queue is empty at dequeue #%d", i)
			return
		}
		intVal := val.(int)
		valueCounts[intVal]--
		if valueCounts[intVal] < 0 {
			t.Errorf("dequeued unexpected extra value %d", intVal)
		}
	}

	for v, c := range valueCounts {
		if c != 0 {
			t.Errorf("value %d was not fully dequeued, final count is %d", v, c)
		}
	}

	val, ok := q.Dequeue()
	if ok {
		t.Errorf("expected empty queue, got val=%v (ok=true)", val)
	}
}
