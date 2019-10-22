package main

import (
	"fmt"
	"math"

	"github.com/tuananh1998hust/go_tutorial/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Abs(-4))
	fmt.Println(strutil.Reverse("Hello World"))
}
