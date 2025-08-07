/*
Goroutine

A lightweight thread managed by the Go runtime.
Allows you to run functions concurrently (like async/await in JS).
Starts with the keyword: go
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func goFunc() {
	go sayHello() // run concurrently

	// Give time for goroutine to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function done.")
}

func printMsg(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

// Wait group :To wait for multiple goroutines to finish:
func waitGroup() {
	var wg sync.WaitGroup

	wg.Add(2) // we’ll run 2 goroutines

	go printMsg("Hello", &wg)
	go printMsg("World", &wg)

	wg.Wait() // block until all done
	fmt.Println("All done")
}
