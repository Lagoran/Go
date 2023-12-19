package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}

	ages["Yordan"] = 77

	if ages["Yordan"] < 65 {
		fmt.Println("Yordan is not of retirement age.")
	} else if ages["Yordan"] > 65 {
		fmt.Println("Yordan needs to retire asap")
	}

}
