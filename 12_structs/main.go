package main

import (
	"fmt"
	"strconv"
)

// Person is a person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// value receiver (not changing anything)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ", and I am " + strconv.Itoa(p.age) + " years old"
}

// pointer receiver (changing something)
func (p *Person) hasBirthday() {
	p.age++
}

// pointer receiver (changing something)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Jane", lastName: "Smith", city: "Seattle", gender: "F", age: 40}
	person2 := Person{"Jane", "Smith", "Seattle", "F", 40}
	fmt.Println(person1, person2)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person1.getMarried("Jensen")
	fmt.Println(person1.greet())
}
