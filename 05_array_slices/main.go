package main

import "fmt"

func main() {
	var fruitArray [2]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	vegeArray := [2]string{"broccoli", "sprout"}

	fruitSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(fruitArray)
	fmt.Println(vegeArray)
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])
	fmt.Println(len(fruitSlice))
}
