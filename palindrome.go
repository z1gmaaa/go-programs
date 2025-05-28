package main

import "fmt"

func main() {
	var n int
	var d int
	var rev = 0
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	var temp = n
	for n != 0 {
		d = n % 10
		rev = rev*10 + d
		n = n / 10

	}
	fmt.Println("reverse is ", rev)
	if temp == rev {
		fmt.Println("number is palindrome")
	} else {
		fmt.Println("number is not palindrome")
	}
}
