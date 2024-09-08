package main

// User struct to store user details
type User struct {
	ID    string
	Name  string
	Email string
	Phone string
}

//Constructor
func NewUser(id, name, email, phone string) User{
	return User{
		ID: id,
		Name: name,
		Email: email,
		Phone: phone,
	}
}