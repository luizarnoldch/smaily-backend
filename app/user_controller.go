package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/mesapara2_backend/src/users/application"
)

type UserController struct {
	service application.UserService
}

// GetAllUsers retrieves all users from the UserController.
//
// c is a pointer to a fiber.Ctx object that represents the HTTP request context.
// It is used to handle the response and any potential errors.
//
// The function returns an error if there is an issue retrieving the users.
// Otherwise, it returns the retrieved users as a JSON response.
func (controller *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := controller.service.FindAllUsers()

	if err != nil {
		c.JSON(err)
	}

	return c.JSON(users)
}