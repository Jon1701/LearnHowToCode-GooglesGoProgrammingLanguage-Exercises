package main

import (
	"fmt"
)

// 1. Use var to DECLARE three variables. The variables should have package
// level scope.
// Do not assign VALUES to the variables. Use the following IDENTIFIERS for
// the variables and make sure the variables store values of the following TYPE.
//
// a. identifier "x" type int
// b. identifier "y" type string
// c. identifier "z" type bool
//
// 2. In func main:
// a. print out the values for each identifier
// b. The compiler assigned values to the variables. What are these values called.

var x int
var y string
var z bool

func main() {
	// Printing out values for each identifier
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// The values are called zero values.
}
