package main

import (
	math "fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := math.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}
