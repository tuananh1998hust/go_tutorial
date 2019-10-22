package main

import (
	"fmt"
)

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := add()
	for i := 0; i < 5; i++ {
		fmt.Println(sum(i))
	}
}
