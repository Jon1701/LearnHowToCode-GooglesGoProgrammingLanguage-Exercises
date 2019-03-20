package main

import "fmt"

// Write a program which prints a number in decimal, binary, and hex

func main() {
	registry := 1701

	fmt.Printf("The number %v in decimal notation is: %d\n", registry, registry)
	fmt.Printf("The number %v in binary notation is: %b\n", registry, registry)
	fmt.Printf("The number %v in hexadecimal notation is: %#x\n", registry, registry)
}
