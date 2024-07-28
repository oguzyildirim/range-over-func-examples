package main

import (
	"fmt"
	"iter"
)

func iterate(yield func(int) bool) {
	items := []int{1, 5, 10}
	for _, i := range items {
		if !yield(i) {
			return
		}
	}
}

func iterate2() iter.Seq[int] {
	return func(yield func(int) bool) {
		items := []int{1, 2, 3}
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	for i := range iterate {
		fmt.Println(i)

		if i > 10 {
			break
		}
	}

	for i := range iterate2() {
		fmt.Println(i)

		if i > 10 {
			break
		}
	}
}
