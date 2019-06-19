package postalcode

import models "github.com/miliyash/ms-contact-manager/app/models"

// IPostalCode ...
type IPostalCode interface {
	GetPostalCodeData(postalCode string) (result *models.PostalCode, err error)
}

//GetNewPostalCodeService ...
func GetNewPostalCodeService() IPostalCode {
	return &USPostalCode{}
}
