package utility

import (
	"encoding/json"
	"net/http"
)

type BaseError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func CreateBaseError(err error,statusCode int,message string) BaseError  {
	var _error BaseError
	_error.Error = err.Error()
	_error.Message = message
	_error.StatusCode = statusCode
	return _error
}

func GenerateError(w http.ResponseWriter,err error,statusCode int,message string)  {
	var _err = CreateBaseError(err, statusCode, message)
	http.Error(w, "", statusCode)
	json.NewEncoder(w).Encode(_err)
}