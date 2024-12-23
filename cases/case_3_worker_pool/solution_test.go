package case_3_worker_pool

import (
	"testing"
	"time"
)

func TestBasicProcessing(t *testing.T) {
	numJobs := 5
	numWorkers := 2

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go startWorkerPool(numWorkers, jobs, results)

	expectedResults := map[int]bool{2: true, 4: true, 6: true, 8: true, 10: true}
	for r := range results {
		if _, ok := expectedResults[r]; !ok {
			t.Errorf("Unexpected result: %d", r)
		}
		delete(expectedResults, r)
	}

	if len(expectedResults) > 0 {
		t.Errorf("Not all results were processed: remaining %v", expectedResults)
	}
}

func TestEmptyJobsChannel(t *testing.T) {
	numWorkers := 3
	jobs := make(chan int)
	results := make(chan int)

	close(jobs)

	go startWorkerPool(numWorkers, jobs, results)

	select {
	case _, ok := <-results:
		if ok {
			t.Errorf("Expected no results, but got some")
		}
	case <-time.After(1 * time.Second):
	}
}

func TestManyWorkersFewJobs(t *testing.T) {
	numWorkers := 10
	numJobs := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go startWorkerPool(numWorkers, jobs, results)

	expectedResults := map[int]bool{2: true, 4: true, 6: true}
	for r := range results {
		if _, ok := expectedResults[r]; !ok {
			t.Errorf("Unexpected result: %d", r)
		}
		delete(expectedResults, r)
	}

	if len(expectedResults) > 0 {
		t.Errorf("Not all results were processed: remaining %v", expectedResults)
	}
}

func TestFewWorkersManyJobs(t *testing.T) {
	numWorkers := 2
	numJobs := 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go startWorkerPool(numWorkers, jobs, results)

	expectedResults := make(map[int]bool)
	for i := 1; i <= numJobs; i++ {
		expectedResults[i*2] = true
	}

	for r := range results {
		if _, ok := expectedResults[r]; !ok {
			t.Errorf("Unexpected result: %d", r)
		}
		delete(expectedResults, r)
	}

	if len(expectedResults) > 0 {
		t.Errorf("Not all results were processed: remaining %v", expectedResults)
	}
}
