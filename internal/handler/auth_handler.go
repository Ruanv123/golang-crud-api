package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruanv123/api-go-crud/internal/model"
	"github.com/ruanv123/api-go-crud/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "dados invalidos",
		})
	}

	token, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "crendemciais invalidas",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Dados inválidos",
		})
	}

	// Registrar usuário
	newUser := &model.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	err := h.authService.Register(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao registrar usuário",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuário registrado com sucesso",
	})
}
