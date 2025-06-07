package handler

import (
	"go-fiber/database"
	"go-fiber/helper"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/model/response"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func Createuser(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	validate := validator.New()

	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	// Check Email
	if user.Email != "" {
		check := checkEmail(user.Email)
		if check == false {
			return ctx.Status(402).JSON(fiber.Map{
				"message": "Email already used",
			})
		}
	}
	errCreate := database.DB.Create(&newUser).Error

	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user entity.User
	userId := ctx.Params("id")

	UserRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(UserRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	// Check User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	// Update User
	user.Name = helper.CheckString(UserRequest.Name, user.Name)
	user.Email = helper.CheckString(UserRequest.Email, user.Email)
	user.Phone = helper.CheckString(UserRequest.Phone, user.Phone)
	user.Address = helper.CheckString(UserRequest.Address, user.Phone)

	// Check Email
	if UserRequest.Email != "" {
		check := checkEmail(user.Email)
		if check == false {
			return ctx.Status(402).JSON(fiber.Map{
				"message": "Email already used",
			})
		}
	}

	if errUpdate := database.DB.Save(&user).Error; errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user,
	})

}
func DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	errDelete := database.DB.Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to delete data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    "User has been deleted",
	})
}
func checkEmail(email string) bool {
	var user entity.User
	result := database.DB.First(&user, "email = ?", email)

	if result.RowsAffected > 0 {
		return false
	}
	return true

}
