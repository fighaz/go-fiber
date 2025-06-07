package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Init DB
	database.DatabaseInit()
	// Migration
	migration.RunMigration()
	// Init APP
	app := fiber.New()

	// Init Route

	route.RouteInit(app)

	app.Listen(":3000")
}
