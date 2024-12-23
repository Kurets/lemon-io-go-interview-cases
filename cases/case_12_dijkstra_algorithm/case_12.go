package case_12_dijkstra_algorithm

import (
	"container/heap"
	"math"
)

type PriorityQueueItem struct {
	node     int
	distance int
}

type PriorityQueue []PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(PriorityQueueItem))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph map[int][]struct{ to, weight int }, start int) map[int]int {
	dist := make(map[int]int)

	for node := range graph {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, PriorityQueueItem{node: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(PriorityQueueItem)
		currentNode, currentDist := current.node, current.distance

		if currentDist > dist[currentNode] {
			continue
		}

		// Relax edges
		for _, edge := range graph[currentNode] {
			nextNode, edgeWeight := edge.to, edge.weight
			newDist := currentDist + edgeWeight
			if newDist < dist[nextNode] {
				dist[nextNode] = newDist
				heap.Push(pq, PriorityQueueItem{node: nextNode, distance: newDist})
			}
		}
	}

	return dist
}

/*
func main() {
	// Example usage:
	graph := map[int][]struct{to, weight int}{
		0: {{to: 1, weight: 4}, {to: 2, weight: 1}},
		1: {{to: 3, weight: 1}},
		2: {{to: 1, weight: 2}, {to: 3, weight: 5}},
		3: {},
	}
	start := 0
	shortestDistances := dijkstra(graph, start)
	// shortestDistances should contain the shortest distance from node 0 to all other nodes.
	// Example: From node 0: dist[0] = 0, dist[1] = 3, dist[2] = 1, dist[3] = 4
	for node, d := range shortestDistances {
		println("Node:", node, "Distance:", d)
	}
}
*/
