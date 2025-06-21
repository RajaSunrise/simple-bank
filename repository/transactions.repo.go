package repository

import (
	"context"

	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)


type TransactionRepository interface {
	Create(ctx context.Context, transaction *schema.Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*schema.Transaction, error)
	GetByAccountID(ctx context.Context, accountID uuid.UUID, page, pageSize int) ([]schema.Transaction, error)
	GetByReferenceNumber(ctx context.Context, refNumber string) (*schema.Transaction, error)
	Transfer(ctx context.Context, fromAccountID, toAccountID uuid.UUID, amount decimal.Decimal, description string) (*schema.Transaction, error)
}
