package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/mesapara2_backend/config/env"
)

func Start() {
	config, err := env.LoadConfig(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	app := fiber.New()

	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	URL_API := fmt.Sprint(config.MICRO.API.HOST,":",config.MICRO.API.PORT)
	app.Listen(URL_API)
}