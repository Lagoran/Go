package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Learn some Go Math")
	fmt.Println("Addition:", 1+2)
	fmt.Println("Substraction:", 1-2)
	fmt.Println("Multiplication:", 1*2)
	fmt.Println("Division:", 1/2)
	fmt.Println("Division with float:", 1.0/3)

	fmt.Println("Power function:", math.Pow(5, 5))
	fmt.Printf("Remainder of the devision is: %.1f", math.Remainder(100, 30))
}
