package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number: ", num)
}

// Prints and adds multiple numbers, returning the sum
func Add(nums ...int) int {
	sum := 0
	for _, value := range nums {
		printNum(value)
		sum += value
	}
	return sum
}
