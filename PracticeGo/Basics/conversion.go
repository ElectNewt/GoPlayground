package main

import (
	"fmt"
	"strconv"
)

func main() {

	age := 10
	fmt.Println("the age is", strconv.Itoa(age))

	age_string := "12"
	age_int, err := strconv.Atoi(age_string)

	if err == nil {
		fmt.Println("the converted age is", age_int)
	}
}
