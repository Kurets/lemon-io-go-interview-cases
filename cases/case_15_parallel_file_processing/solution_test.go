package case_15_parallel_file_processing

import (
	"sync"
	"testing"
	"time"
)

func TestAllFilesProcessed(t *testing.T) {
	files := []string{"file1", "file2", "file3", "file4"}
	processed := make(map[string]bool)
	var mu sync.Mutex

	testProcessFunc = func(file string) {
		mu.Lock()
		processed[file] = true
		mu.Unlock()
	}

	processFiles(files, 2)

	for _, f := range files {
		if !processed[f] {
			t.Errorf("file %s was not processed", f)
		}
	}
}

func TestEmptyInput(t *testing.T) {
	testProcessFunc = func(file string) {
		t.Errorf("no files should be processed, but got: %s", file)
	}
	processFiles([]string{}, 5)
}

func TestSingleWorker(t *testing.T) {
	files := []string{"a", "b", "c"}
	var processedCount int
	var mu sync.Mutex

	testProcessFunc = func(file string) {
		mu.Lock()
		processedCount++
		mu.Unlock()
	}

	processFiles(files, 1)

	if processedCount != len(files) {
		t.Errorf("expected %d files processed, got %d", len(files), processedCount)
	}
}

func TestMultipleWorkersPerformance(t *testing.T) {
	files := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		files[i] = "file" + string(rune(i))
	}

	var processedCount int
	var mu sync.Mutex
	testProcessFunc = func(file string) {
		time.Sleep(1 * time.Millisecond)
		mu.Lock()
		processedCount++
		mu.Unlock()
	}

	start := time.Now()
	processFiles(files, 10)
	elapsed := time.Since(start)

	if processedCount != 1000 {
		t.Errorf("not all files processed, got %d", processedCount)
	}

	if elapsed > 2*time.Second {
		t.Errorf("processing took too long, might not be running in parallel. took: %v", elapsed)
	}
}

func TestWorkersSpeedup(t *testing.T) {
	files := make([]string, 100)
	for i := 0; i < 100; i++ {
		files[i] = "file" + string(rune(i))
	}

	testProcessFunc = func(file string) {
		time.Sleep(10 * time.Millisecond)
	}

	startSingle := time.Now()
	processFiles(files, 1)
	elapsedSingle := time.Since(startSingle)

	startMulti := time.Now()
	processFiles(files, 10)
	elapsedMulti := time.Since(startMulti)

	if elapsedMulti >= elapsedSingle {
		t.Errorf("expected multiple workers to be faster than single worker. single: %v, multi: %v", elapsedSingle, elapsedMulti)
	}
}
