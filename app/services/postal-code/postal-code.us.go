package postalcode

import (
	"fmt"

	models "github.com/miliyash/ms-contact-manager/app/models"
)

type USPostalCode models.PostalCode

func (p *USPostalCode) GetPostalCodeData(postalCode string) (result *models.PostalCode, err error) {

	result = models.PostalCodeDict[postalCode]
	if result == nil {
		return nil, fmt.Errorf("No Data found for postal code %s", postalCode)
	}
	result.ID = postalCode

	return
}
