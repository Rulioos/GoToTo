package productDomain

//go:generate stringer -type=productType -output=stringEnumsMethods.go

type productType int

const (
	TOY productType = iota
	FOOD
	CLOTH
)
