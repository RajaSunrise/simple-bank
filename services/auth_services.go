package services

import (
	"context"
	"errors"
	"time"

	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/models/dto/response"
	"github.com/RajaSunrise/simple-bank/models/schema"
	"github.com/RajaSunrise/simple-bank/repository"
	"github.com/RajaSunrise/simple-bank/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(ctx context.Context, req request.RegisterUser) (*response.AuthResponse, error)
	Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error)
}

type authService struct {
	authRepo repository.AuthRepository
	jwtUtil  utils.JWTUtil
}

func NewAuthService(
	authRepo repository.AuthRepository,
	jwtUtil utils.JWTUtil,
) AuthService {
	return &authService{
		authRepo: authRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *authService) RegisterUser(ctx context.Context, req request.RegisterUser) (*response.AuthResponse, error) {
	customerID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		return nil, errors.New("invalid customer ID")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &schema.AuthUser{
		CustomerID:   customerID,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
	}

	if err := s.authRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, expiresAt, err := s.jwtUtil.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		UserID:      user.ID.String(),
		CustomerID:  user.CustomerID.String(),
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}

func (s *authService) Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := s.authRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	user.LastLogin = time.Now()
	if err := s.authRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, expiresAt, err := s.jwtUtil.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		UserID:      user.ID.String(),
		CustomerID:  user.CustomerID.String(),
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}
