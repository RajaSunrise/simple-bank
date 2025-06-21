package repository

import (
	"context"

	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/google/uuid"
)


type CustomerRepository interface {
	Create(ctx context.Context, customer *schema.Customer) error
	GetByID(ctx context.Context, id uuid.UUID) (*schema.Customer, error)
	GetByEmail(ctx context.Context, email string) (*schema.Customer, error)
	GetByNIK(ctx context.Context, nik string) (*schema.Customer, error)
	Update(ctx context.Context, customer *schema.Customer) error
	Delete(ctx context.Context, id uuid.UUID) error
}
