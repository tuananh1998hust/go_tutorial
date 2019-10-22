package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign Value
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	fruitArr := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[2:])
}
