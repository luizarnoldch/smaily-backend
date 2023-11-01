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

// FindUserByID retrieves a user by their ID.
//
// c is a pointer to a fiber.Ctx object that represents the HTTP request context.
// It is used to handle the response and any potential errors.
//
// The function returns an error if there is an issue retrieving the user.
// Otherwise, it returns the retrieved user as a JSON response.
func (controller *UserController) FindUserByID(c *fiber.Ctx) error {
	user, err := controller.service.FindUserByID()
	if err != nil {
		c.JSON(err)
	}
	return c.JSON(user)
}

// UpdateUserByID updates a user by their ID.
//
// c is a pointer to a fiber.Ctx object that represents the HTTP request context.
// It is used to handle the response and any potential errors.
//
// The function returns an error if there is an issue updating the user.
// Otherwise, it returns the updated user as a JSON response.
func (controller *UserController) UpdateUserByID(c *fiber.Ctx) error {
	user, err := controller.service.UpdateUserByID()
	if err != nil {
		c.JSON(err)
	}
	return c.JSON(user)
}

// DeleteUserByID deletes a user by their ID.
//
// c is a pointer to a fiber.Ctx object that represents the HTTP request context.
// It is used to handle the response and any potential errors.
//
// The function returns an error if there is an issue deleting the user.
// Otherwise, it returns the deleted user as a JSON response.
func (controller *UserController) DeleteUserByID(c *fiber.Ctx) error {
	user, err := controller.service.DeleteUserByID()
	if err != nil {
		c.JSON(err)
	}
	return c.JSON(user)
}
