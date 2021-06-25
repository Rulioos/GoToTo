package main

//@tsInterface[context="internal"]
type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Given string `json:"given"`
}

//@tsInterface[context="external"]
type Person struct {
	Name  string `json:"name"`
	Given string `json:"given"`
}

//@tsInterface
type Product struct {
	Price int
	Name  string
}
