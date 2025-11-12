package db

import (
	"context"
	"database/sql"
	"errors"
)

type Queries interface {
	CreateAccount(ctx context.Context, accountId int64, initialBalance float64) error
	GetAccount(ctx context.Context, accountId int64) (*Account, error)
	CreateTransaction(ctx context.Context, sourceAccountId int64, destinationAccountId int64, amount float64) error
	UpdateAccount(ctx context.Context, accountId int64, initialBalance float64) error
}

type queries struct {
	db *sql.DB
}

func NewQueries(db *sql.DB) Queries {
	return &queries{db: db}
}

func (q *queries) CreateAccount(ctx context.Context, accountId int64, initialBalance float64) error {
	_, err := q.db.ExecContext(ctx, "INSERT INTO accounts (account_id, balance) VALUES ($1, $2)", accountId, initialBalance)
	return err
}

func (q *queries) UpdateAccount(ctx context.Context, accountId int64, balance float64) error {
	_, err := q.db.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE account_id = $2", balance, accountId)
	return err
}

func (q *queries) GetAccount(ctx context.Context, accountId int64) (*Account, error) {
	var account Account

	err := q.db.QueryRowContext(ctx,
		"SELECT account_id, balance FROM accounts WHERE account_id = $1",
		accountId,
	).Scan(&account.AccountID, &account.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("account not found with the given accountId")
		}
		return nil, err
	}

	return &account, nil
}

func (q *queries) CreateTransaction(ctx context.Context, sourceAccountId int64, destinationAccountId int64, amount float64) error {
	_, err := q.db.ExecContext(ctx, "INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)", sourceAccountId, destinationAccountId, amount)
	return err
}
