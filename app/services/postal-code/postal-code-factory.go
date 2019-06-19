package postalcode

import models "github.com/miliyash/ms-contact-manager/app/models"

type IPostalCode interface {
	GetPostalCodeData(postalCode string) (result *models.PostalCode, err error)
}

func GetNewPostalCodeService() IPostalCode {
	return &INPostalCode{}
}
