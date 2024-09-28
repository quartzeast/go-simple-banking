package domain

import "github.com/quartzeast/go-simple-banking/internal/dto"

type Account struct {
	ID          string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.ID}
}

type AccountRepository interface {
	Save(Account) (*Account, error)
	SaveTransaction(transaction Transaction) (*Transaction, error)
	FindBy(id string) (*Account, error)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
