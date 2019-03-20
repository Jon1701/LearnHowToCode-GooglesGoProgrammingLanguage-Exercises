package main

import "fmt"

// Write a program that:
// - assigns an int to a variable
// - prints that int in decimal, binary, and hex
// - shifts the bits of that int over 1 position to the left, and assigns that
//   to a variable
// - prints that variable in decimal, binary, and hex

func main() {
	// Declare int.
	var x int

	// Assign int to a variable.
	x = 1701

	// Prints the int in decimal, binary, and hex.
	fmt.Printf("The number %v in decimal is: %v\n", x, x)
	fmt.Printf("The number %v in binary is: %b\n", x, x)
	fmt.Printf("The number %v in hex is: %#x\n", x, x)

	// Shift x over 1 position to the left, and assign it to a variable
	y := x << 1

	// Print the shifted variable in decimal, binary, and hex.
	fmt.Printf("The number %v in decimal is: %v\n", y, y)
	fmt.Printf("The number %v in binary is: %b\n", y, y)
	fmt.Printf("The number %v in hex is: %#x\n", y, y)
}
