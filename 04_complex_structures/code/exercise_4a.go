package main

import "fmt"

// func average(first_float, second_float, third_float float32) float32 {
// 	return (first_float + second_float + third_float) / 3
// }

func average(nums ...float32) float32 {
	var sum float32 = 0
	for _, num := range nums {
		sum += num
	}
	return sum / float32(len(nums))
}

func main() {
	fmt.Println(average(32.7, 40.0, 12.6, 14.2, 73.5))
}
