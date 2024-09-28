package dto

import (
	"fmt"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() error {
	if r.Amount < 500 {
		return fmt.Errorf("to open a new account you need to deposit at least 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return fmt.Errorf("account type should be checking or saving")
	}
	return nil
}
