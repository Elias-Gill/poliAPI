package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SheetsHandler(r chi.Router) {
	r.Get("/sheets", getExcelFilesList)
	r.Get("/sheets/subjects", getSubjectsFromExcel)
}

// @Summary		Listar las carreras y archivos excel disponibles
// @Description Retorna una lista de los archivos excel disponibles en el servidor, asi como de las carreras
// @Tags		sheets
// @Accept		json
// @Produce		json
// @Failure		500 query string "Error en el servidor"
// @Succes		200 query {object} TODO: hacer el objecto
// @Router		/sheets [get]
func getExcelFilesList(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Listar las materias por carrera contenidas en un excel
// @Description Retorna la lista de materias disponibles por carrera en el archivo excel
// @Tags		sheets
// @Accept		json
// @Produce		json
// @Failure		500 query string "Error en el servidor"
// @Succes		200 query {object} TODO: hacer el objecto
// @Router		/sheets/subjects [get]
func getSubjectsFromExcel(w http.ResponseWriter, r *http.Request) {
	return
}
