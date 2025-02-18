package functions

import (
	"fmt"
	"loan_calculator/dto"
)

func CalculateLoan(userData dto.UserData) {

	fmt.Println(userData)

	rateforAmount := calRateForAmount(userData.Amount)

	fmt.Println(rateforAmount)

}

func calRateForAmount(amount float64) float64 {
	var rate float64
	if amount > 100000 {
		rate = 15
	} else if amount > 50000 {
		rate = 10
	} else if amount > 25000 {
		rate = 8.5
	} else {
		rate = 5
	}
	return rate
}
