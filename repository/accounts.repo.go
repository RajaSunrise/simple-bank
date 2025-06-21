package repository

import (
	"context"

	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountRepository interface {
	Create(ctx context.Context, account *schema.Account) error
	GetByID(ctx context.Context, id uuid.UUID) (*schema.Account, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (*schema.Account, error)
	GetByCustomerID(ctx context.Context, customerID uuid.UUID) ([]schema.Account, error)
	UpdateBalance(ctx context.Context, accountID uuid.UUID, amount decimal.Decimal) error
	UpdateStatus(ctx context.Context, accountID uuid.UUID, status string) error
}

