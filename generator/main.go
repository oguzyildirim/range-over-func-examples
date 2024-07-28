package main

import (
	"fmt"
	"iter"
)

func ToInfinity(start, step int) iter.Seq[int] {
	if step > 0 {
		// Forward iteration from start
		return func(yield func(int) bool) {
			for i := start; ; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}

	if step < 0 {
		// Backward iteration from start
		return func(yield func(int) bool) {
			for i := start; ; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}

	return func(yield func(int) bool) {
		// yield only the 'start' value when step 0
		yield(start)
	}
}

func main() {
	start := 1
	step := -5
	for i := range ToInfinity(start, step) {
		fmt.Println(i)

		if i > 100 {
			break
		}
	}
}
