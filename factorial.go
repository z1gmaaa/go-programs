package main

import "fmt"

func main() {
	var x int
	var fct = 1
	fmt.Println("Enter the number")
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fct *= i
	}
	fmt.Println("Factorial of", x, "is", fct)
}
