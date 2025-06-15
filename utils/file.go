package utils

import (
	"fmt"
	"log"
	"os"

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
		fileName = &file.Filename

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
