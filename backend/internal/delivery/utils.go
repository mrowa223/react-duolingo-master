package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type envelope map[string]any

func readJSON(r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	err := writeJSON(w, status, message, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("[HTTP] Internal server error:\n%s", err)

	message := envelope{
		"message": "the server encountered a problem and could not process your request",
		"error":   err.Error(),
	}
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	message := envelope{
		"message": "validation failed",
		"errors":  errors,
	}

	errorResponse(w, r, http.StatusUnprocessableEntity, message)
}

func editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	errorResponse(w, r, http.StatusConflict, message)
}

func recordDuplicationResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to save the record due to a duplication error"
	errorResponse(w, r, http.StatusConflict, message)
}

func rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	errorResponse(w, r, http.StatusTooManyRequests, message)
}

func invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	errorResponse(w, r, http.StatusUnauthorized, message)
}
