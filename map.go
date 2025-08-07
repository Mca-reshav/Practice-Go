/*
Maps : which are the Go equivalent of JavaScript objects or Python dictionaries.

A map is a key-value data structure.
Keys must be comparable types (e.g., string, int).
Values can be any type.
*/

package main

import "fmt"

func obj() {
	// Declare and initialize
	person := map[string]string{
		"name": "Alice",
		"city": "Mumbai",
	}
	//Map keys must be of comparable types.
	//[]int (a slice) is not comparable, so it cannot be used as a key

	// Accessing
	fmt.Println("Name:", person["name"])

	// Adding/updating a key
	person["job"] = "Engineer"
	person["city"] = "Delhi" // updates

	// Deleting
	delete(person, "job")

	// Looping over map
	for key, value := range person {
		fmt.Println(key, "→", value)
	}

	// check if exist
	val, exists := person["age"]
	if exists {
		fmt.Println("Age is", val)
	} else {
		fmt.Println("Age not found")
	}

	// zero value
	m := map[string]int{}
	fmt.Println(m["unknown"]) // 0 (not error)
	//When a key doesn’t exist in a map, Go returns the zero value for that type

}
