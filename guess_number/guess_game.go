package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Game Started. Random value generated between 0 - 10.")
	randVal := rand.Intn(10)

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
