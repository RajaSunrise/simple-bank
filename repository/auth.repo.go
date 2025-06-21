package repository

import (
	"context"

	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/google/uuid"
)


type AuthRepository interface {
	CreateUser(ctx context.Context, user *schema.AuthUser) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*schema.AuthUser, error)
	GetUserByUsername(ctx context.Context, username string) (*schema.AuthUser, error)
	GetUserByCustomerID(ctx context.Context, customerID uuid.UUID) (*schema.AuthUser, error)
	UpdateUser(ctx context.Context, user *schema.AuthUser) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
