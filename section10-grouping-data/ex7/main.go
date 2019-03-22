package main

import "fmt"

// Create a slice of a slice of string.
// Store the following data in the multi-dimensional slice.
// - "James", "Bond", "Shaken, not stirred"
// - "Miss", "Moneypenny", "Hello James"
// Range over the records, th en range over the data in each record

func main() {
	// Main slice
	slice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hello James"}}

	// Print the data.
	for _, v1 := range slice {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
