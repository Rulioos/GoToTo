package main

//@tsInterface[context="internal"]
type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Given string `json:"given,omitempty"`
}

//@tsInterface[context="external"]
type Adress struct {
	Name string `json:"given"`
}

//@tsInterface[context="external"]
type Person struct {
	Name     string   `json:"name"`
	Given    string   `json:"given"`
	Adresses []Adress `json:"adresses,omitempty"`
}

//@tsInterface
type Product struct {
	Price int    `json:"price"`
	Name  string `json:"name"`
}
