package service

import "github.com/quartzeast/go-simple-banking/domain"

// CustomerService 定义了一个 合约 contract
// 表示中心的业务逻辑可以对外提供 Customer 服务，表明我可以提供 GetAllCustomers 和 GetCustomer 两个服务
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

// 4. 业务逻辑实现 CustomerService port（这是 Primary port），并且依赖 repository
// 所以首先创建 primary port，然后创建实现，然后注入依赖（repository）
type DefaultCustomerService struct {
	repo domain.CustomerRepository // 这也是一个合约，表示 DefaultCustomerService 依赖 某个 Repository 提供的能力，但不去考虑具体的实现
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repository}
}

// 在 业务领域通过接口定义 port，有两种类型，一个是 primary port，另一个是 secondary port
// primary port 定义了 service，secondary port 定义了 repository
// service 端口是业务逻辑对外提供的能力，repository 端口定义了外部存储要遵循的规范

// 可以把业务逻辑看作一台机器，这台机器对外提供了 service 端口，表示如果外部可以使用这台机器的哪些能力，就可从插入这个端口获取这些能力
// 所以外部系统依赖 service 端口

// repository 定义的是外部存储如果要接入这台机器，要遵循哪些规范，拿这个项目来说，外部存储要实现 Customerrepository 接口定义的方法
