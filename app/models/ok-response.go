package models

import "github.com/miliyash/ms-contact-manager/app/utils"

// OkRespose ..
type OkRespose struct {
	utils.ServiceResult
	Data string `json:",omitempty"`
}
