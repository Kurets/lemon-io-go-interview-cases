package case_11_tarjan_algorithm

func tarjanSCC(graph map[int][]int) [][]int {
	index := 0
	var stack []int
	onStack := make(map[int]bool)
	indices := make(map[int]int)
	lowLink := make(map[int]int)
	for v := range graph {
		indices[v] = -1
	}

	var sccs [][]int

	var strongConnect func(v int)
	strongConnect = func(v int) {
		indices[v] = index
		lowLink[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		for _, w := range graph[v] {
			if indices[w] == -1 {
				strongConnect(w)
				if lowLink[w] < lowLink[v] {
					lowLink[v] = lowLink[w]
				}
			} else if onStack[w] {
				if indices[w] < lowLink[v] {
					lowLink[v] = indices[w]
				}
			}
		}

		if lowLink[v] == indices[v] {
			var scc []int
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			sccs = append(sccs, scc)
		}
	}

	for v := range graph {
		if indices[v] == -1 {
			strongConnect(v)
		}
	}
	return sccs
}
