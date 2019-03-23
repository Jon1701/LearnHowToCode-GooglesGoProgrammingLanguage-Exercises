package main

import "fmt"

// Use the "defer" keyword to show that a deferred func runs after func containing
// it exits

func normalFunction() {
	fmt.Println("This function is normal")
}

func deferredFunction() {
	fmt.Println("This function is deferred")
}

func main() {
	// This function should execute last.
	defer deferredFunction()

	// This function should execute first.
	normalFunction()
}
