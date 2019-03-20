package main

import "fmt"

// Create a program which uses a switch statement with no switch
// expression specified.

func main() {
	switch {
	case true:
		fmt.Println("case 1")

	default:
		fmt.Println("default case")
	}
}
