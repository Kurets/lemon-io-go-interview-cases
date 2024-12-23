package case_13_topological_sort

func topologicalSort(graph map[int][]int) []int {
	inDegree := make(map[int]int)
	for u := range graph {
		if _, ok := inDegree[u]; !ok {
			inDegree[u] = 0
		}
		for _, v := range graph[u] {
			inDegree[v]++
		}
	}

	queue := []int{}
	for v, d := range inDegree {
		if d == 0 {
			queue = append(queue, v)
		}
	}

	var order []int
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, w := range graph[v] {
			inDegree[w]--
			if inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return order
}
