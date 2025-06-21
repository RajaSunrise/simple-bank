package handlers

import (
	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/services"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	services services.AuthService
}

func NewAuthHandler(services services.AuthService) *AuthHandler {
	return &AuthHandler{services: services}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validate := c.Locals("validate").(*validator.Validate)
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.services.RegisterUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validate := c.Locals("validate").(*validator.Validate)
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.services.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.JSON(response)
}
