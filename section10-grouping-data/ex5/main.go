package main

import "fmt"

// Follow these steps:
//
// - start with this slice:
//     x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//
// - use append and slicing to get the values here which you should then print:
//     [42, 43, 44, 48, 49, 50, 51]

func main() {
	// Initial slice.
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Modify the slice.
	x = append(x[:3], x[6:]...)

	// Print the slice.
	fmt.Println(x)
}
