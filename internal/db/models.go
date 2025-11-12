package db

type Account struct {
	AccountID int64   `db:"account_id" json:"account_id"`
	Balance   float64 `db:"balance" json:"balance"`
}

type Trasnsaction struct {
	ID                   int64   `db:"id" json:"id"`
	SourceAccountID      int64   `db:"source_account_id" json:"source_account_id"`
	DestinationAccountID int64   `db:"destination_account_id" json:"destination_account_id"`
	CreateAt             string  `db:"created_at" json:"created_at"`
	Amount               float64 `db:"amount" json:"amount"`
}
