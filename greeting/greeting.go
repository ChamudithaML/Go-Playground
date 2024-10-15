package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter Name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello %s !!!", name)
}
