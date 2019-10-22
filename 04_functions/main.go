package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func total(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Tuan Anh"))
	fmt.Println(total(2, 3))
}
