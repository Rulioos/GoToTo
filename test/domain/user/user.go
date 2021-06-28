package userDomain

//@tsInterface[context="user"]
type User struct {
	Person   Person `json:"person"`
	Password string `json:"pwd"`
	Login    string `json:"login"`
}

//@tsInterface[context="user"]
type Person struct {
	Name   string `json:"name"`
	Given  string `json:"given"`
	Gender string `json:"gender"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email"`
}
