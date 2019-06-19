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

	//db := database.NewDataBaseService()
	//result, err := db.GetDB()
	//result, err = result.CreateDatabase()
	//fmt.Println(result, err)

	w.WriteHeader(http.StatusCreated)
}
