package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 1000; i++ {

		if i%3 == 0 || i%5 == 0 {

			hold := sum
			sum = sum + i
			fmt.Println(hold, " + ", i, " = ", sum)
			//fmt.Println("FizzBuzz")
		}

		/*else if i%3 == 0 {

			//fmt.Println("Fizz")
		} else if i%5 == 0 {

			//fmt.Println("Buzz")
		} else {
			//println(i)
		}*/

	}

	fmt.Println("Sum of multiples of 3 and 5 below 100", sum)

}
