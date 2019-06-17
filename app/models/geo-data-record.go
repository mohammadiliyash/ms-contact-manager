package models

// GeoDataRecord holds imported geo data records
type GeoDataRecord struct {
	Zip            string
	FIPS           string `logging:"fips-code"`
	City           string
	CountyNameType string
	ZipType        string
	CityAbb        string
	LL             string
	State          string
	CountyName     string
}
