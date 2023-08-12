package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TODO: cambiar todos los mensajes de swagger
// Returns a new handler for the /schedules path
func ScheduleHandler(r chi.Router) {
	r.Get("/", getUserSchedules)
	r.Post("/", newSchedule)
	r.Put("/", modifySchedule)
	r.Delete("/", deleteSchedule)
}

// @Summary		Iniciar Sesion
// @Description Iniciar sesion en una cuenta de usuario y generar un token
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		name query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		paswd query	string	true "La contrasena" minlength(5)  maxlength(30)
// @Failure		400 query string "Parametros invalidos"
// @Succes		200 query Token "JWT para autenticacion"
// @Router		/login [get]
func getUserSchedules(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Iniciar Sesion
// @Description Iniciar sesion en una cuenta de usuario y generar un token
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		name query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		paswd query	string	true "La contrasena" minlength(5)  maxlength(30)
// @Failure		400 query string "Parametros invalidos"
// @Succes		200 query Token "JWT para autenticacion"
// @Router		/login [get]
func newSchedule(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Iniciar Sesion
// @Description Iniciar sesion en una cuenta de usuario y generar un token
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		name query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		paswd query	string	true "La contrasena" minlength(5)  maxlength(30)
// @Failure		400 query string "Parametros invalidos"
// @Succes		200 query Token "JWT para autenticacion"
// @Router		/login [get]
func deleteSchedule(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Iniciar Sesion
// @Description Iniciar sesion en una cuenta de usuario y generar un token
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		name query	string	true "Nombre del usuario" minlength(5)  maxlength(30)
// @Param		paswd query	string	true "La contrasena" minlength(5)  maxlength(30)
// @Failure		400 query string "Parametros invalidos"
// @Succes		200 query Token "JWT para autenticacion"
// @Router		/login [get]
func modifySchedule(w http.ResponseWriter, r *http.Request) {
	return
}
