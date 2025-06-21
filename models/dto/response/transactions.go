package response

import "time"

type Transaction struct {
	ID               string    `json:"id"`
	AccountID        string    `json:"account_id"`
	RelatedAccountID string    `json:"related_account_id,omitempty"`
	Amount           string    `json:"amount"`
	Type             string    `json:"type"`
	Description      string    `json:"description"`
	ReferenceNumber  string    `json:"reference_number"`
	Status           string    `json:"status"`
	TransactionDate  time.Time `json:"transaction_date"`
}

type TransferResponse struct {
	TransactionID     string `json:"transaction_id"`
	ReferenceNumber   string `json:"reference_number"`
	Status            string `json:"status"`
	FromAccountNumber string `json:"from_account_number"`
	ToAccountNumber   string `json:"to_account_number"`
	Amount            string `json:"amount"`
	Timestamp         string `json:"timestamp"`
}
