package main

import "fmt"

// Create a func which returns a func
//
// Assign a returned func to a variable
//
// Call the returned func

func main() {
	// Define (outer) function which returns a(n) (inner) function.
	f := func() func() {
		return func() {
			fmt.Println("Hello World")
		}
	}

	// Save the inner function to variable.
	printMeHelloWorld := f()

	// Call inner function.
	printMeHelloWorld()
}
