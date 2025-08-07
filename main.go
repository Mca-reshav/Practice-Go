// main.go

/*
Def :
    Go is a statically typed, compiled programming language designed at Google.
    It was designed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson.

Features :
    - Simplicity and readability
    - Fast compilation
    - Built-in concurrency support
    - Strong support for modularity, testing, and tooling
*/

package main

import "fmt" // standard package lib : format

func main() {
	showVar()

	const addFnVal = 5 // constant expression only
	fmt.Println("Const result of 2+3 =", addFnVal)

	result := add(2, 3) // use result for function
	fmt.Println("Function result of 2+3 =", result)

	fmt.Println("Hello, Go!")
	fmt.Print("No newline here.")
	fmt.Printf("My age is %d\n", 25) // just like C language

	// IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("IIFE in Go")
	}()

	// use single quote only

	const sal int = 40000
	conditions(sal)
	conditions2(sal)

	loopIter(5, "FOR")
	loopIter(4, "FOR_EACH")
	loopIter(2, "FOR_OF")

	str()
	formatName("   john doe  ")
	staticAry()

	slices()

	obj()

	goFunc()
	waitGroup()

	chann()
	prodConsumer()

}
