package app

import (
	"fmt"

	log "github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/mesapara2_backend/config/env"
	"github.com/luizarnoldch/mesapara2_backend/src/users/application"
	"github.com/luizarnoldch/mesapara2_backend/src/users/domain"
	// "github.com/luizarnoldch/mesapara2_backend/src/users/infrastructure"
)

func GreetHandler(c *fiber.Ctx) error {
	return c.SendString("Hello this is the API for Smaily")
}

// Start initializes the application and starts the server.
//
// It loads the configuration from the .env file and sets up the database connection.
// Then it creates the necessary repository and application instances.
// Finally, it defines the API routes and starts the server.
func Start() {
	config, err := env.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Environment while loading .env: %s", err)
	}

	// dbClient := infrastructure.GetPostgreSQLClient()
	// defer dbClient.Close()

	// repository := domain.NewUserRepositoryLocal(dbClient)
	repository := domain.NewUserRepositoryLocal()
	application := application.NewUserService(repository)
	userController := UserController{application}
	
	app := fiber.New()

	api := app.Group("/api")
	api.Get("/", GreetHandler)

	api_users := api.Group("/user")
	api_users.Get("/", userController.GetAllUsers)
	api_users.Get("/:id", userController.FindUserByID)
	api_users.Put("/", userController.UpdateUserByID)
	api_users.Delete("/:id", userController.DeleteUserByID)

	URL_API := fmt.Sprint(config.MICRO.API.HOST,":",config.MICRO.API.PORT)
	app.Listen(URL_API)
}