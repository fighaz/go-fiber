package handler

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"log"

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
	file, errFile := ctx.FormFile("cover")

	if errFile != nil {
		log.Println("Error file = ", errFile)
	}

	var fileName string

	if file != nil {
		fileName = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/asset/%s", fileName))

		if errSaveFile != nil {
			log.Println("Failed saving file into directory")
		}
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  fileName,
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
