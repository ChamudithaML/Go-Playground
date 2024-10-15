package main

import (
	"fmt"
)

func main() {
	basicCal()

}

func basicCal() {
	var value1, value2 float32
	fmt.Print("Enter value 1: ")
	fmt.Scan(&value1)
	fmt.Print("Enter value 1: ")
	fmt.Scan(&value2)

	var sum float32

	var operator string
	fmt.Print("Enter operator: ")
	fmt.Scan(&operator)

	if operator == "+" {
		sum = addition(value1, value2)
	} else if operator == "-" {
		sum = substraction(value1, value2)
	} else if operator == "*" {
		sum = multiplication(value1, value2)
	} else if operator == "/" {
		sum = division(value1, value2)
	}

	fmt.Printf("Sum of values: %.2f", sum)
}

func addition(a float32, b float32) float32 {
	return a + b
}

func substraction(a float32, b float32) float32 {
	return a - b
}

func division(a float32, b float32) float32 {
	return a / b
}

func multiplication(a float32, b float32) float32 {
	return a * b
}
