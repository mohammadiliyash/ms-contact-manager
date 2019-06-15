package controllers

import (
	"net/http"
)

// SayError returns a 500
func SayError(w http.ResponseWriter, r *http.Request) {
	//e.ReturnError(w, e.GenericError, "something went wrong")
	w.WriteHeader(http.StatusInternalServerError)
}

// SayOK returns a 200
func SayOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
