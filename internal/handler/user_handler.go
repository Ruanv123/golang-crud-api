package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruanv123/api-go-crud/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.userService.Profile(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
