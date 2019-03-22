package main

import "fmt"

// - Using a composite literal:
//   - create a slice of type int
//   - assign 10 values
// - range over the slice and print the values out
// - using format printing
//   - print out the type of the slice

func main() {
	// Create a slice of type int, and assign 10 values
	slice := []int{1701, 1764, 1031, 1227, 1864, 74205, 74656, 73811, 1017, 1305}

	// Range over the slice and print the values out
	for _, v := range slice {
		fmt.Println(v)
	}

	// Print out the type of the slice.
	fmt.Printf("%T\n", slice)
}
