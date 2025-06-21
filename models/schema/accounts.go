package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID            uuid.UUID       `db:"account_id"`
	CustomerID    uuid.UUID       `db:"customer_id"`
	BranchCode    string          `db:"branch_code"`
	AccountNumber string          `db:"account_number"`
	AccountType   string          `db:"account_type"`
	Balance       decimal.Decimal `db:"balance"`
	Currency      string          `db:"currency"`
	Status        string          `db:"status"`
	OpenedDate    time.Time       `db:"opened_date"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}
