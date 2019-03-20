package main

import "fmt"

// Using the following operators, write expressions and assign their
// values to variables.
//
// 		==		<=		>=		!=		<		>
func main() {
	isEqual := "1701" == "1701A"
	isLessThanOrEqualTo := 1031 <= 1701
	isGreaterThanOrEqualTo := 1701 >= 74656
	isNotEqualTo := 1701 != 1764
	isLessThan := 1701 < 1017
	isGreaterThan := 1701 > 74205

	fmt.Printf("== %v\n", isEqual)
	fmt.Printf("<= %v\n", isLessThanOrEqualTo)
	fmt.Printf(">= %v\n", isGreaterThanOrEqualTo)
	fmt.Printf("!= %v\n", isNotEqualTo)
	fmt.Printf("< %v\n", isLessThan)
	fmt.Printf("> %v\n", isGreaterThan)

}
