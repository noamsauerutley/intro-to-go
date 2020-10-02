package main

import "fmt"

func average(first_float, second_float, third_float float32) float32 {
	return (first_float + second_float + third_float) / 3
}

func main() {
	fmt.Println(average(32.7, 40.0, 12.6))
}
