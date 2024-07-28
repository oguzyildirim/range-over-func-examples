package main

import (
	"fmt"
	"time"
)

func main() {
	for range every(2 * time.Second) {
		fmt.Println("ping")
	}
}

func every(d time.Duration) func(func() bool) {
	return func(yield func() bool) {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if !yield() {
					return
				}
			}
		}
	}
}
