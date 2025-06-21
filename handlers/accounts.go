package handlers

import (
	"strconv"

	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/services"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AccountHandler struct {
	services services.AccountService
}

func NewAccountHandler(services services.AccountService) *AccountHandler {
	return &AccountHandler{services: services}
}

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var req request.CreateAccount
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

	account, err := h.services.CreateAccount(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(account)
}

func (h *AccountHandler) GetAccount(c *fiber.Ctx) error {
	accountID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid account ID",
		})
	}

	account, err := h.services.GetAccount(c.Context(), accountID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	return c.JSON(account)
}

func (h *AccountHandler) GetAccountStatement(c *fiber.Ctx) error {
	accountID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid account ID",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	statement, err := h.services.GetAccountStatement(c.Context(), accountID, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get account statement",
		})
	}

	return c.JSON(statement)
}

func (h *AccountHandler) TransferFunds(c *fiber.Ctx) error {
	var req request.TransferFunds
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

	response, err := h.services.TransferFunds(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}
