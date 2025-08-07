/* 	GO does not support arrow function as like Nodejs
 	But support named and self calling / anonymous function
	Unlike Nodejs GO returns multiple values
*/

// you must declare add fn in main
package main

func add(x int, y int) int {
	return x + y
}

// ref iife in main.go

// func() {
//     fmt.Println("IIFE in Go")
// }()
