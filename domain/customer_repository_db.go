package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/quartzeast/go-simple-banking/errs"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := c.db.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.BirthDate, &c.Status)
		if err != nil {
			log.Println("Error while querying customers table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
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
			log.Println("Error while querying customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/banking")
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
