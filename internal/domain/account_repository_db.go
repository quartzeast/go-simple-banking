package domain

import (
	"context"
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

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, error) {
	// starting the database transaction
	ctx := context.Background()
	tx, err := d.client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, fmt.Errorf("Error while starting a new transaction for bank account transaction: %w", err)
	}

	// inserting bank account transaction
	insertSQL := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`
	result, err := tx.Exec(insertSQL, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		return nil, fmt.Errorf("Error while inserting bank account transaction: %w", err)
	}

	// updating bank account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE id = ?", t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE id = ?", t.Amount, t.AccountID)
	}

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error while updating bank account balance: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("Error while committing bank account transaction: %w", err)
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Error while getting last inserted transaction id: %w", err)
	}

	account, err := d.FindBy(t.AccountID)
	if err != nil {
		return nil, fmt.Errorf("Error while getting last inserted transaction id: %w", err)
	}

	t.ID = strconv.FormatInt(transactionID, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindBy(accountId string) (*Account, error) {
	var account Account
	getAccountSQL := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	row := d.client.QueryRow(getAccountSQL, accountId)
	err := row.Scan(&account.ID, &account.CustomerID, &account.OpeningDate, &account.AccountType, &account.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &account, nil
}
