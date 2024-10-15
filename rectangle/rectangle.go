package main

import (
	"fmt"
)

func main() {
	var length, width float32
	fmt.Print("Enter width: ")
	fmt.Scan(&width)

	fmt.Print("Enter length: ")
	fmt.Scan(&length)

	fmt.Printf("Area of rectangle is: %f", length*width*.5)
	// fmt.Print(length * width * .5)
}
