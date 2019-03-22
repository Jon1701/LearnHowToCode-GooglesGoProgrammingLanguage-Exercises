package main

import "fmt"

// Using the code from the previous example, delete a record to your map.
// Now print the map out using the range loop.

func main() {
	// Create a map.
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
		"person_unknown":  []string{"Lorem", "Ipsum", "Hello World"},
	}

	// Delete a record.
	if _, ok := x["no_dr"]; ok {
		delete(x, "no_dr")
	}

	// Print out the map.
	for lastName, array := range x {
		fmt.Printf("%v\n", lastName)

		for _, v := range array {
			fmt.Printf("\t%v\n", v)
		}
		fmt.Printf("\n")
	}
}
