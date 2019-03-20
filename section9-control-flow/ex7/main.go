package main

import "fmt"

// Create a program that uses "else if" and else.

func main() {
	for i := -100; i <= 100; i++ {
		if i < 0 {
			fmt.Printf("%v is negative\n", i)
		} else if i == 0 {
			fmt.Printf("%v is 0\n", i)
		} else {
			fmt.Printf("%v is positive\n", i)
		}
	}
}
