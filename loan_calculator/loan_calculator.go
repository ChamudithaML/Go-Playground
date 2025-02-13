package main

import (
	"fmt"
)

var interest = map[int]int{
	1: 110,
	2: 108,
	3: 105,
}

func main() {
	fmt.Print("Enter the amount: ")
	var amount float32
	fmt.Scan(&amount)

	fmt.Print("Enter the Months: ")
	var months int
	fmt.Scan(&months)

	pLoan := calculateLoan(amount, months)
	fmt.Printf("\nFinal amount:  %.2f", pLoan)
}

func calculateLoan(amount float32, months int) float32 {

	var val int
	if months < 12 {
		val = 1
	} else if months < 24 {
		val = 2
	} else {
		val = 3
	}

	return amount * float32(months) * (float32(interest[val]) / 100)
}

// commits := map[string]int{
//     "rsc": 3711,
//     "r":   2138,
//     "gri": 1908,
//     "adg": 912,
// }

// var interest map[int]int
// interest = make(map[int]int)
// interest[6] = 10
// interest[12] = 20
