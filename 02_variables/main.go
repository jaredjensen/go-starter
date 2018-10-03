package main

import "fmt"

// Variables
// var name string = "Bob" // only specify type if it can't be inferred
var name = "Bob"
var age = 40

const isTired = true // can't redefine const later

func main() {

	// can use shorthand declarations only inside functions
	x := "Test"

	// can do multiple assignments at once
	y, z := "1", "2"

	// print values to console
	fmt.Println(name, age, x, y, z)
	fmt.Printf("%d\n", age)

	// show variable type
	fmt.Printf("%T\n", age)
}
