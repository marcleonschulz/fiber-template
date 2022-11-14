package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/fiber-template/config"
	"github.com/marcleonschulz/fiber-template/internal/database"
	"github.com/marcleonschulz/fiber-template/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	err = database.Init()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	router.SetupRoutes(app)
}
