package main

import "fmt"

// Take the code from the previous exercise, then store the values of type
// person in a map with the key of last name. Access each value in the map.
// Print out the values, ranging over the slice.

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:               "Lorem",
		lastName:                "Ipsum",
		favoriteIceCreamFlavors: []string{"Chocolate", "Vanilla", "Strawberry"},
	}

	p2 := person{
		firstName:               "Hello",
		lastName:                "World",
		favoriteIceCreamFlavors: []string{"Rocky Road", "Cookies & Creme", "Mint Chocolate"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for lastName, p := range m {
		fmt.Printf("Last Name: %v\n", lastName)
		fmt.Printf("\tFirst Name: %v\n", p.firstName)
		fmt.Printf("\tLast Name: %v\n", p.lastName)

		fmt.Printf("\tFavorite Ice Cream Flavors:\n")
		for _, flavor := range p.favoriteIceCreamFlavors {
			fmt.Printf("\t\t%v\n", flavor)
		}

		fmt.Println()
	}
}
