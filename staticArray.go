package main

import "fmt"

func staticAry() {
	// Declare an array of 5 integers
	var nums [5]int
	fmt.Println("Empty array:", nums)

	// Assign values
	nums[0] = 10
	nums[1] = 20
	fmt.Println("Updated array:", nums)

	// Initialize with values
	letters := [3]string{"A", "B", "C"}
	fmt.Println("Letters:", letters)

	// Get length
	fmt.Println("Length:", len(letters))

	// Looping through array
	for i, v := range letters {
		fmt.Println("Index:", i, "Value:", v)
	}

	// try to conversion
	// var ary1 [4]int;
	// var ary2 [3]int;
	//ary2 = ary1; // prevent its conversion

	// behavior by value
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // [1 2 3] — modified!
}
