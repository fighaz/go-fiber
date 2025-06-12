package route

import (
	"go-fiber/config"
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("login", handler.Login)
	r.Get("/user", middleware.Auth, handler.GetAllUser)
	r.Get("/user/:id", handler.GetUserById)
	r.Post("/user", handler.Createuser)
	r.Put("/user/:id", handler.UpdateUser)
	r.Delete("/user/:id", handler.DeleteUser)
}
