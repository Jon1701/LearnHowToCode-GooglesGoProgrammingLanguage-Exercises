package main

import "fmt"

// Create a value and assign it to a variable.
//
// Print the address of that value.

func main() {
	x := 3.14

	fmt.Println(&x)
}
