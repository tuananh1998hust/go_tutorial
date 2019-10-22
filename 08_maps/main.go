package main

import (
	"fmt"
)

func main() {
	// // Define map
	// emails := make(map[string]string)
	// // Assign
	// emails["Bob"] = "bob@gmail.com"
	// emails["John"] = "john@gmail.com"
	// emails["Sakura"] = "sakura@gmail.com"
	emails := map[string]string{"Bob": "bob@gmail.com", "John": "john@gmail.com", "Sakura": "sakura@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
}
