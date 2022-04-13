package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	course string
	module int
	clip   int
)

func main() {
	fmt.Println("name is the type", reflect.TypeOf(name))

}
