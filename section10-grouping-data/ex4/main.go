package main

import "fmt"

// Follow these steps:
//
// - start with this slice:
//    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//
// - append to that slice this value:
//    52
//
// - print out the slice
//
// - in one statement, append to that slice these values
//
// - print out the slice
//
// - append to the slice this slice:
//    x := []int{56, 57, 58, 59, 60}
//
// - print out the slice

func main() {
	// Initial slice.
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Append to initial slice.
	x = append(x, 52)

	// Print out slice.
	fmt.Println(x)

	// Add 53, 54, and 55 to the slice.
	x = append(x, []int{53, 54, 55}...)

	// Append another slice to the initial slice.
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	// Print out the slice.
	fmt.Println(x)
}
