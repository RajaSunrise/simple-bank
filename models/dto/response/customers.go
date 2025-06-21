package response

import "time"

type Customer struct {
	ID          string    `json:"id"`
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

type CustomerWithAccounts struct {
	Customer
	Accounts []Account `json:"accounts"`
}
