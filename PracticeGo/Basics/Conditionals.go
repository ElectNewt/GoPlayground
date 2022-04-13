package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 25
	if x > y {
		fmt.Println("x is bigger than Y")
	} else if x == y {
		fmt.Println("y and X are equal")
	} else {
		fmt.Println(" y is bigger")
	}

	if x != y {
		fmt.Println("x and y are not equal")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("iteration number: ", i+1)
	}
}
