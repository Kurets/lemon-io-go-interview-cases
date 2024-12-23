package case_21_streaming_file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcessFileBasic(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_process_file")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "test.txt")
	lines := []string{
		"line one",
		"line two",
		"line three",
	}
	if err := os.WriteFile(testFile, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	err = processFile(testFile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProcessFileEmpty(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test_process_file_empty")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	emptyFile := filepath.Join(tempDir, "empty.txt")
	if err := ioutil.WriteFile(emptyFile, []byte(""), 0644); err != nil {
		t.Fatalf("failed to write empty file: %v", err)
	}

	err = processFile(emptyFile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProcessFileNonExistent(t *testing.T) {
	err := processFile("non_existent_file.txt")
	if err == nil {
		t.Error("expected an error for non-existent file, got nil")
	}
}
