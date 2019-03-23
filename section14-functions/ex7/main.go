package main

import "fmt"

// Assign a func to a variable, then call that func

func main() {
	sum := func(x []float64) float64 {
		sum := 0.0

		for _, v := range x {
			sum += v
		}

		return sum
	}

	result := sum([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(result)
}
