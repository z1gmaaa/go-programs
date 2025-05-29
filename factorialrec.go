package main

import "fmt"

func fact(a int) int {
	if a == 1 {
		return a
	} else {
		return (a * fact(a-1))
	}
}

func main() {
	var x int
	fmt.Print("Enter a number: ")
	fmt.Scan(&x)
	fmt.Println("factorial:", fact(x))
}
