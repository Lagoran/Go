package main

import "fmt"

func main() {
	i := 0
	for i < 122 {
		if i%2 == 0 {
			i++
			continue
		} else if i == 117 {
			break
		}

		fmt.Println("We're counting:", i)
		i++
	}
}
