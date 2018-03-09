package main

import

(
	"fmt"
)

var smallnumber = 0
var bignumber = 0

func main() {

fmt.Println("Enter small numbber : ")
fmt.Scan(&smallnumber)
fmt.Println("Enter big number")
fmt.Scan(&bignumber)

fmt.Println("Remainder of ", bignumber, " divided by ", smallnumber, " is ", bignumber % smallnumber)

}