package main

import "fmt"

// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
//   - first name
//   - last name
//   - favorite ice cream flavors
//
// Create two VALUES of TYPE person.
//
// Print out the values, ranging over the elements in the slice which stores
// the favorite flavors.

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:               "Lorem",
		lastName:                "Ipsum",
		favoriteIceCreamFlavors: []string{"Chocolate", "Vanilla"},
	}

	p2 := person{
		firstName:               "Hello",
		lastName:                "World",
		favoriteIceCreamFlavors: []string{"Rocky Road", "Cookies & Creme"},
	}

	// Create a slice containing the 2 person structs.
	persons := []person{p1, p2}

	// Iterate over the persons, and display values.
	for i, p := range persons {
		fmt.Println(i, p)
		fmt.Printf("\tFirst Name: %v\n", p.firstName)
		fmt.Printf("\tLast Name: %v\n", p.lastName)

		fmt.Printf("\t\tFavorite Ice Cream Flavors:\n")
		for _, flavor := range p.favoriteIceCreamFlavors {
			fmt.Printf("\t\t%v\n", flavor)
		}

		fmt.Println()
	}
}
