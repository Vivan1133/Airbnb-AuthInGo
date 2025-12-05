package utils

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {	
	return validator.New(validator.WithRequiredStructEnabled())
}


func WriteJsonSucessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{
		"message": message,
		"success": true,
		"data": data,
		"err": nil,
	}

	return WriteJsonResponse(w, status, response)
}

func WriteErrorResponse(w http.ResponseWriter, status int, message string, err error) error { 
	response := map[string]any {
		"message": message,
		"error": err,
		"data": nil,
		"success": false,
	}

	return WriteJsonResponse(w, status, response)
}

func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func ReadJsonRequest(r *http.Request, result any) error {

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	return decoder.Decode(result)

}
