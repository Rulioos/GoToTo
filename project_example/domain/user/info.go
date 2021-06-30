package userDomain

//@tsInterface[context="newContext"]
type Address struct {
	Id          uint   `json:"id"`
	Street      string `json:"street"`
	HouseNumber string `json:"h_number"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
}
