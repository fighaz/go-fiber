package middleware

import (
	"go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": " Unauthenticated",
		})
	}
	_, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": " Unauthenticated",
		})
	}

	role := claims["role"].(string)

	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": " Forbidden Access",
		})
	}

	ctx.Locals("userInfo", claims)

	return ctx.Next()
}
