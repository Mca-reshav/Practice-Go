/*
Slices

- A slice is a dynamically-sized, flexible view into an array.
- Unlike arrays, slices don't have fixed length.
- They are reference types (like JS arrays).
*/

package main

import "fmt"

func slices() {
	// Declare a slice
	var nums []int                     // nil slice (no memory yet)
	letters := []string{"A", "B", "C"} // initialized slice

	// Print
	fmt.Println("Empty:", nums)
	fmt.Println("Letters:", letters)

	// Length
	fmt.Println("Length:", len(letters))

	// Append (like JS push)
	letters = append(letters, "D")
	fmt.Println("Appended:", letters)

	// Slicing an array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // Includes index 1, excludes index 4
	fmt.Println("Sliced:", slice)

	// make an array
	nums1 := make([]int, 3, 5) // length = 3, capacity = 5
	nums1[0] = 10
	fmt.Println(nums1)                  // [10 0 0]
	fmt.Println(len(nums1), cap(nums1)) // length and memory allocation
	// you can discard cap if not specific

	// looping
	for i, v := range letters {
		fmt.Println("Index:", i, "Value:", v)
	}

	// behavior by reference
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // [100 2 3] — modified!

}
