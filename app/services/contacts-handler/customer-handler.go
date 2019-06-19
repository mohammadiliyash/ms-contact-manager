package contactshandler

import (
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"
)

// GetNewCustomerHandler ...
func GetNewCustomerHandler() ICustomer {
	return &customer{}
}

//customer ...
type customer dbModels.Customer

// ICustomer ...
type ICustomer interface {
}
