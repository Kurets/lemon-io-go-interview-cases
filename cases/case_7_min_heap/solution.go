package case_7_min_heap

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.bubbleDown(0)
	}
	return min, true
}

func (h *MinHeap) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] < h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) bubbleDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < len(h.data) && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < len(h.data) && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}
