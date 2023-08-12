package server

import (
	"fmt"

	"github.com/elias-gill/poliapi/src/utils"
)

// Compares credentials and then returns a new JWT token for the user
func LoginUser(user string, pasw string) (*string, error) {
	var data User
	err := db.Select(
		&data,
		"SELECT u.password, u.name FROM usuarios u (WHERE user = $1)",
		user)
	if err != nil {
		return nil, err
	}

	// compare pasw and encripted pasw
	if err = utils.ComparePasw(pasw, *data.Pasw); err != nil {
		return nil, err
	}

	// returns a new JWT token
	return utils.GenerateJWT(user)
}

// TODO: pensar que hacer con esto
func UpdateUserData(name string, new User) error {
	if new.Name != nil {
		_, err := db.Query("update usuarios u set nombre = $1 where u.nombre = $2", new.Name, name)
		if err != nil {
			return err
		}
	}
	if new.Email != nil {
		_, err := db.Query("update usuarios u set email = $1 where u.nombre = $2", new.Email, name)
		if err != nil {
			return err
		}
	}
	if new.Pasw != nil {
		hash, err := utils.EncriptPasw(*new.Pasw)
		if err != nil {
			return err
		}
		_, err = db.Query("update usuarios u set password = $1 where u.nombre = $2", new.Name, hash)
		if err != nil {
			return err
		}
	}
	return nil
}

// validar la validez de los campos para el registro
func validRegistration(u User) error {
	if u.Email == nil || u.Pasw == nil || u.Name == nil {
		return fmt.Errorf("Los parametros no pueden ser nulos")
	}
	return nil
}

// Inserts a new user into the database. Returns an error if the username is unavailable
func RegisterNewUser(u User) error {
	err := validRegistration(u)
	if err != nil {
		return err
	}

	// comparobar que el nombre no este repetido
	if err = nameIsNotRepeated(*u.Name); err != nil {
		return err
	}

	// encriptar la contrasena
	hashedPwd, err := utils.EncriptPasw(*u.Pasw)
	if err != nil {
		return err
	}

	// insertar el nuevo usuario
	err = inserUserInDb(u, hashedPwd)
	if err != nil {
		return err
	}
	return nil
}

func inserUserInDb(u User, hashedPwd string) error {
	// insertar
	_, err := db.Query(
		"insert into users (nombre, password, email) values ($1, $2, $3)", u.Name, hashedPwd, u.Email)
	if err != nil {
		return err
	}
	return nil
}

func nameIsNotRepeated(name string) error {
	// buscar en la db
	rows, err := db.Query(
		"select count(*) from usuarios where nombre = $1", name)
	if err != nil {
		return err
	}
	// guardar la cantidad de matches en un contador
	var res *int
	rows.Scan(res)
	// si el contador es mayor a 0 entonces es repetido
	if *res != 0 {
		return fmt.Errorf("Nombre de usuario no disponible")
	}
	return nil
}

// Delete a user from the database, including all his schedules
func DeleteUser(user string) error {
	_, err := db.Query("delete from usuarios where nombre = $1 cascade", user)
	if err != nil {
		return err
	}
	return nil
}
