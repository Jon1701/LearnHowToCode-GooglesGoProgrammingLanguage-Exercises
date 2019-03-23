package main

import "fmt"

// Create a person struct
//
// create a func called "changeMe" with a "*person" as a
// parameter
//     - change the value stored at the *person address
//
// create a value of type person
//     - print out the value
//
// in func main
//     - call "changeMe"
//
// in func main
//     - print out the value
//
// important
//     - to deference a struct use (*value).field
//     - as an exception, if the type of x is a named pointer type
//       and (*x).f is a valid selector expression denoting a field
//       (but not a method), x.f is shorthand for (*x).f

// person struct
type person struct {
	age int
}

func changeMe(p *person) {
	p.age = p.age + 1

	fmt.Println((*p).age)
}

func main() {
	// Create person
	p1 := person{
		30,
	}

	changeMe(&p1)
	changeMe(&p1)
	changeMe(&p1)
}
