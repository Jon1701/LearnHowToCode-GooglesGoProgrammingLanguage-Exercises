package main

import (
	"fmt"
)

// 1. At the package level scope, assign the following values to the three variables
// a. for x assign 42
// b. for y assigned "James Bond"
// c. for z assign true
//
// 2. in func main
// a. use fmt.Sprintf to print all of the VALUES to one string. Assign the returned
// value of TYPE string using the short declaration operator to a VARIABLE with the
// IDENTIFIER "s"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
