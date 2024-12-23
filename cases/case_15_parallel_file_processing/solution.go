package case_15_parallel_file_processing

import "sync"

var testProcessFunc = func(file string) {
	// default no-op
}

// processFiles processes a list of files using a specified number of workers in parallel.
func processFiles(files []string, workers int) {
	if len(files) == 0 {
		// No files to process, nothing to do.
		return
	}
	if workers <= 0 {
		// If invalid workers count, default to 1
		workers = 1
	}

	fileChan := make(chan string)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range fileChan {
				// Call the test function so the tests know the file was processed
				testProcessFunc(f)
			}
		}()
	}

	// Send files to workers
	for _, f := range files {
		fileChan <- f
	}
	close(fileChan)

	// Wait for all workers to finish
	wg.Wait()
}
