package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Game Started....")
	startVal, endVal, gap := 0, 0, 0

	for !(gap > 0) {
		fmt.Println("\nEnter the range to generate random number...")
		fmt.Print("\nStarting value: ")
		fmt.Scan(&startVal)
		fmt.Print("Ending value: ")
		fmt.Scan(&endVal)

		fmt.Printf("\nstart val %d. end val %d \n", startVal, endVal)
		gap = endVal - startVal
		if gap <= 0 {
			fmt.Println("\nEnd Value should greater than Starting value")
		}
	}

	// random val between the start value and end value
	randVal := rand.Intn(gap) + startVal

	answer, count := -5, 0

	for answer != randVal {
		count++
		fmt.Print("\nEnter value between 1 - 10: ")
		fmt.Scan(&answer)

		if answer != randVal {
			fmt.Println("\nTry again")
		} else {
			fmt.Println("\nCorrect")
		}

		if count%3 == 0 {
			mid := randVal / 2
			if answer != randVal {
				if randVal > mid {
					fmt.Printf("Random value greater than %d", mid)
				} else {
					fmt.Printf("Random value less than %d", mid)
				}
			}
		}
	}
}
