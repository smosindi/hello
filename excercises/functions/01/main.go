package main

import "fmt"

func main() {

	fmt.Println(tworeturns(53))
}

func tworeturns(n int) (res float64, check bool) {

	if n%2 == 0 {
		check = true
	}

	res = float64(n) / 2

	return res, check

}
