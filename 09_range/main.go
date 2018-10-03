package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// "_" avoids errors when not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	cats := map[string]string{"cappy": "orange", "elbie": "brown", "thora": "tabby"}

	for k, v := range cats {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range cats {
		fmt.Printf("Name: %s\n", k)
	}
}
