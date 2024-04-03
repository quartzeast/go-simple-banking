package dto

import (
	"strings"

	"github.com/quartzeast/go-simple-banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() error {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account, amount must be greater than 5000")
	}
	if strings.ToLower(r.AccountType) != "checking" && strings.ToLower(r.AccountType) != "saving" {
		return errs.NewValidationError("Account type must be checking or saving")
	}
	return nil
}
