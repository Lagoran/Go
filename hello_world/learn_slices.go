package main

import (
	"fmt"
)

func main() {

	names := []string{} //shorthand array assignment
	names = append(names, "Yordan")
	names = append(names, "Rosi", "Eli", "Kalina")

	fmt.Println(names[0])
	fmt.Println(names[2])
	fmt.Println(names)

	/*
		Or we can have another definition of the array, in case we know what the size would be
	*/

	surnames := make([]string, 4)
	surnames[0] = "Strahinov"
	surnames[1] = "Strahinova"
	fmt.Println(names[0], surnames[0])
	fmt.Println(surnames)
}
