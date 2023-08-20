package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elias-gill/poliapi/types"
)

// small function to boilerplate json response
func writeJsonResponse(w http.ResponseWriter, status int, msg interface{}) {
    encoder := json.NewEncoder(w)
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(status)
    if msg != nil {
        encoder.Encode(msg)
    }
}

// To generate a new http error and a new log message
func generateHttpError(w http.ResponseWriter, status int, err error, msg string) types.HttpError {
    log.Println(err.Error())
    writeJsonResponse(w, status, msg)
    return types.HttpError{Error: msg}
}
