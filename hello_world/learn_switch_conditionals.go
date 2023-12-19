package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}

	ages["Yordan"] = 7

	switch {
	case ages["Yordan"] > 78:
		fmt.Println("Time to rest maan!!")
	case ages["Yordan"] > 33:
		fmt.Println("Time to fun maan!!")
	default:
		fmt.Println("nothing special here about Yordan")
	}

}
