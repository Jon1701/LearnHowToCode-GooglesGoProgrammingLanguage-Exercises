package main

import "fmt"

// Create a func with the identifier foo that returns an int.
//
// Create a func with the identifier bar that returns an int and a string
//
// Call both funcs
//
// Print out their results

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1701, "Enterprise"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
