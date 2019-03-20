package main

import "fmt"

func main() {
	birthYear := 1987
	currentYear := 2019

	for {
		if birthYear <= currentYear {
			fmt.Println(birthYear)
		} else {
			break
		}
		birthYear++
	}
}
