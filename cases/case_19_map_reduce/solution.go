package case_19_map_reduce

import (
	"sync"
)

func mapReduce(inputs []int, mapper func(int) int, reducer func(int, int) int) int {
	if len(inputs) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	mapped := make([]int, len(inputs))

	for i, in := range inputs {
		wg.Add(1)
		go func(i, val int) {
			defer wg.Done()
			mapped[i] = mapper(val)
		}(i, in)
	}
	wg.Wait()

	result := mapped[0]
	for i := 1; i < len(mapped); i++ {
		result = reducer(result, mapped[i])
	}
	return result
}
