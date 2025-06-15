package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/utils"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CreatePhoto(ctx *fiber.Ctx) error {
	photo := new(request.PhotoRequest)

	if err := ctx.BodyParser(photo); err != nil {
		return err
	}
	validate := validator.New()

	errValidate := validate.Struct(photo)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Validasi File Upload
	fileNames := ctx.Locals("fileNames")
	var newPhotos []entity.Photo
	if fileNames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "Image Photo is required",
		})
	} else {
		fileNamesData := fileNames.([]string)
		for _, fileName := range fileNamesData {
			newPhoto := entity.Photo{
				Image:      fileName,
				CategoryId: photo.CategoryId,
			}

			errCreate := database.DB.Create(&newPhoto).Error

			if errCreate != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"message": "Failed to store some data",
				})
			} else {
				newPhotos = append(newPhotos, newPhoto)
			}
		}

	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newPhotos,
	})
}

func DeletePhoto(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")
	var photo entity.Photo
	err := database.DB.First(&photo, "id = ?", photoId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Photo Not Found",
		})
	}
	// Handle Delete
	errDelete := database.DB.Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to delete data",
		})
	} else {
		// Handle Delete file
		errDeleteFile := utils.HandleRemoveFile(photo.Image)
		if errDeleteFile != nil {
			log.Println("Failed to Remove File from project")
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    "Photo has been deleted",
	})
}
