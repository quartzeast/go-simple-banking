package service

import (
	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/dto"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0, cap(customers))
	for _, customer := range customers {
		response = append(response, customer.ToDTO())
	}
	return response, nil
}
