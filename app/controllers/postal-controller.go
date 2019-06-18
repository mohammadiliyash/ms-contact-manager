package controllers

import (
	"net/http"

	//Framework
	"github.com/gorilla/mux"
	models "github.com/miliyash/ms-contact-manager/app/models"
	pcs "github.com/miliyash/ms-contact-manager/app/services/postal-code"
	"github.com/miliyash/ms-contact-manager/app/utils"
)

//HandlePostalCodeLookup
func HandlePostalCodeLookup(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var jsonAPIError *utils.JSONAPIError
	var result *models.PostalCode
	args := mux.Vars(r)
	postalCode := args["code"]

	result, err := pcs.GetPostalCodeData(postalCode)

	if err != nil {
		jsonAPIError = &utils.JSONAPIError{Status: "400", Detail: err.Error()}
	}
	HandlePostResponse(c, w, jsonAPIError == nil, jsonAPIError, result)
}

// HandleLegacyPostalCodeLookup returns rating data for a postal code, wrapped in a legacy struct
func HandleLegacyPostalCodeLookup(w http.ResponseWriter, r *http.Request) {

	c := r.Context()
	var jsonAPIError *utils.JSONAPIError
	var resultPc *models.PostalCode
	args := mux.Vars(r)
	postalCode := args["code"]
	var result models.LegacyContainer
	resultPc, err := pcs.GetPostalCodeData(postalCode)

	if err == nil {

		dataType := "postalCodes"
		legacyData := models.LegacyData{Attributes: resultPc}
		result = models.LegacyContainer{Data: &legacyData, ID: &postalCode, Type: &dataType}
		//result, _ = json.MarshalIndent(legacyDataContainer, "", "  ")
	}
	if err != nil {
		jsonAPIError = &utils.JSONAPIError{Status: "400", Detail: err.Error()}
	}
	HandlePostResponse(c, w, jsonAPIError == nil, jsonAPIError, result)

}
