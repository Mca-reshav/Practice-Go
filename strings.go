package main

import (
	"fmt"
	"strings"
)

func str() {
	// Declaration
	str := "Hello, Go!"
	rawStr := `Line 1
Line 2`

	fmt.Println("Normal:", str)
	fmt.Println("Raw:", rawStr)

	// Length
	fmt.Println("Length:", len(str)) // in bytes

	// Accessing characters (as bytes)
	fmt.Println("First character (byte):", str[0])     // 72 (ASCII of 'H')
	fmt.Printf("First character (char): %c\n", str[0]) // H

	// Slicing
	fmt.Println("Slice:", str[1:5]) // Hello

	// Concatenation
	newStr := str + " How are you?"
	fmt.Println("Concat:", newStr)

	// Strings package examples
	fmt.Println("ToUpper:", strings.ToUpper(str))
	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))
	fmt.Println("Split by comma:", strings.Split(str, ","))
}

func formatName(name string) {
	trimmed := strings.TrimSpace(name)     // Step 1
	titleCase := strings.Title(trimmed)    // Step 2 (deprecated but simple)
	words := strings.Split(titleCase, " ") // Step 3

	if len(words) >= 2 {
		fmt.Println("First Name:", words[0])
		fmt.Println("Last Name:", words[1])
	} else {
		fmt.Println("Full Name:", titleCase)
	}

	str := "hello"
	for i, char := range str {
		fmt.Print(i, char, " "+string(char), "\n")
		fmt.Printf("Index: %d, Char: %c\n", i, char)
	}

}
