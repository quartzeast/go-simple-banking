package service

import (
	"time"

	"github.com/quartzeast/go-simple-banking/domain"
	"github.com/quartzeast/go-simple-banking/dto"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountID:   "",
		CustomerID:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToDTO()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return DefaultAccountService{repo: repo}
}
