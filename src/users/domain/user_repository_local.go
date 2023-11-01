package domain

import "encoding/json"

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
		{ClientID: 1, ClientName: "Alice", ClientDirection: "123 Main St", ClienteEmail: "alice@example.com", ClientePassword: "password123", ClienteTelephone: "123-456-7890"},
		{ClientID: 2, ClientName: "Bob", ClientDirection: "456 Elm St", ClienteEmail: "bob@example.com", ClientePassword: "bobpassword", ClienteTelephone: "987-654-3210"},
		{ClientID: 3, ClientName: "Charlie", ClientDirection: "789 Oak St", ClienteEmail: "charlie@example.com", ClientePassword: "charliepass", ClienteTelephone: "111-222-3333"},
	}

	return users, nil
}

// GetUserByID retrieves a user from the local user repository based on the user ID.
// It creates a new User struct with sample user data and returns a pointer to it along with a nil error.
func (repo *UserRepositoryLocal) GetUserByID() (*User, error) {
	user := User{ClientID: 1, ClientName: "Alice", ClientDirection: "123 Main St", ClienteEmail: "alice@example.com", ClientePassword: "password123", ClienteTelephone: "123-456-7890"}

	return &user, nil
}

// UpdateUserByID updates a user in the local user repository based on the user ID.
// It creates a new User struct with updated sample user data and returns a pointer to it along with a nil error.
func (repo *UserRepositoryLocal) UpdateUserByID() (*User, error) {
	user := User{ClientID: 1, ClientName: "Bob", ClientDirection: "456 Elm St", ClienteEmail: "bob@example.com", ClientePassword: "bobpassword", ClienteTelephone: "987-654-3210"}

	return &user, nil
}

// DeleteUserByID deletes a user from the local user repository based on the user ID.
// It marshals the string "User deleted" into a byte slice and returns it along with a nil error.
func (repo *UserRepositoryLocal) DeleteUserByID() ([]byte, error) {
	return json.Marshal("User deleted")
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