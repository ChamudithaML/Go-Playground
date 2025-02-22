package main

import (
	"fmt"
	"loan_calculator/dto"
	"loan_calculator/functions"
	"strings"
)

func main() {

	fmt.Print("Enter the Name: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("Enter the amount: ")
	var amount float64
	fmt.Scan(&amount)

	fmt.Print("Enter months: ")
	var months int
	fmt.Scan(&months)

	fmt.Print("Enter loan type: ")
	var loanType string
	fmt.Scan(&loanType)

	fmt.Print("Employeed? Yes or No: ")
	var isEmployeed string
	fmt.Scan(&isEmployeed)

	fmt.Print("\nProcessing...\n")

	// fmt.Print("Enter Repayment Period: ")
	// var repaymentPeriod string
	// fmt.Scan(&repaymentPeriod)

	var isSafe bool
	if strings.ToLower(isEmployeed) == "yes" {
		isSafe = true
	} else {
		isSafe = false
	}

	var repaymentPeriod string
	if months > 60 {
		repaymentPeriod = "long-term"
	} else if months > 24 {
		repaymentPeriod = "mid-term"
	} else {
		repaymentPeriod = "short-term"
	}

	userData := dto.UserData{
		Name:            name,
		Amount:          amount,
		IsSafe:          isSafe,
		RepaymentPeriod: repaymentPeriod,
		Months:          months,
		LoanType:        loanType,
	}

	functions.CalculateLoan(userData)
}
