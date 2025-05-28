package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println("Enter a number")
	fmt.Scanln(&x)
	if x%2 == 0 {
		fmt.Println("it is an even number")
	} else {
		fmt.Println("it is an odd number")
	}
}
