package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// don't need to declare type of all parameters if they're the same
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Bob"))
	fmt.Println(getSum(3, 4))
}
