package domain

type Customer struct {
	Id        string
	Name      string
	City      string
	ZipCode   string
	BirthDate string
	Status    string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
