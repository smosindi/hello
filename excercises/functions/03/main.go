package main

import (
	"fmt"
)

func main() {

	numbers := []float64{1, 9, 3, 11, 5, 6, 7}

	fmt.Println(maxnumber(numbers...))

	fmt.Println("Answer ", ((true && false) || (false && true) || !(false && true)))

}

func maxnumber(args ...float64) float64 {

	var holder float64

	for _, n := range args {

		if n > holder {

			holder = n
		}
	}

	return holder

}
