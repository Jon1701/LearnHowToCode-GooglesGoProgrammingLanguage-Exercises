package main

import "fmt"

// Build and use an anonymous function.

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	func(x []int) {
		sum := 0

		for _, v := range x {
			sum += v
		}

		fmt.Println(sum)
	}(s)
}
