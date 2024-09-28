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

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	var err error
	var rows *sql.Rows

	if status == "" {
		findAllSQL := "SELECT id, name, city, postcode, birth_date, status FROM customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT id, name, city, postcode, birth_date, status FROM customers where status = ?"
		rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		return nil, fmt.Errorf("error when querying all customers: %w", err)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.City, &customer.Postcode, &customer.BirthDate, &customer.Status)
		if err != nil {
			return nil, fmt.Errorf("error when scanning customer: %w", err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindByID(id string) (*Customer, error) {
	findByIDSQL := "SELECT id, name, city, postcode, birth_date, status FROM customers WHERE id = ?"

	var customer Customer
	err := d.client.QueryRow(findByIDSQL, id).
		Scan(&customer.ID, &customer.Name, &customer.City, &customer.Postcode, &customer.BirthDate, &customer.Status)
	if err != nil {
		return nil, fmt.Errorf("error when scanning customer: %w", err)
	}
	return &customer, nil
}
