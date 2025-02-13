package dto

type UserData struct {
	Name            string
	Amount          float64
	IsSafe          bool
	RepaymentPeriod string // long-term, short-term
	Months          int
	LoanType        string // Personal, Home, Vehicle, Business
}
