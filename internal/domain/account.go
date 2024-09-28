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
}
