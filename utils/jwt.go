package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTUtil interface {
	GenerateToken(userID uuid.UUID, role string) (string, time.Time, error)
	ExtractToken(c *fiber.Ctx) (*jwt.Token, error)
}

type jwtUtil struct {
	secret string
}

func NewJWTUtil(secret string) JWTUtil {
	return &jwtUtil{secret: secret}
}

func (j *jwtUtil) GenerateToken(userID uuid.UUID, role string) (string, time.Time, error) {
	expiresAt := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"role":    role,
		"exp":     expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))

	return tokenString, expiresAt, err
}

func (j *jwtUtil) ExtractToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("missing authorization header")
	}

	if len(tokenString) > 7 && strings.ToUpper(tokenString[0:6]) == "BEARER" {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})

	return token, err
}
