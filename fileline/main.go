package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

func LineReader(filename string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if !yield(line) {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning file: %s\n", err)
			return
		}
	}
}

func main() {
	filename := "fileline/file.txt"

	for line := range LineReader(filename) {
		fmt.Println(line)
	}

}
