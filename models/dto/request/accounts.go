package request

import "time"

type CreateAccount struct {
	CustomerID    string    `json:"customer_id" validate:"required,uuid4"`
	BranchCode    string    `json:"branch_code" validate:"required"`
	Balance       string    `json:"balance" validate:"required"`
	AccountType   string    `json:"account_type" validate:"required,oneof=SAVINGS CHECKING BUSINESS LOAN"`
	AccountNumber string    `json:"account_number" validate:"required,numeric"`
	OpenedDate    time.Time `json:"opened_date" validate:"required"`
}

type TransferFunds struct {
	FromAccountID string `json:"from_account_id" validate:"required,uuid4"`
	ToAccountID   string `json:"to_account_id" validate:"required,uuid4"`
	Amount        string `json:"amount" validate:"required,decimal"`
	Description   string `json:"description" validate:"required,min=5"`
}
