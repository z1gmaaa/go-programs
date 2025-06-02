package main

import "fmt"

func main() {
	var n int
	var val int
	fmt.Println("Enter the array size:")
	fmt.Scan(&n)
	var arr1 = []int{}
	for i := 0; i < n; i++ {
		fmt.Println("Enter the array element:")
		fmt.Scan(&val)
		arr1 = append(arr1, val)
	}
	sort(arr1, n)
	fmt.Println("sorted array is :")
	for i := 0; i < n; i++ {

		fmt.Println(arr1[i])

	}
	var x int
	fmt.Println("Enter the element to search")
	fmt.Scan(&x)

	search(arr1, n, x)
}

func sort(arr []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				var temp int
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp

			}
		}
	}
	return arr
}

func search(arr []int, n int, x int) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := high - low/2

		if arr[mid] == x {
			fmt.Println("element found at ", mid+1)
			return mid
		}
		if arr[mid] < x {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}

	fmt.Println("element not found")
	return -1

}
