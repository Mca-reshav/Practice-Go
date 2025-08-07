// Conditional statement : if, else if, switch
package main

import "fmt"

// Function to return condition based on salary
func conditions(sal int) {
	var cond string

	if sal < 0 {
		cond = "Invalid Salary"
	} else if sal >= 0 && sal < 15000 {
		cond = "Poor"
	} else if sal >= 15000 && sal < 40000 {
		cond = "Below Medium"
	} else if sal >= 40000 && sal < 75000 {
		cond = "Medium"
	} else {
		cond = "Rich"
	}

	fmt.Println("Salary Condition:", cond)
}

func conditions2(sal int) {
	var cond string

	switch {
	case sal < 0:
		cond = "Invalid Salary"
	case sal < 25000:
		cond = "Poor"
	case sal < 50000:
		cond = "Below Medium"
	case sal < 95000:
		cond = "Medium"
	default:
		cond = "Rich"
	}

	fmt.Println("Salary Condition 2:", cond)
}
