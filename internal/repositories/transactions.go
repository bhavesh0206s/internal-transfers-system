package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/thebhavesh02/internal-transfers-service/internal/db"
)

type TransactionsRepository interface {
	CreateTransaction(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error
	CreateAccountDetails(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error
}

type transactionRepository struct {
	queries db.Queries
}

func NewTransactionRepository(queries db.Queries) TransactionsRepository {
	return &transactionRepository{
		queries: queries,
	}
}

func (tr *transactionRepository) CreateAccountDetails(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {
	amountFloat, err := strconv.ParseFloat(*amount, 64)

	if err != nil {
		return err
	}

	sourceAccount, err := tr.queries.GetAccount(ctx, *sourceAccountId)

	if err != nil {
		return err
	}

	destinationAccount, err := tr.queries.GetAccount(ctx, *destinationAccountId)

	if err != nil {
		return err
	}

	if amountFloat > sourceAccount.Balance {
		return errors.New("insufficient funds")
	}

	sourceAccounNewBalance := sourceAccount.Balance - amountFloat
	destinationAccountNewBalance := destinationAccount.Balance + amountFloat

	err = tr.queries.UpdateAccount(ctx, *sourceAccountId, sourceAccounNewBalance)

	if err != nil {
		return err
	}

	err = tr.queries.UpdateAccount(ctx, *destinationAccountId, destinationAccountNewBalance)

	if err != nil {
		return err
	}

	return nil
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {

	amountFloat, err := strconv.ParseFloat(*amount, 64)

	if err != nil {
		return err
	}

	if err := tr.CreateAccountDetails(ctx, sourceAccountId, destinationAccountId, amount); err != nil {
		return err
	}

	if err := tr.queries.CreateTransaction(ctx, *sourceAccountId, *destinationAccountId, amountFloat); err != nil {
		return err
	}

	return nil
}
