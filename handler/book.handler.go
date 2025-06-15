package handler

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(ctx *fiber.Ctx) error {
	book := new(request.BookRequest)

	if err := ctx.BodyParser(book); err != nil {
		return err
	}
	validate := validator.New()

	errValidate := validate.Struct(book)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Handle File Upload
	var fileNameString string
	fileName := ctx.Locals("fileName")

	if fileName == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "Image cover is required",
		})
	} else {
		fileNameString = fmt.Sprintf("%v", fileName)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  fileNameString,
	}

	errCreate := database.DB.Create(&newBook).Error

	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
