package main

import "fmt"

// Create a slice to store the names of all the states in the United States of
// America.
//
// - What is the length of your slice?
// - What is the capacity?
// - Print out all of the values, along with their index position in the slice,
//   without using the range clause

func main() {
	// Initial slice.
	states := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
	}

	// Length of the slice.
	fmt.Printf("Length of the slice: %v\n", len(states))

	// Capacity of the slice.
	fmt.Printf("Capacity of the slice: %v\n", cap(states))

	// Print out all the values, along with their index, without using range.
	for i := 0; i < len(states); i++ {
		fmt.Printf("%v\t%v\n", i, states[i])
	}
}
