package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/utils"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *fiber.Ctx) error {

	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}
	validate := validator.New()

	errValidate := validate.Struct(loginRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	// Check Email
	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credential",
		})
	}
	// Check Password
	validPass := utils.CheckPassword(loginRequest.Password, user.Password)
	if !validPass {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credential",
		})
	}

	// Generate Jwt
	claims := jwt.MapClaims{}

	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	claims["role"] = user.Role

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		log.Println(errToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credential",
		})

	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
