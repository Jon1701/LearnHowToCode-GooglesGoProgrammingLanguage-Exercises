package main

import "fmt"

// Create and use an anonymous struct.

func main() {
	p1 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Lorem",
		lastName:  "Ipsum",
	}

	fmt.Println(p1)
}
