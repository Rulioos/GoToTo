package main

//@tsInterface[context="internal"]
type User struct {
	Id    uint
	Name  string
	Given string
}

//@tsInterface[context="external"]
type Person struct {
	Name  string
	Given string
}

//@tsInterface
type Product struct {
	Price int
	Name  string
}
