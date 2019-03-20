package main

import "fmt"

// Create a for loop using the for condition syntax.
// Have it printout the years you have been alive.
func main() {
	yob := 1987

	for yob <= 2019 {
		fmt.Println(yob)
		yob++
	}
}
