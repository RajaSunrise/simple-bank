package services

import (
	"context"
	"errors"

	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/models/dto/response"
	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/RajaSunrise/simple-bank/repository"
	"github.com/RajaSunrise/simple-bank/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountService interface {
	CreateAccount(ctx context.Context, req request.CreateAccount) (*response.Account, error)
	GetAccount(ctx context.Context, accountID uuid.UUID) (*response.Account, error)
	GetAccountStatement(ctx context.Context, accountID uuid.UUID, page, pageSize int) (*response.AccountStatement, error)
	TransferFunds(ctx context.Context, req request.TransferFunds) (*response.TransferResponse, error)
}

type accountService struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func NewAccountService(
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
) AccountService {
	return &accountService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *accountService) CreateAccount(ctx context.Context, req request.CreateAccount) (*response.Account, error) {
	customerID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	balance, err := decimal.NewFromString(req.Balance)
	if err != nil {
		balance = decimal.Zero
	}

	account := &schema.Account{
		CustomerID:    customerID,
		BranchCode:    req.BranchCode,
		AccountNumber: req.AccountNumber,
		AccountType:   req.AccountType,
		Balance:       balance,
		OpenedDate:    req.OpenedDate,
	}

	if err := s.accountRepo.Create(ctx, account); err != nil {
		return nil, err
	}

	return utils.ToAccountResponse(*account), nil
}

func (s *accountService) GetAccount(ctx context.Context, accountID uuid.UUID) (*response.Account, error) {
	account, err := s.accountRepo.GetByID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return utils.ToAccountResponse(*account), nil
}

func (s *accountService) GetAccountStatement(ctx context.Context, accountID uuid.UUID, page, pageSize int) (*response.AccountStatement, error) {
	account, err := s.accountRepo.GetByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	transactions, err := s.transactionRepo.GetByAccountID(ctx, accountID, page, pageSize)
	if err != nil {
		return nil, err
	}

	res := &response.AccountStatement{
		Account: *utils.ToAccountResponse(*account),
	}

	for _, t := range transactions {
		res.Transactions = append(res.Transactions, *utils.ToTransactionResponse(t))
	}

	return res, nil
}

func (s *accountService) TransferFunds(ctx context.Context, req request.TransferFunds) (*response.TransferResponse, error) {
	fromAccountID, err := uuid.Parse(req.FromAccountID)
	if err != nil {
		return nil, errors.New("invalid from account ID")
	}

	toAccountID, err := uuid.Parse(req.ToAccountID)
	if err != nil {
		return nil, errors.New("invalid to account ID")
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil || amount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("invalid amount")
	}

	// Check sufficient balance
	fromAccount, err := s.accountRepo.GetByID(ctx, fromAccountID)
	if err != nil {
		return nil, errors.New("source account not found")
	}

	if fromAccount.Balance.LessThan(amount) {
		return nil, errors.New("insufficient balance")
	}

	// Execute transfer in transaction
	transaction, err := s.transactionRepo.Transfer(ctx, fromAccountID, toAccountID, amount, req.Description)
	if err != nil {
		return nil, err
	}

	// Get account numbers for response
	toAccount, err := s.accountRepo.GetByID(ctx, toAccountID)
	if err != nil {
		return nil, err
	}

	return &response.TransferResponse{
		TransactionID:     transaction.ID.String(),
		ReferenceNumber:   transaction.ReferenceNumber,
		Status:            transaction.Status,
		FromAccountNumber: fromAccount.AccountNumber,
		ToAccountNumber:   toAccount.AccountNumber,
		Amount:            amount.String(),
		Timestamp:         transaction.TransactionDate.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
