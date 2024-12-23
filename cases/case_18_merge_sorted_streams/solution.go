package case_18_merge_sorted_streams

import "container/heap"

type streamItem struct {
	val int
	ch  <-chan int
}

type minHeap []streamItem

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(streamItem)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeSortedChannels(chs []<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		pq := &minHeap{}
		heap.Init(pq)
		for _, c := range chs {
			if val, ok := <-c; ok {
				heap.Push(pq, streamItem{val, c})
			}
		}

		for pq.Len() > 0 {
			item := heap.Pop(pq).(streamItem)
			out <- item.val
			if val, ok := <-item.ch; ok {
				heap.Push(pq, streamItem{val, item.ch})
			}
		}
		close(out)
	}()
	return out
}
