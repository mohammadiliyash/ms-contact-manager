package models

import "github.com/miliyash/ms-contact-manager/app/utils"

//ContactModel => Model holds user's contact details
type ContactModel struct {
	utils.ServiceResult
	// UserID         string
	ID             int
	ContactName    string
	ContactPhone   string
	ContactAddress string
}
