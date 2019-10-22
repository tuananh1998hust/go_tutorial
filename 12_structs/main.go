package main

import (
	"fmt"
)

// Person :
type Person struct {
	name string
	age  int
}

func (p Person) greeting() string {
	return "Hello " + p.name
}

func main() {
	person := Person{name: "Tuan Anh", age: 20}

	fmt.Println(person)
	fmt.Println(person.greeting())
	person.age++
	fmt.Println(person)
}
