package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			log.Printf("Error: %v", err)

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"error":   err.Error(),
				"message": "Ocorreu um erro interno. Por favor, tente novamente mais tarde.",
			})
		}
		return nil
	}
}
