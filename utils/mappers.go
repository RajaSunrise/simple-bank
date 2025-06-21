package utils

import (
	"time"

	"github.com/RajaSunrise/simple-bank/models/dto/response"
	"github.com/RajaSunrise/simple-bank/models/schema"
)

func ToCustomerResponse(c schema.Customer) *response.Customer {
	return &response.Customer{
		ID:          c.ID.String(),
		NIK:         c.NIK,
		FullName:    c.FullName,
		Email:       c.Email,
		Phone:       c.Phone,
		Address:     c.Address,
		DateOfBirth: c.DateOfBirth,
		CreatedAt:   c.CreatedAt,
	}
}

func ToAccountResponse(a schema.Account) *response.Account {
	return &response.Account{
		ID:            a.ID.String(),
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance.String(),
		Currency:      a.Currency,
		Status:        a.Status,
		OpenedDate:    a.OpenedDate,
		CreatedAt:     a.CreatedAt,
	}
}

func ToTransactionResponse(t schema.Transaction) *response.Transaction {
	res := &response.Transaction{
		ID:              t.ID.String(),
		AccountID:       t.AccountID.String(),
		Amount:          t.Amount.String(),
		Type:            t.TransactionType,
		Description:     t.Description,
		ReferenceNumber: t.ReferenceNumber,
		Status:          t.Status,
		TransactionDate: t.TransactionDate,
	}

	if t.RelatedAccountID.Valid {
		res.RelatedAccountID = t.RelatedAccountID.UUID.String()
	}

	return res
}

func ToAuthResponse(user schema.AuthUser, token string, expires time.Time) *response.AuthResponse {
	return &response.AuthResponse{
		UserID:      user.ID.String(),
		CustomerID:  user.CustomerID.String(),
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: token,
		ExpiresAt:   expires,
	}
}
