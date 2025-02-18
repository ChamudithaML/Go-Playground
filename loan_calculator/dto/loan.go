package dto

type Loan struct {
	Type        string  `json:"Type"`
	Client      string  `json:"Client"`
	Amount      float64 `json:"Amount"`
	Interest    float64 `json:"Interest"`
	IsSafe      bool    `json:"IsSafe"`
	SafetyAdded bool    `json:"SafetyAdded"`
}
