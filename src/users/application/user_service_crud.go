package application

import (
	"encoding/json"
	"sync"

	"github.com/luizarnoldch/mesapara2_backend/src/users/domain"
)

type UserServiceCRUD struct {
	repo domain.UserRepository
}

// FindAllUsers returns a list of all users and any error encountered.
//
// It does not take any parameters.
// It returns a slice of domain.UserResponse and an error.
func (service *UserServiceCRUD) FindAllUsers() ([]domain.UserResponse, error) {
	users, err := service.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var response []domain.UserResponse

	wg.Add(len(users))

	for _, user := range users {
		go func(u domain.User) {
			defer wg.Done()

			userResponse := u.ToUserResponse()

			mu.Lock()
			response = append(response, userResponse)
			mu.Unlock()
		}(user)
	}

	wg.Wait()

	return response, nil
}

// FindUserByID retrieves a user by their ID.
//
// It returns a domain.UserResponse and an error.
func (service *UserServiceCRUD) FindUserByID() (domain.UserResponse, error){
	user, err := service.repo.GetUserByID()
	if err != nil {
		return domain.UserResponse{}, err
	}

	response := user.ToUserResponse()

	return response, nil
}

// UpdateUserByID updates a user by their ID.
//
// It returns a domain.UserResponse and an error.
func (service *UserServiceCRUD) UpdateUserByID() (domain.UserResponse, error){
	user, err := service.repo.UpdateUserByID()
	if err != nil {
		return domain.UserResponse{}, err
	}

	response := user.ToUserResponse()
	
	return response, nil
}

// DeleteUserByID deletes a user by their ID.
//
// This function does not take any parameters.
// It returns a byte slice and an error.
func (service *UserServiceCRUD) DeleteUserByID() ([]byte, error){
	user, err := service.repo.DeleteUserByID()

	if err != nil {
		return []byte{}, err
	}
	return  json.Marshal(&user)
}


// NewUserService creates a new instance of UserServiceCRUD.
//
// It takes a domain.UserRepository as a parameter and returns a pointer to UserServiceCRUD.
func NewUserService(repo domain.UserRepository) *UserServiceCRUD {
	return &UserServiceCRUD{
		repo: repo,
	}
}