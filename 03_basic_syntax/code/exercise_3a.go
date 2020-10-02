package main

import "fmt"

func main() {
	var sentence = "The dog had a lovely day playing in the park."
	for i, letter := range sentence {
		if i%2 != 0 {
			fmt.Println(string(letter))
		}
	}
}
