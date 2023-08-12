package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elias-gill/poliapi/src/server"
	"github.com/go-chi/chi/v5"
)

type JWTResponse struct {
	Token string `json:"token"`
}

type httpError struct {
	Error string `json:"error"`
}

// Returns a new handler for the /users path
func UsersHandler(r chi.Router) {
	r.Get("/login", login)
	r.Delete("/", deleteUser)
	r.Put("/", registerUser)
	r.Post("/", updateUser)
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
func login(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-type", "application/json")
	user, pasw, ok := r.BasicAuth() // get user credentials
	if !ok {
		// error de autenticacion
		log.Println("Error al parsear credenciales " + user + " " + pasw)
		w.WriteHeader(400)
		encoder.Encode(httpError{Error: "Formato de credenciales invalido"})
		return
	}
	token, err := server.LoginUser(user, pasw)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(401)
		encoder.Encode(httpError{Error: "Usuario o contrasena invalidos"})
		return
	}
	// mandar el jwt
	w.WriteHeader(200)
	encoder.Encode(JWTResponse{Token: *token})
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
func registerUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-type", "application/json")
	var body server.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Print("Error de formateo: ", err.Error())
		w.WriteHeader(400)
		encoder.Encode(httpError{Error: "No es posible parsear la request, formato invalido"})
		return
	}
	// crear el nuevo usario en la db
	if err := server.RegisterNewUser(body); err != nil {
		log.Println("Error al crear nuevo usuario: ", err.Error())
		w.WriteHeader(400)
		encoder.Encode(httpError{Error: err.Error()})
		return
	}
	w.WriteHeader(200)
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
func updateUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-type", "application/json")
	userName, _, _ := r.BasicAuth()
	// extraer los datos de la request
	var body server.User
	json.NewDecoder(r.Body).Decode(&body)
	err := server.UpdateUserData(userName, body)
	if err != nil {
		log.Println("No se pudo modificar usuario" + err.Error())
		w.WriteHeader(400)
		encoder.Encode(httpError{Error: "No se puede realizar la operacion"})
		return
	}
	w.WriteHeader(200)
}

// @Summary		Eliminar un usuario
// @Description Eliminar todos los registros de un usuario de la base de datos
// @Tags		users
// @Accept		json
// @Produce		json
// @Failure		403 {string} string "Parametros invalidos"
// @Succes		200
// @Router		/ [delete]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-type", "application/json")
	user, _, _ := r.BasicAuth()
	err := server.DeleteUser(user)
	if err != nil {
		log.Println("Delete error: ", err.Error())
		w.WriteHeader(403)
		encoder.Encode(httpError{Error: "No se pudo realizar la operacion"})
		return
	}
	w.WriteHeader(200)
}
