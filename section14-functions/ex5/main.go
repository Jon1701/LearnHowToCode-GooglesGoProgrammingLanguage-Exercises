package main

import "fmt"

// Create a type square
//
// create a type circle
//
// Attach a method to each that calculates area and returns it
//
// Create a type shape which defines an interface as anything which has the
// area method.
//
// Create a value of type square
//
// Create a value of type cirlce
//
// Use func info to print the area of the square
//
// Use func info to print the area of the circle

// Square struct.
type square struct {
	length float32
	width  float32
}

// Circle struct.
type circle struct {
	radius float32
}

// Area method for square.
func (s square) area() float32 {
	return s.length * s.width
}

// Area method for circle
func (c circle) area() float32 {
	return 3.141592654 * c.radius * c.radius
}

// Define a shape interface as anything which has an area method.
type shape interface {
	area() float32
}

// Function which prints the area of the shape.
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// Define a square.
	s1 := square{
		17,
		64,
	}

	// Define a circle.
	c1 := circle{
		5,
	}

	// Print area of square.
	info(s1)

	// Print area of circle.
	info(c1)
}
