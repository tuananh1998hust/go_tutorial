package main

import (
	"fmt"
)

// Switch Func
func line(color string) {
	// Switch
	switch color {
	case "red":
		fmt.Println("STOP")
	case "yellow":
		fmt.Println("Slow")
	case "green":
		fmt.Println("GO")
	}
}

func main() {
	x := 5
	y := 10

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal than %d", x, y)
	} else {
		fmt.Printf("%d is more than %d", x, y)
	}

	line("red")
}
