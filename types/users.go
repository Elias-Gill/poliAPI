package types

import "fmt"

type NewUserRequest struct {
	Name  *string `json:"userName"`
	Pasw  *string `json:"password"`
	Email *string `json:"email"`
}

type UpdateUserRequest struct {
	Name  *string `json:"userName" db:"nombre"`
	Pasw  *string `json:"password" db:"password"`
	Email *string `json:"email" db:"email"`
}

type UserGetResponse struct {
    Name  *string `json:"userName"`
    Email *string `json:"email"`
}

type User struct {
	Name  *string `json:"userName" db:"nombre"`
	Pasw  *string `json:"password" db:"password"`
	Email *string `json:"email" db:"email"`
}

// Validates all inputs and then sends a new user
func NewUserFromRequest(req NewUserRequest) (*User, error) {
	if req.Email == nil || req.Pasw == nil || req.Name == nil {
		return nil, fmt.Errorf("Los parametros no pueden ser nulos")
	}

	return &User{
		Email: req.Email,
		Name:  req.Name,
		Pasw:  req.Pasw,
	}, nil
}
