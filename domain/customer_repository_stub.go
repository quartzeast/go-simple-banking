package domain

import "github.com/quartzeast/go-simple-banking/errs"

// 3. 创建一个 adpater 实现定义的端口（port）
type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll(status string) ([]Customer, error) {
	return c.customers, nil
}

func (c CustomerRepositoryStub) ById(id string) (*Customer, error) {
	for _, customer := range c.customers {
		if customer.Id == id {
			return &customer, nil
		}
	}
	return nil, errs.NewNotFoundError("Customer not found")
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{Id: "1", Name: "John", City: "New York", ZipCode: "10001", BirthDate: "2000-01-01", Status: "active"},
		{Id: "2", Name: "Jane", City: "Paris", ZipCode: "75001", BirthDate: "2000-01-01", Status: "active"},
		{Id: "3", Name: "Bob", City: "Berlin", ZipCode: "10115", BirthDate: "2000-01-01", Status: "active"},
	}
	return CustomerRepositoryStub{customers: customers}
}
