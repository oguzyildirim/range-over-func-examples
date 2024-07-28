package main

import (
	"fmt"
	"time"
)

func main() {
	greeting := "Hello, World!"

	// Iterate over string using range
	for index, char := range greeting {
		fmt.Printf("Character at index %d is %c\n", index, char)
	}

	nums := []int{2, 3, 4, 5}

	// Iterate over slice using range
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	ch := make(chan string)
	go sender(ch)

	// Iterate over channel using range
	for msg := range ch {
		fmt.Println("Received:", msg)
		time.Sleep(time.Second)
	}
}

func sender(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}
