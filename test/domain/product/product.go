package productDomain

//@tsInterface[context="Product"]
type Item struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	InStock bool   `json:"in_stock"`
}
