package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a
	c := 20

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)

	// Change value pointer
	*b = 10
	fmt.Println(b)
	fmt.Println(a, *b)

	b = &c
	fmt.Println(b)
	fmt.Println(a, *b, c)
}
