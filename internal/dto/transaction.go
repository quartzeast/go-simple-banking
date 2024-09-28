package dto

import "errors"

type TransactionType string

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerID      string  `json:"-"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (r TransactionRequest) Validate() error {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errors.New("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errors.New("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
