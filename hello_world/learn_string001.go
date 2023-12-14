package main

import "fmt"

func main() {
	fmt.Println("Simple\nstring") //for multi line string
	fmt.Println(`
	dsdsdsd
	sdsdsd
	`)
	fmt.Println("Ʊ")
	fmt.Println("\u2722") // If you use unicode, you need to put an escape sequence
	fmt.Println("Connectivity request")
	fmt.Println(`Connectivity`)
}
