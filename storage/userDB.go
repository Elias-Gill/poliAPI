package storage

import (
	//We are using the pgx driver to connect to PostgreSQL

	"github.com/elias-gill/poliapi/types"
	"github.com/elias-gill/poliapi/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type UsersHandler struct {
	db UserStorer
}

// Compares credentials and then returns a new JWT token for the user
func (ud *UsersHandler) LoginUser(user string, pasw string) (*string, error) {
	// fetch for user data
	data, err := ud.db.GetById(user)
	if err != nil {
		return nil, err
	}

	// compare pasw and encripted pasw
	if err = utils.ComparePasw(pasw, *data.Pasw); err != nil {
		return nil, err
	}

	// returns a new JWT token
	return utils.GenerateJWT(string(user))
}

// TODO: pensar que hacer con esto
func (ud *UsersHandler) UpdateUserData(name string, new types.User) error {
	if new.Name != nil {
		// TODO: generar update con la funcion "UpdateSection"
	}
	if new.Email != nil {
		// TODO: generar update con la funcion "UpdateSection"
	}
	if new.Pasw != nil {
		_, err := utils.EncriptPasw(*new.Pasw)
		if err != nil {
			return err
		}
		// TODO: generar update con la funcion "UpdateSection"
	}
	return nil
}

// Inserts a new user into the database. Returns an error if the username is unavailable
func (ud *UsersHandler) RegisterNewUser(u types.User) error {
	err := u.ValidRegistration()
	if err != nil {
		return err
	}

	// encriptar la contrasena
	hashedPwd, err := utils.EncriptPasw(*u.Pasw)
	if err != nil {
		return err
	}
    u.Pasw = &hashedPwd

	// insertar el nuevo usuario
	err = ud.db.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user from the database, including all his schedules
func (ud *UsersHandler) DeleteUser(user string) error {
	return ud.db.Delete(user)
}
