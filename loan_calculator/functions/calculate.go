package functions

import (
	"fmt"
	"loan_calculator/dto"
)

func CalculateLoan(userData dto.UserData) {

	rateforAmount := calRateForAmount(userData.Amount)
	rateForTerm := calRateForTerm(userData.RepaymentPeriod)

	fmt.Println(rateforAmount, rateForTerm)

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

func calRateForTerm(term string) float64 {
	if term == "long-term" {
		return 5
	} else if term == "mid-term" {
		return 3
	} else {
		return 2.5
	}
}
