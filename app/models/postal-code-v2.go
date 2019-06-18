package models

import "github.com/miliyash/ms-contact-manager/app/utils"

type LegacyContainer struct {
	utils.ServiceResult
	Data *LegacyData `json:"data"`
	ID   *string     `json:"id"`
	Type *string     `json:"type"`
}

type LegacyData struct {
	Attributes *PostalCode `json:"attributes"`
}
