package models

import "github.com/miliyash/ms-contact-manager/app/utils"

// LegacyContainer ...
type LegacyContainer struct {
	utils.ServiceResult
	Data *LegacyData `json:"data"`
	ID   *string     `json:"cid"`
	Type *string     `json:"type"`
}

// LegacyData ...
type LegacyData struct {
	Attributes *PostalCode `json:"attributes"`
}
