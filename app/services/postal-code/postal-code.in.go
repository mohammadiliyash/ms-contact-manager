package postalcode

import (
	"fmt"

	models "github.com/miliyash/ms-contact-manager/app/models"
)

type INPostalCode models.PostalCode

func (p *INPostalCode) GetPostalCodeData(postalCode string) (result *models.PostalCode, err error) {

	var city = "Mock City Delhi"

	result = &models.PostalCode{
		ID:   "001",
		City: &city,
	}
	if result == nil {
		return nil, fmt.Errorf("No Data found for postal code %s", postalCode)
	}
	result.ID = postalCode

	return
}
