package main

import "fmt"

func main() {
	var name, place string
	var time int
	var validity bool

	fmt.Println("Name:")
	fmt.Scan(&name)
	fmt.Println("Residence:")
	fmt.Scan(&place)
	fmt.Println("Years of residence:")
	fmt.Scan(&time)
	fmt.Println("The weather is nice: (true/false)")
	fmt.Scan(&validity)
	fmt.Printf("My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t.\n", name, place, time, validity)
}
