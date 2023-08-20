package api

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/poliapi/storage"
	"github.com/elias-gill/poliapi/types"
	"github.com/elias-gill/poliapi/utils"
	"github.com/go-chi/chi/v5"
)

const defaultHttpError = "No se pudo realizar la operacion"

type UsersHandler struct {
	storer storage.UserStorer
}

func NewUsersHandler(store storage.UserStorer) UsersHandler {
	return UsersHandler{
		storer: store,
	}
}

// Returns a new handleUsers for the /users path
func (u UsersHandler) HandleUsers(r chi.Router) {
	r.Get("/login", u.LoginUser)
	r.Delete("/", u.DeleteUser)
	r.Put("/", u.RegisterUser)
	r.Post("/", u.UpdateUser)
}

// @Summary		Iniciar Sesion
// @Description Iniciar sesion en una cuenta de usuario y generar un token de sesion
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		userName query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		password query	string	true "La contrasena" minlength(5)  maxlength(30)
// @Failure		400 query string "Parametros invalidos"
// @Failure		403 query string "Usuario o contasena invalidos"
// @Succes		200 query Token "JWT para autenticacion"
// @Router		/ [get]
func (u UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, pasw, ok := r.BasicAuth()
	if !ok {
		// error de autenticacion
		logMsg := "Error al parsear credenciales " + user + " " + pasw
		msg := "Formato de credenciales invalido"
		writeJsonResponse(w, 400, generateHttpError(msg, logMsg))
		return
	}

	// generar un nuevo jwt para el usuario
	token, err := u.login(user, pasw)
	if err != nil {
		msg := "Usuario o contrasena invalidos"
		writeJsonResponse(w, 401, generateHttpError(msg, err.Error()))
		return
	}

	// mandar el jwt
	writeJsonResponse(w, 200, &types.JWTResponse{Token: *token})
	return
}

// @Summary		Crear usuario
// @Description Crear una nueva cuenta de usuario con los datos proporcionados
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		userName query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		password query	string	true "Contrasena del usuario" minlength(5)  maxlength(30)
// @Param		email query	string	true "User email" minlength(5)  maxlength(30)
// @Failure		400 {string} string "Parametros invalidos"
// @Succes		200
// @Router		/ [post]
func (u UsersHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body types.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logMsg := "Error de formateo: " + err.Error()
		msg := "No es posible parsear la request, formato invalido"
		writeJsonResponse(w, 400, generateHttpError(msg, logMsg))
		return
	}

	// crear el nuevo usario en la db
	if err := u.RegisterNewUser(body); err != nil {
		logMsg := "Error al crear nuevo usuario: " + err.Error()
		writeJsonResponse(w, 400, generateHttpError(err.Error(), logMsg))
		return
	}
	writeJsonResponse(w, 200, nil)
}

// @Summary		Modificar usuario
// @Description Permite modificar email, nombre de usuario o contrasena
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		userName query	string	false "El nuevo nombre de usuario" minlength(5)  maxlength(30)
// @Param		password query	string	false "La nueva contrasena para el usuario" minlength(5)  maxlength(30)
// @Param		email query	string	false "La nueva direccion email" minlength(5)  maxlength(30)
// @Failure		400 {string} string "Parametros invalidos"
// @Succes		200
// @Router		/ [put]
func (u UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// extraer los datos de la request
	userName, _, _ := r.BasicAuth()
	var body types.User
	json.NewDecoder(r.Body).Decode(&body)
	// hacer lo que se tenga que hacer
	err := u.updateUserData(userName, body)
	if err != nil {
		logMsg := "No se pudo modificar usuario" + err.Error()
		msg := "No se puede realizar la operacion"
		writeJsonResponse(w, 400, generateHttpError(msg, logMsg))
		return
	}
	writeJsonResponse(w, 200, nil)
}

// @Summary		Eliminar un usuario
// @Description Eliminar todos los registros de un usuario de la base de datos
// @Tags		users
// @Accept		json
// @Produce		json
// @Failure		403 {string} string "Parametros invalidos"
// @Succes		200
// @Router		/ [delete]
func (u UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, _, _ := r.BasicAuth()
	err := u.storer.Delete(user)
	if err != nil {
		writeJsonResponse(w, 403, generateHttpError(defaultHttpError, err.Error()))
		return
	}
	writeJsonResponse(w, 200, nil)
}

// Compares credentials and then returns a new JWT token for the user
func (ud UsersHandler) login(user string, pasw string) (*string, error) {
	// fetch for user data
	data, err := ud.storer.GetById(user)
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
func (ud *UsersHandler) updateUserData(name string, new types.User) error {
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
    err := u.ValidateParameters()
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
    err = ud.storer.Insert(u)
    if err != nil {
        return err
    }
    return nil
}
