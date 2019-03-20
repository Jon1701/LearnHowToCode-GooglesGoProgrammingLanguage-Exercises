package main

import "fmt"

// Create a program that uses a switch statment with the switch
// expression specified as a variable of TYPE string with the
// IDENTIFIER "favSport"

func main() {
	favSport := "tetris"

	switch favSport {
	case "tetris":
		fmt.Println("Tetris is my favourite sport")

	default:
		fmt.Println("Unknown sport")
	}

}
