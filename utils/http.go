package utils

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/elias-gill/poliapi/types"
)

// small function to boilerplate json response
func WriteJsonResponse(w http.ResponseWriter, status int, msg interface{}) {
    encoder := json.NewEncoder(w)
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(status)
    if msg != nil {
        encoder.Encode(msg)
    }
}

// To generate a new http error and a new log message
func GenerateHttpError(w http.ResponseWriter, status int, err error, msg string) types.HttpError {
    log.Println(err.Error())
    WriteJsonResponse(w, status, msg)
    return types.HttpError{Error: msg}
}

