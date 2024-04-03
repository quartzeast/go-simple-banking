package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/quartzeast/go-simple-banking/errs"
	"github.com/quartzeast/go-simple-banking/logger"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (a AccountRepositoryDb) Save(account Account) (*Account, error) {

	insertSQL := `INSERT INTO accounts 
		(account_id, customer_id, opening_date, account_type, amount, status)
		VALUES (?, ?, ?, ?, ?, ?)`

	result, err := a.db.Exec(insertSQL,
		account.AccountID,
		account.CustomerID,
		account.OpeningDate,
		account.AccountType,
		account.Amount,
		account.Status,
	)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account.AccountID = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{db: db}
}
