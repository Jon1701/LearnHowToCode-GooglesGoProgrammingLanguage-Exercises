package main

import "fmt"

// A "callback" is when we pass a func into a func as an argument.
// For this exercise, pass a func into a func as an argument.

func info(s []float64, f func(x []float64) float64) {
	fmt.Printf("Applying function to slice yields this result: %v\n", f(s))
}

func main() {
	// Slice of floats.
	s := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Define a summation function.
	sum := func(x []float64) float64 {
		sum := 0.0

		for _, v := range x {
			sum += v
		}

		return sum
	}

	info(s, sum)
}
