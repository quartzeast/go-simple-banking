package domain

import (
	"database/sql"
	"fmt"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sql.DB
}

func NewAccountRepositoryDB(client *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client}
}

func (d AccountRepositoryDB) Save(a Account) (*Account, error) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	stmt, err := d.client.Prepare(sqlInsert)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare sql: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to exec sql: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id from new account: %w", err)
	}

	a.ID = strconv.FormatInt(id, 10)
	return &a, nil
}
