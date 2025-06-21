package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID               uuid.UUID       `db:"transaction_id"`
	AccountID        uuid.UUID       `db:"account_id"`
	RelatedAccountID uuid.NullUUID   `db:"related_account_id"`
	Amount           decimal.Decimal `db:"amount"`
	TransactionType  string          `db:"transaction_type"`
	Description      string          `db:"description"`
	ReferenceNumber  string          `db:"reference_number"`
	Status           string          `db:"status"`
	TransactionDate  time.Time       `db:"transaction_date"`
	CreatedAt        time.Time       `db:"created_at"`
}
