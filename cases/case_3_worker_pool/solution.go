package case_3_worker_pool

import (
	"sync"
)

func startWorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		go func() {
			defer wg.Done()
			for j := range jobs {
				// Example processing: just double the input
				res := j * 2
				results <- res
			}
		}()
	}

	wg.Wait()
	close(results)
}
