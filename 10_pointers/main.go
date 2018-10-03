package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	*b = 6

	fmt.Println(a, b)
}
