package application

import (
	"github.com/luizarnoldch/mesapara2_backend/src/users/domain"
)

type UserService interface {
	FindAllUsers() ([]domain.UserResponse, error)
	FindUserByID() (domain.UserResponse, error)
	UpdateUserByID() (domain.UserResponse, error)
	DeleteUserByID() ([]byte, error)
}