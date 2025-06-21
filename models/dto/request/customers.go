package request

import "time"

type CreateCustomer struct {
	NIK         string    `json:"nik" validate:"required,len=16"`
	FullName    string    `json:"full_name" validate:"required,min=3"`
	Email       string    `json:"email" validate:"required,email"`
	Phone       string    `json:"phone" validate:"required,e164"`
	Address     string    `json:"address" validate:"required,min=5"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type UpdateCustomer struct {
	FullName    string    `json:"full_name" validate:"min=3"`
	Email       string    `json:"email" validate:"email"`
	Phone       string    `json:"phone" validate:"e164"`
	Address     string    `json:"address" validate:"min=5"`
}
