package services

import (
	"context"

	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/models/dto/response"
	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/RajaSunrise/simple-bank/repository"
	"github.com/RajaSunrise/simple-bank/utils"
	"github.com/google/uuid"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, req request.CreateCustomer) (*response.Customer, error)
	GetCustomer(ctx context.Context, customerID uuid.UUID) (*response.CustomerWithAccounts, error)
	UpdateCustomer(ctx context.Context, customerID uuid.UUID, req request.UpdateCustomer) (*response.Customer, error)
}

type customerService struct {
	customerRepo repository.CustomerRepository
	accountRepo  repository.AccountRepository
}

func NewCustomerService(
	customerRepo repository.CustomerRepository,
	accountRepo repository.AccountRepository,
) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
		accountRepo:  accountRepo,
	}
}

func (s *customerService) CreateCustomer(ctx context.Context, req request.CreateCustomer) (*response.Customer, error) {
	customer := &schema.Customer{
		NIK:         req.NIK,
		FullName:    req.FullName,
		Email:       req.Email,
		Phone:       req.Phone,
		Address:     req.Address,
		DateOfBirth: req.DateOfBirth,
	}

	if err := s.customerRepo.Create(ctx, customer); err != nil {
		return nil, err
	}

	return utils.ToCustomerResponse(*customer), nil
}

func (s *customerService) GetCustomer(ctx context.Context, customerID uuid.UUID) (*response.CustomerWithAccounts, error) {
	customer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	accounts, err := s.accountRepo.GetByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	res := &response.CustomerWithAccounts{
		Customer: *utils.ToCustomerResponse(*customer),
	}

	for _, account := range accounts {
		res.Accounts = append(res.Accounts, *utils.ToAccountResponse(account))
	}

	return res, nil
}

func (s *customerService) UpdateCustomer(ctx context.Context, customerID uuid.UUID, req request.UpdateCustomer) (*response.Customer, error) {
	customer, err := s.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.FullName != "" {
		customer.FullName = req.FullName
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Address != "" {
		customer.Address = req.Address
	}

	if err := s.customerRepo.Update(ctx, customer); err != nil {
		return nil, err
	}

	return utils.ToCustomerResponse(*customer), nil
}
