package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type httpError struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}

type successMessage struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	e := httpError{
		Status:    http.StatusInternalServerError,
		Error:     "Internal Server Error",
		Message:   "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		Path:      r.RequestURI,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}

func notFound(message string, w http.ResponseWriter, r *http.Request) {
	e := httpError{
		Status:    http.StatusNotFound,
		Error:     "Not Found",
		Message:   message,
		Path:      r.RequestURI,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	e := httpError{
		Status:    http.StatusUnauthorized,
		Error:     "Unauthorized",
		Message:   "The request has not been completed because it lacks valid authentication credentials for the requested resource.",
		Path:      r.RequestURI,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("WWW-Authenticate", "Custom realm=\"Login via /api/login\"")
	w.WriteHeader(http.StatusUnauthorized)
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	e := httpError{
		Status:    http.StatusForbidden,
		Error:     "Unauthorized",
		Message:   "The access is permanently forbidden and tied to the application logic, such as insufficient rights to a resource.",
		Path:      r.RequestURI,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}

func ok(obj interface{}, w http.ResponseWriter, _ *http.Request) {
	payload, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}

func badRequest(message string, w http.ResponseWriter, r *http.Request) {
	e := httpError{
		Status:    http.StatusBadRequest,
		Error:     "Bad Request",
		Message:   message,
		Path:      r.RequestURI,
		Timestamp: time.Now(),
	}

	payload, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
	}
}
