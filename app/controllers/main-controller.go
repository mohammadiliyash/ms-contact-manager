package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	//Framework

	"github.com/gorilla/mux"
	models "github.com/miliyash/ms-contact-manager/app/models"
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"

	ch "github.com/miliyash/ms-contact-manager/app/services/contacts-handler"
	pc "github.com/miliyash/ms-contact-manager/app/services/postal-code"
	"github.com/miliyash/ms-contact-manager/app/utils"
)

var chContact = make(chan *dbModels.Contact, 1000)

// GetMainController ...
func GetMainController(postalcode pc.IPostalCode, contactHandler ch.IContact, customerHandler ch.ICustomer) *MainController {
	return &MainController{Postalcode: postalcode, ContactHandler: contactHandler, CustomerHandler: customerHandler}
}

// MainController struct
type MainController struct {
	Postalcode      pc.IPostalCode
	ContactHandler  ch.IContact
	CustomerHandler ch.ICustomer
}

// HandlePostalCodeLookup ...
func (m *MainController) HandlePostalCodeLookup(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	var jsonAPIError *utils.JSONAPIError
	var result *models.PostalCode
	args := mux.Vars(r)
	postalCode := args["code"]

	result, err := m.Postalcode.GetPostalCodeData(postalCode)

	if err != nil {
		jsonAPIError = &utils.JSONAPIError{Status: "400", Detail: err.Error()}
	}
	HandlePostResponse(c, w, jsonAPIError == nil, jsonAPIError, result)
}

// HandleLegacyPostalCodeLookup returns rating data for a postal code, wrapped in a legacy struct
func (m *MainController) HandleLegacyPostalCodeLookup(w http.ResponseWriter, r *http.Request) {

	c := r.Context()
	var jsonAPIError *utils.JSONAPIError
	var resultPc *models.PostalCode
	args := mux.Vars(r)
	postalCode := args["code"]
	var result models.LegacyContainer

	resultPc, err := m.Postalcode.GetPostalCodeData(postalCode)

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

//HandleCreateCustomer ...
func (m *MainController) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {

	c := r.Context()
	var jsonAPIError *utils.JSONAPIError
	result := &models.OkRespose{}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var resource dbModels.Contact

	err := json.Unmarshal(body, &resource)

	if err == nil {

		err = resource.Validate()
		if err == nil {

			chContact <- &resource
			m.ContactHandler.CreateContactFromchannel(chContact)
		}

	}
	if err != nil {
		jsonAPIError = &utils.JSONAPIError{Status: "400", Detail: err.Error()}
	}
	HandlePostResponse(c, w, jsonAPIError == nil, jsonAPIError, result)

}
