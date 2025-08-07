// Loop iterator in GO : for, while, forEach, range (for of), range with map (for in)
package main

import (
	"fmt"
	"time"
)

func loopIter(iter int, loopType string) {
	if iter > 10 {
		fmt.Println("Please reduce the iter size")
	} else if iter <= 0 {
		fmt.Println("Invalid iter passed")
	} else {
		fmt.Println("Loop iter started at", time.Now())

		switch loopType {
		case "FOR":
			fmt.Println("Running FOR loop")
			for i := 0; i < iter; i++ {
				fmt.Println(i)
			}
		case "WHILE":
			fmt.Println("Running WHILE loop")
			i := 0
			for i < iter {
				fmt.Println(i)
				i++
			}
		case "FOR_EACH":
			fmt.Println("Running FOR_EACH loop")
			num := make([]int, iter)
			for i := range num {
				num[i] = i
			}
			for _, val := range num {
				fmt.Println(val)
			}
		case "FOR_OF", "FOR_IN":
			fmt.Println("Running RANGE loop (like FOR_OF/FOR_IN)")
			numMap := map[string]int{}
			for i := 0; i < iter; i++ {
				key := fmt.Sprintf("key%d", i)
				numMap[key] = i
			}
			for key, val := range numMap {
				fmt.Printf("%s => %d\n", key, val)
			}
		default:
			fmt.Println("Invalid loop type")
		}

		fmt.Println("Loop iter terminated at", time.Now())
	}
}
