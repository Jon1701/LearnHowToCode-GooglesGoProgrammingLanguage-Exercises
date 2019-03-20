package main

import "fmt"

// Using iota, create 4 constants for the last 4 years. Print the constant values
func main() {
	const (
		year1 = 2019 - iota
		year2
		year3
		year4
	)

	fmt.Println(year1, year2, year3, year4)
}
