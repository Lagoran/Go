package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}

	ages["Yordan"] = 36
	ages["Rosi"] = 33
	ages["Elena"] = 7
	ages["Kalina"] = 1

	for name, age := range ages {
		switch age {
		case 36:
			fmt.Println(name, ", time to rest!!")
		case 33:
			fmt.Println(name, ", happy momming!!")
		default:
			fmt.Println(fmt.Sprintln("Nothing special here about %s's current age", name))
		}

	}
}
