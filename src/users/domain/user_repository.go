package domain

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserByID() (*User, error)
	UpdateUserByID() (*User, error)
	DeleteUserByID() ([]byte, error)
}