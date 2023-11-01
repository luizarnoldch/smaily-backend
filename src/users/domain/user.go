package domain

type UserRequest struct {
	ClienteEmail    string `json:"cliente_email"`
	ClientePassword string `json:"cliente_password"`
}

type User struct {
	ClientID         int    `db:"cliente_id"`
	ClientName       string `db:"cliente_nombre"`
	ClientDirection  string `db:"cliente_direccion"`
	ClienteEmail     string `db:"cliente_email"`
	ClientePassword  string `db:"cliente_password"`
	ClienteTelephone string `db:"cliente_telefono"`
}

func (user *User) ToUserResponse() UserResponse {
	return UserResponse{
		ClientID:     user.ClientID,
		ClienteEmail: user.ClienteEmail,
		ClientName:   user.ClientName,
	}
}

type UserResponse struct {
	ClientID     int    `json:"cliente_id"`
	ClienteEmail string `json:"cliente_email"`
	ClientName   string `db:"cliente_nombre"`
}