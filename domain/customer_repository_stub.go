package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", City: "New York", ZipCode: "10001", BirthDate: "2000-01-01", Status: "active"},
		{Id: "2", Name: "Jane", City: "Paris", ZipCode: "75001", BirthDate: "2000-01-01", Status: "active"},
		{Id: "3", Name: "Bob", City: "Berlin", ZipCode: "10115", BirthDate: "2000-01-01", Status: "active"},
	}
	return CustomerRepositoryStub{customers: customers}
}
