package main

import "fmt"

// Create a map with a key of type string, which is a person's last_first name,
// and a value of type []string which stores their favorite things.
// Store three records in your map.
// Print out all of the values, along with their index position in the slice.

func main() {
	// Create a map.
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for i, v1 := range x {
		fmt.Println(i, v1)

		for j, v2 := range v1 {
			fmt.Println(j, v2)
		}
	}
}
