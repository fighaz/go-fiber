package route

import (
	"go-fiber/config"
	"go-fiber/handler"
	"go-fiber/middleware"
	"go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	// Auth
	r.Post("login", handler.Login)
	// User
	r.Get("/user", middleware.Auth, handler.GetAllUser)
	r.Get("/user/:id", handler.GetUserById)
	r.Post("/user", handler.Createuser)
	r.Put("/user/:id", handler.UpdateUser)
	r.Delete("/user/:id", handler.DeleteUser)
	// Book
	r.Post("/book", utils.HandleSingleFile, handler.CreateBook)
	// Photo
	r.Post("/gallery", utils.HandleMultipleFile, handler.CreatePhoto)
	r.Delete("/gallery/:id", handler.DeletePhoto)

}
