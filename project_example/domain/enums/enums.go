package enums

//go:generate stringer -type=paymentMethods,roles -output=stringEnumsMethods.go

type paymentMethods int
type roles int

//@tsEnum
const (
	CB paymentMethods = iota
	CASH
	GIFT_CARD

	OWNER roles = iota
	GENERAL_MANAGER
	CHEF_COOK
	SERVER
)
