package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// You have to explicitly specify fallthrough
	switch x {
	case 10:
		fmt.Printf("%d is 10\n", x)
	case 15:
		fmt.Printf("%d is 15\n", x)
		fallthrough
	default:
		fmt.Printf("%d is not 10 or 15\n", x)
	}
}
