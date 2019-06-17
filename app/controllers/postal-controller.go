package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	//Framework
	"github.com/gorilla/mux"
	models "github.com/miliyash/ms-contact-manager/app/models"
	log "github.com/miliyash/ms-contact-manager/logging"
)

//HandlePostalCodeLookup
func HandlePostalCodeLookup(w http.ResponseWriter, r *http.Request) {

	c := r.Context()
	args := mux.Vars(r)
	postalCode := args["code"]
	lw := log.WithField("postal-code", postalCode)
	url := r.URL
	if url == nil {
		lw.CError(c, "HandlePostalCodeLookup: Nil URL!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	lw = lw.WithField("request-url", url)
	pcRecord := models.PostalCodeDict[postalCode]
	if pcRecord == nil {
		lw.WithError(errors.New("Postal code not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		pcRecord.ID = postalCode
		lw.WithError(errors.New("Postal code not found"))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		foo, _ := json.MarshalIndent(pcRecord, "", "  ")

		w.Write(foo)
	}
}

// HandleLegacyPostalCodeLookup returns rating data for a postal code, wrapped in a legacy struct
func HandleLegacyPostalCodeLookup(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	args := mux.Vars(r)
	postalCode := args["code"]
	lw := log.WithField("postal-code", postalCode)
	url := r.URL
	if url == nil {
		lw.CError(c, "HandlePostalCodeLookup: Nil URL!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	lw = lw.WithField("request-url", url)
	pcRecord := models.PostalCodeDict[postalCode]
	if pcRecord == nil {
		lw.WithError(errors.New("Postal code not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		pcRecord.ID = postalCode
		lw.WithError(errors.New("Postal code not found"))
		dataType := "postalCodes"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		quux := models.LegacyData{Attributes: pcRecord}
		fum := models.LegacyContainer{Data: &quux, ID: &postalCode, Type: &dataType}
		foo, _ := json.MarshalIndent(fum, "", "  ")

		w.Write(foo)
	}
}
