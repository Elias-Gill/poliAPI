package types

import "fmt"

type ExcelInfo struct {
	Id     int    `json:"id"`
	Fecha  string `json:"date"`
	Nombre string `json:"fileName"`
}

type JWTResponse struct {
    Token string `json:"token"`
}

type HttpError struct {
    Error string `json:"error"`
}

type User struct {
	Name  *string `json:"userName" db:"nombre"`
	Pasw  *string `json:"password" db:"password"`
	Email *string `json:"email" db:"email"`
}

// valida que todos los campos sean correctos
func (u *User) ValidateParameters() error {
    if u.Email == nil || u.Pasw == nil || u.Name == nil {
        return fmt.Errorf("Los parametros no pueden ser nulos")
    }
    return nil
}
