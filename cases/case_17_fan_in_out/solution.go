package case_17_fan_in_out

import (
	"sync"
)

func fanOut(in <-chan int, workers int) []<-chan int {
	workerIn := make([]chan int, workers)
	for i := 0; i < workers; i++ {
		workerIn[i] = make(chan int)
	}

	outChans := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		out := make(chan int)
		outChans[i] = out

		go func(inCh <-chan int, outCh chan<- int) {
			defer close(outCh)
			for item := range inCh {
				outCh <- item
			}
		}(workerIn[i], out)
	}

	go func() {
		defer func() {
			for _, ch := range workerIn {
				close(ch)
			}
		}()

		i := 0
		for val := range in {
			workerIn[i] <- val
			i = (i + 1) % workers
		}
	}()

	return outChans
}

func fanIn(channels []<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for _, ch := range channels {
			go func(c <-chan int) {
				defer wg.Done()
				for val := range c {
					out <- val
				}
			}(ch)
		}

		wg.Wait()
		close(out)
	}()
	return out
}
