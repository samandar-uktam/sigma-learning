package oop

type User struct {
	name  string // unexported (private)
	Email string // exported (public)
}

func NewUser(name, email string) *User {
	return &User{name: name, Email: email}
}

// Method to access private field
func (u *User) GetName() string {
	return u.name
}

// Method to modify private field
func (u *User) SetName(name string) {
	u.name = name
}

// Here User's name variable is encapsulated through getters and setters. You contron when to modify or access it.
