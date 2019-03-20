package main

import "fmt"

/*
1. At the package level scope, using the var keyword, create a VARIABLE with the
	 IDENTIFIER "y". The variable should be of the UNDERLYING TYPE of your custom
	 type "x".

2. In func main:
	a. this should already be done.
		i. print out the value of the variable "x"
		ii. print out the type of the variable "x"
		iii. assign your own value to the variable "x" using the "=" operator
		iv. print out the value of the variable x
	b. now do this
		i. now use conversion to convert the type of the value stored in "x" to the
		   underlying type
			1. then use the short declaration operator to assign that value to "y"
*/

// Declare custom type.
type custom int

// Declare variable x of type custom.
var x custom

// Declare variable y of type custom.
var y int

func main() {
	// Print out value and type of x.
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// Assign value to x
	x = 42

	// Print value of x
	fmt.Println(x)

	// Convert type of x to the underlying type, and assign to y
	y = int(x)

	// Print type of y
	fmt.Printf("%T\n", y)
}
