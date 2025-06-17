package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/asset/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("cover")

	if errFile != nil {
		log.Println("Error file = ", errFile)
	}

	var fileName *string

	if file != nil {
		// Validasi Tipe File
		errCheckContentType := CheckContentType(file, "image/jpg", "image/png", "image/jpeg")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}
		// Save File
		fileName = &file.Filename
		currentTime := time.Now().Format(time.RFC3339)
		extFile := filepath.Ext(*fileName)
		generatedFileName := fmt.Sprintf("image-%s%s", currentTime, extFile)
		fileName = &generatedFileName
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/asset/%s", *fileName))

		if errSaveFile != nil {
			log.Println("Failed saving file into directory")
		}

	}

	if fileName != nil {
		ctx.Locals("fileName", *fileName)
	} else {
		ctx.Locals("fileName", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()

	if errForm != nil {
		log.Println("Error Read Multipart Form ,Error = ", errForm)
	}
	files := form.File["photos"]

	var fileNames []string

	for i, file := range files {
		var fileName string

		if file != nil {
			// Validasi Tipe File
			errCheckContentType := CheckContentType(file, "image/jpg", "image/png", "image/jpeg")
			if errCheckContentType != nil {
				return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"message": errCheckContentType.Error(),
				})
			}

			fileName = fmt.Sprintf("%d-%s", i, file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/asset/%s", fileName))

			if errSaveFile != nil {
				log.Println("Failed saving file into directory")
			}
		}
		if fileName != "" {
			fileNames = append(fileNames, fileName)
		}

	}
	ctx.Locals("fileNames", fileNames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("Failed To remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed To remove file")
			return err
		}
	}

	return nil
}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("Type File Not Allowed")

	} else {
		return errors.New("Content Not Found ")
	}
}
