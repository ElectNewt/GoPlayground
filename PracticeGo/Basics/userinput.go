package main

import "fmt"

func main() {

	var age int

	fmt.Print("please write your age: ")
	fmt.Scanf("%d\n", &age)
	fmt.Println("next year you will have ", age+1, " years")

	fmt.Print("please write your name: ")
	var name string
	fmt.Scanf("%s\n", &name)

	fmt.Println("your name is", name, "and your age is", age)

	formatedString := fmt.Sprintf("your name is %s and your age is %d", name, age)
	fmt.Println(formatedString)
}
