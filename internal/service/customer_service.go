package service

import (
	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/dto"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, error)
	GetCustomer(id string) (*dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0, cap(customers))
	for _, customer := range customers {
		response = append(response, customer.ToDTO())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, error) {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	response := customer.ToDTO()
	return &response, nil
}
