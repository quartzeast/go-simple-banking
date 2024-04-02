package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/logger"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := c.db.Query(findAllSQL)
	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.Error("Error while scanning customers table " + err.Error())
		return nil, err
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

func NewCustomerRepository() CustomerRepository {
	db, err := sqlx.Open("mysql", "root:root123@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{
		db: db,
	}
}
