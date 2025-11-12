package repositories

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/thebhavesh02/internal-transfers-service/internal/db"
)

type AccountRepository interface {
	GetAccountBalance(ctx context.Context, accountId int64) (string, string, error)
	CreateAccount(ctx context.Context, accountId *int64, initialBalance *string) error
}

type accountRepository struct {
	queries db.Queries
}

func NewAccountRepository(queries db.Queries) AccountRepository {
	return &accountRepository{queries: queries}
}

func (ar *accountRepository) GetAccountBalance(ctx context.Context, accountId int64) (string, string, error) {
	if accountId <= 0 {
		return "", "", errors.New("invalid accountId")
	}
	account, err := ar.queries.GetAccount(ctx, accountId)
	if err != nil {
		log.Default().Println("Error fetching account: ", err)
		return "", "", err
	}
	return strconv.FormatInt(account.AccountID, 10), strconv.FormatFloat(account.Balance, 'f', 2, 64), err
}

func (ar *accountRepository) CreateAccount(ctx context.Context, accountId *int64, initialBalance *string) error {
	if accountId == nil {
		return errors.New("accountId is nil")
	}
	if initialBalance == nil {
		return errors.New("initialBalance is nil")
	}

	balance, err := strconv.ParseFloat(*initialBalance, 64)
	if err != nil {
		return err
	}

	if err := ar.queries.CreateAccount(ctx, *accountId, balance); err != nil {
		return err
	}
	return nil
}
