package main

import "fmt"

// - Using a composite literal:
//   - create an array which holds 5 values of type int
//   - assigns values to each index position
// - range over the array and print the values out
// - using format printing
//   - print out the type of the array

func main() {
	// Create array which holds 5 values of type int.
	arr := [5]int{1701, 1031, 1227, 1764, 1867}

	// Range over the array and print the values out.
	for _, v := range arr {
		fmt.Println(v)
	}

	// Print the type of the array.
	fmt.Printf("%T\n", arr)
}
