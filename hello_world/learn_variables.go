package main

import (
	"fmt"
)

func main() {
	var A int = 16
	var B, ok = "yes", true
	var _ int = 33 //we can use the blank identifier "_" to point a var that we don't intent to use in case we don't want to the delete it
	myInt := 11    //shorthand variable assignmest
	var name string

	name = "Name is what??"

	fmt.Println("The first variable is:", A)
	fmt.Println("The second variable is:", B)
	fmt.Println("The third variable is:", ok)
	fmt.Println("Value of myInt variable is:", myInt)
	fmt.Println("Value of name variable is:", name)
}
