package main

import "fmt"

func main() {
	birthdays := map[string]string{ //first string is the type of the key; second string is the type of the value of the map
		"Yordan": "18/11/1987",
		"Rosi":   "30/10/1990",
	}

	ages := map[string]int{}

	ages["Yordan"] = 36
	ages["Rosi"] = 33
	ages["Yordan"] = 999 //this way you are overwriting the previous value for this map entry!!

	delete(ages, "Yordan")

	//	fmt.Println(birthdays["Yordan"])
	fmt.Println(birthdays)
	fmt.Println(ages)
}
