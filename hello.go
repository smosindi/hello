package main

import (
	"fmt"
)

var name = ""

func main() {

	//fmt.Printf(stringutil.Reverse("!ssoG ,olleH"))
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Hello my name is : ", name)

}
