package main

import "fmt"

func array_average(scores [5]float64) float64 {
	var sum float64 = 0
	for _, value := range scores {
		sum += value
	}
	return sum / float64(len(scores))
}

func add_groceries(groceries []string, item string) []string {
	groceries = append(groceries, item)
	return groceries
}

func has_pet(pets map[string]string, pet string) bool {
	_, ok := pets[pet]
	return ok
}

func main() {
	scores := [...]float64{9, 2.7, 18, 32, 4}
	var groceries = []string{"apples", "bananas", "lemons", "milk", "bread"}
	fruits := groceries[0:3]
	var more_fruits []string = append(fruits, "Cherries", "Melon", "Blueberries")
	var pets = map[string]string{"daisy": "bulldog", "periwinkle": "poodle", "harold": "skink", "jeremiah": "snake"}

	fmt.Println("Average Score: ", array_average(scores))
	fmt.Println("Current Grocery List Length: ", len(groceries), "Grocery List Capacity: ", cap(groceries))
	fmt.Println("Grocery List: ", append(add_groceries(groceries, "honey"), "lettuce"))
	fmt.Println("Current Number of Fruits: ", len(fruits), "Fruits Capacity: ", cap(fruits))
	fmt.Println("Fruits: ", fruits)
	fmt.Println(more_fruits)
	fmt.Println("Updated Number of Fruits: ", len(more_fruits), "Updated Fruits Capacity: ", cap(more_fruits))
	fmt.Println("Has Billy?: ", has_pet(pets, "billy"))
	fmt.Println("Has Daisy? :", has_pet(pets, "daisy"))
}
