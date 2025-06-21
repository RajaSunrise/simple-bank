package handlers

import (
	"github.com/RajaSunrise/simple-bank/models/dto/request"
	"github.com/RajaSunrise/simple-bank/services"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CustomersHandlers struct {
	services services.CustomerService
}

func NewCustomersHandlers(services services.CustomerService) *CustomersHandlers {
	return &CustomersHandlers{services: services}
}

func (h *CustomersHandlers) CreateCustomer(c *fiber.Ctx) error {
	var req request.CreateCustomer
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

	customer, err := h.services.CreateCustomer(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create customer",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}

func (h *CustomersHandlers) GetCustomer(c *fiber.Ctx) error {
	customerID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	customer, err := h.services.GetCustomer(c.Context(), customerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.JSON(customer)
}

func (h *CustomersHandlers) UpdateCustomer(c *fiber.Ctx) error {
	customerID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid customer ID",
		})
	}

	var req request.UpdateCustomer
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

	customer, err := h.services.UpdateCustomer(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update customer",
		})
	}

	return c.JSON(customer)
}
