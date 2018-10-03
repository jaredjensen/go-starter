package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	// add values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")

	fmt.Println(emails)

	// shorthand declare/define
	cats := map[string]string{"cappy": "orange", "elbie": "brown", "thora": "tabby"}

	fmt.Println(cats)
}
