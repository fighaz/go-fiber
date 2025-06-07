package route

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.GetAllUser)
	r.Get("/user/:id", handler.GetUserById)
	r.Post("/user", handler.Createuser)
	r.Put("/user/:id", handler.UpdateUser)
	r.Delete("/user/:id", handler.DeleteUser)
}
