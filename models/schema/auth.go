package schema

import (
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID           uuid.UUID `db:"user_id"`
	CustomerID   uuid.UUID `db:"customer_id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	LastLogin    time.Time `db:"last_login"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
