package case_21_streaming_file

import (
	"bufio"
	"fmt"
	"os"
)

func processFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Process line
		fmt.Println(line)
	}

	return scanner.Err()
}
