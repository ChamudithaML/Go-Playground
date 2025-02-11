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
	// fmt.Printf("\n************** Random value %d: **************\n", randVal)

	answer, count, mid := -5, 0, 0

	for answer != randVal {
		count++
		fmt.Printf("\nEnter value between %d - %d: ", startVal, endVal)
		fmt.Scan(&answer)

		if answer != randVal {
			fmt.Println("\nTry again")
		} else {
			fmt.Println("\nCorrect")
		}

		mid = (startVal + endVal) / 2

		if count%3 == 0 {
			if answer != randVal {
				if randVal > mid {
					startVal = mid
					fmt.Printf("Random value greater than %d", mid)
				} else {
					endVal = mid
					fmt.Printf("Random value less than %d", mid)
				}
			}
			gap = endVal - startVal
		}

		fmt.Printf("\ngap %d. start val %d. end val %d. mid %d. count %d", gap, startVal, endVal, mid, count)
	}
}
