package main

import "fmt"

func main() {
	var testArray = [3]string{"first", "second", "third"}
	fmt.Println(testArray)

	var testArray2 = []string{"first", "second", "third"}
	fmt.Println(testArray2)
	slice := testArray2[1:3]
	fmt.Println(slice)
	fmt.Println("size of the array: ", cap(testArray))

	testArray2 = append(testArray2, "newone")
	fmt.Println(testArray2)
}
