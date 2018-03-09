package main

import (
	"fmt"
)

func main() {

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{5, 2, 3, 4}
	foo(aSlice...)
	foo()

	fmt.Println(foo(aSlice...))
}

func foo(args ...int) int {

	for _, n := range args {

		return n
	}

	return 1
}
