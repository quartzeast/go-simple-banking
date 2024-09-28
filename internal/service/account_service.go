package service

import (
	"time"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/dto"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerID:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDTO()
	return &response, nil
}
