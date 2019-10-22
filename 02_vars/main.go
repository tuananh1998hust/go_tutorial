package main

import (
	"fmt"
)

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Tuan Anh"
	var age int8 = 20
	var isCool = true
	// Shorthand
	name, email := "Tuan Anh", "tuananh1998.hust@gmail.com"
	fmt.Println(name, age, email)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
