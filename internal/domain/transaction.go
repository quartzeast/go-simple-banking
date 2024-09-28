package domain

import "github.com/quartzeast/go-simple-banking/internal/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	ID              string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionID:   t.ID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
