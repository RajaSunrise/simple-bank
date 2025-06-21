package response

import "time"

type Account struct {
	ID            string    `json:"id"`
	AccountNumber string    `json:"account_number"`
	AccountType   string    `json:"account_type"`
	Balance       string    `json:"balance"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	OpenedDate    time.Time `json:"opened_date"`
	CreatedAt     time.Time `json:"created_at"`
}

type AccountStatement struct {
	Account      Account       `json:"account"`
	Transactions []Transaction `json:"transactions"`
}
