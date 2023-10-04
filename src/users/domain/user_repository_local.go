package domain

// import "github.com/jmoiron/sqlx"

type UserRepositoryLocal struct {
	// client *sqlx.DB
}

// GetAllUsers returns all users from the local user repository.
//
// It does not take any parameters.
// It returns a slice of UserResponse and an error.
func (repo *UserRepositoryLocal) GetAllUsers() ([]User, error) {
	users := []User{
		{ID: "1", Email: "user1@example.com", Password: "password1"},
		{ID: "2", Email: "user2@example.com", Password: "password2"},
	}

	return users, nil
}

// NewUserRepositoryLocal initializes a new instance of UserRepositoryLocal.
//
// It takes a client parameter of type *sqlx.DB and returns an instance of UserRepository.
func NewUserRepositoryLocal() *UserRepositoryLocal {
	// return &UserRepositoryLocal{
	// 	client: client,
	// }
	return &UserRepositoryLocal{}
}