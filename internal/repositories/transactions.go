package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/thebhavesh02/internal-transfers-service/internal/db"
)

type TransactionsRepository interface {
	CreateTransaction(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error
}

type transactionRepository struct {
	queries db.Queries
}

func NewTransactionRepository(queries db.Queries) TransactionsRepository {
	return &transactionRepository{
		queries: queries,
	}
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {
	amountFloat, err := strconv.ParseFloat(*amount, 64)
	if err != nil {
		return err
	}
	amountMicros := int64(amountFloat * 10000)

	return tr.queries.ExecTx(ctx, func(q db.Queries) error {
		sourceAccount, err := q.GetAccount(ctx, *sourceAccountId)
		if err != nil {
			return err
		}

		destinationAccount, err := q.GetAccount(ctx, *destinationAccountId)
		if err != nil {
			return err
		}

		if amountMicros > sourceAccount.Balance {
			return errors.New("insufficient funds")
		}

		sourceAccounNewBalance := sourceAccount.Balance - amountMicros
		destinationAccountNewBalance := destinationAccount.Balance + amountMicros

		err = q.UpdateAccount(ctx, *sourceAccountId, sourceAccounNewBalance)
		if err != nil {
			return err
		}

		err = q.UpdateAccount(ctx, *destinationAccountId, destinationAccountNewBalance)
		if err != nil {
			return err
		}

		if err := q.CreateTransaction(ctx, *sourceAccountId, *destinationAccountId, amountMicros); err != nil {
			return err
		}

		return nil
	})
}
