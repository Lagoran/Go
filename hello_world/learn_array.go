package main

import (
	"fmt"
)

func main() {

	names := [4]string{"Yordan", "Rosi", "Eli", "Kali"} //shorthand array assignment
	var surnames [4]string                              //normal variable initialization

	surnames[0] = "Strahinov"
	surnames[1] = "Strahinova"
	surnames[2] = "Strahinova"
	surnames[3] = "Strahinova"

	fmt.Println("Array entries:", names)
	fmt.Println("Array entries:", names, surnames)
	fmt.Println("Third element of names array is:", names[3] == "nil")
	//fmt.Println("Is there a fourth element in the names array?:", names[4] == "nil")

}
