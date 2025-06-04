package handlers

import (
	"context"

	"github.com/RajaSunrise/simple-bank/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	Queries   *sqlc.Queries
}

func NewUsersHandler(queries *sqlc.Queries) *UserHandlers {
	return &UserHandlers{Queries: queries}
}

// CreateUsers Handler
func (h *UserHandlers) CreateUsers(c *fiber.Ctx) error {
	var reqUser struct {
		Username		string 	`json:"username"`
		Hash_password	string	`json:"hash_password"`
		Email			string	`json:"email"`
	}
	if err := c.BodyParser(&reqUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Invalid Body Parser",
			"Error": err.Error(),
		})
	}

	user, err := h.Queries.CreateUser(context.Background(), sqlc.CreateUserParams{
		Username: reqUser.Username,
		PasswordHash: reqUser.Hash_password,
		Email: reqUser.Email,
		Role: "users",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Internal Server Error",
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Status": "Succes Create Users",
		"Users": user,
	})
}

func GetUsersByAccount()  {

}

func GetUsersByUsername()  {

}
