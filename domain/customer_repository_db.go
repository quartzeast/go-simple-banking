package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/logger"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = c.db.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = c.db.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (c CustomerRepositoryDb) ById(id string) (*Customer, error) {

	findByIdSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := c.db.QueryRow(findByIdSQL, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.ZipCode, &customer.BirthDate, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return CustomerRepositoryDb{
		db: db,
	}
}
