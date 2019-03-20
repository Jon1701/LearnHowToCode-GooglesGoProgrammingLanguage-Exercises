package main

import "fmt"

// Create types and untyped constants.
// Print the values of the constants.

func main() {
	const typed int = 1701
	const untyped = "1701A"

	fmt.Printf("Typed constant: %v\n", typed)
	fmt.Printf("Untyped constant: %v\n", untyped)
}
