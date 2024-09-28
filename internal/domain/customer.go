package domain

import "github.com/quartzeast/go-simple-banking/internal/dto"

type CustomerRepository interface {
	FindAll(status string) ([]Customer, error)
	FindByID(id string) (*Customer, error)
}

type Customer struct {
	ID        string
	Name      string
	City      string
	Postcode  string
	BirthDate string
	Status    string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:        c.ID,
		Name:      c.Name,
		City:      c.City,
		Postcode:  c.Postcode,
		BirthDate: c.BirthDate,
		Status:    c.statusAsText(),
	}
}
