package service

import (
	"github.com/quartzeast/go-simple-banking/domain"
	"github.com/quartzeast/go-simple-banking/dto"
)

// CustomerService 定义了一个 contract
// 表示中心的业务逻辑可以对外提供 Customer 服务，表明我可以提供 GetAllCustomers 和 GetCustomer 两个服务
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, error)
}

// 4. CustomerService 业务逻辑实现，依赖 repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, error) {
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
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDTO())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, error) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDTO()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repository}
}
