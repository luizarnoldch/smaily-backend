package domain

type UserRepository interface {
	GetAllUsers() ([]User, error)
}