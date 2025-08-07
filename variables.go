// Variables and Constants
// Go is statically typed, but support type inference

// Declare mutable and constant

package main

import "fmt"

var name string = "John"

// age := 30  throw error, its a short variables declaration are allowed only inside a function
const pi = 3.14

func showVar() {
	age := 30
	fmt.Println(name, age, pi)
}
