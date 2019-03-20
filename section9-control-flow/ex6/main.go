package main

import "fmt"

// Create a program which shows the "if statement" in action.

func main() {
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
