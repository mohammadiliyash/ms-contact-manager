package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/miliyash/ms-contact-manager/app/models"
	"github.com/stretchr/testify/mock"
)

func TestHandleLogin(t *testing.T) {

	//cfg.Initialize()

	postalCodeservice := &PostalCodeMock{}
	city := "delhi"
	postalCodeData := models.PostalCode{
		City: &city,
	}
	req, err := http.NewRequest("Get", "/login/00601", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	postalCodeservice.On("GetPostalCodeData", "").Return(&postalCodeData, nil)
	controller := GetMainController(postalCodeservice, nil, nil)
	handler := http.HandlerFunc(controller.HandlePostalCodeLookup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"city":"delhi","pid":"","state":null,"value":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

// AuthServiceMock => Mock for AuthService
type PostalCodeMock struct {
	mock.Mock
}

// GetPostalCodeData => Mocked GetPostalCodeData
func (pc *PostalCodeMock) GetPostalCodeData(postalCode string) (result *models.PostalCode, err error) {
	args := pc.Called(postalCode)
	return args.Get(0).(*models.PostalCode), args.Error(1)
}
