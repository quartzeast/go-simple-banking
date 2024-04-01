package service

import "github.com/quartzeast/go-simple-banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// 4. 业务逻辑实现 Service port（这是 Primary port），并且以来 repository
// 所以首先创建 primary port，然后创建实现，然后注入依赖（repository）
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
