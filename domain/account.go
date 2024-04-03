package domain

import "github.com/quartzeast/go-simple-banking/dto"

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountID: a.AccountID,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, error)
}
