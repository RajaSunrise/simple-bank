package schema

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID `db:"customer_id"`
	NIK         string    `db:"nik"`
	FullName    string    `db:"full_name"`
	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	Address     string    `db:"address"`
	DateOfBirth time.Time `db:"date_of_birth"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
