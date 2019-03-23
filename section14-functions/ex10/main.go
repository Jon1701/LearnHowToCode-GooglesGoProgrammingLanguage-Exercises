package main

import "fmt"

// Create a func which encloses the scope of a variable.

func decrementor() func() float64 {
	value := 0.0

	return func() float64 {
		value--
		return value
	}
}

func main() {
	d := decrementor()

	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
}
