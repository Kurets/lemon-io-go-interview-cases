# Line-by-Line File Processing

## Task Description

You need to implement a function that reads a file **line by line** using a buffered approach and processes each line **as it is read**, without loading the entire file into memory.

```go
func processFile(filename string) error {
	// Your code here:
	// 1. Open the file
	// 2. Create a buffered reader (e.g., bufio.NewReader or bufio.NewScanner)
	// 3. Iterate through each line
	// 4. Process the line (e.g., print/log/store it)
	// 5. Close the file
	return nil
}
```

### Requirements

1. **Open the File**
    - Return an error if the file cannot be opened.

2. **Read Line by Line**
    - Use a buffered reader (e.g., `bufio.Scanner` or `bufio.Reader`) to read each line.
    - Do **not** store the entire file’s contents in memory at once.

3. **Process Each Line**
    - After reading a line, do some form of “processing” (printing, logging, or saving to a data structure).
    - The provided tests assume you can handle any lines. Real implementation details can vary.

4. **Close the File**
    - Ensure the file is closed when you’re done.

5. **Error Handling**
    - If reading fails, return the error.
    - If the file doesn’t exist, return an error (e.g., `os.ErrNotExist`).

### Example Usage

```go
package main

import "fmt"

func main() {
	err := processFile("example.txt")
	if err != nil {
		fmt.Println("Failed to process file:", err)
		return
	}
	fmt.Println("File processed successfully")
}
```

### Hints

- Use `os.Open(filename)` to get an `*os.File`.
- Pass that to a `bufio.Scanner` or `bufio.NewReader`.
- Loop over each line with `scanner.Scan()` or `reader.ReadString('\n')`.
- Handle potential partial lines (EOF without a newline).
- Be sure to `defer file.Close()` after successfully opening.

