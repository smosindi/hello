package main

import "fmt"

func main() {

	funcexpre := func(n int) (res float64, check bool) {

		if n%2 == 0 {
			check = true
		}

		res = (float64)(n / 2)

		return res, check

	}

	fmt.Println(funcexpre(4))
}
