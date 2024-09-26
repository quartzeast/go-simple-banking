package domain

import (
	"database/sql"
	"fmt"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func NewCustomerRepositoryDB(client *sql.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client}
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSQL := "SELECT id, name, city, postcode, birth_date, status FROM customers"

	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		return nil, fmt.Errorf("error when querying all customers: %v", err)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.Postcode, &customer.BirthDate, &customer.Status)
		if err != nil {
			return nil, fmt.Errorf("error when scanning customer: %v", err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
