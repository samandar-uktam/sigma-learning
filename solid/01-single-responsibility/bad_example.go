package solid

// This is a bad example because User entity is responsible for both managing user data and interaction with database.

/*
	type User struct {
		FirstName string
		LastName  string
	}

	func (u User) GetFullName(firstName, lastName string) string {
		return u.FirstName + " " + u.LastName
	}

	func (u User) SaveIntoDb() {
		// saved user details into database
	}
*/
