package models

type LegacyContainer struct {
	Data *LegacyData `json:"data"`
	ID   *string     `json:"id"`
	Type *string     `json:"type"`
}

type LegacyData struct {
	Attributes *PostalCode `json:"attributes"`
}
