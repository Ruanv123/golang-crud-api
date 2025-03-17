package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token não fornecido",
			})
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token inválido",
			})
		}

		tokenString := bearerToken[1]

		secretKey := []byte("minha-chave-secreta")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Método de assinatura inválido")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			log.Println("Erro de autenticação:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token inválido ou expirado",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Erro ao decodificar o token",
			})
		}

		userID := int(claims["user_id"].(float64))
		c.Locals("userID", userID)

		return c.Next()
	}
}
