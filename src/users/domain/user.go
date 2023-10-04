package domain

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *UserRequest) ToUser() User {
    return User{
        Email:    request.Email,
        Password: request.Password,
    }
}

type User struct {
	ID string `json:"id,omitempty"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (user *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

type UserResponse struct {
	ID string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}