package main

import (
	"fmt"
)

func main() {
	ids := []int{1, 3, 6, 8, 9}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Range Map
	emails := map[string]string{"Bob": "bob@gmail.com", "John": "john@gmail.com", "Sakura": "sakura@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
