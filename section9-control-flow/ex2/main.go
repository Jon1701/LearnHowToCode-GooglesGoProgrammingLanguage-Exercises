package main

import "fmt"

// Print every rune code point of the uppercase alphabet three times.

func main() {
	// Iterate over the decimal numbers of the uppercase letters
	// in the ASCII table.
	for i := 65; i <= 90; i++ {
		// Print the index, offset by -64.
		fmt.Println(i - 64)
		// Print the decimal number in unicode format three times.
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U %q\n", i, i)
		}

	}
}
