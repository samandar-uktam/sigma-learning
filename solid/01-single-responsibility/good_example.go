package solid

// This is a good example, because each object has one responsibility where User struct is responsible for managing user data
// while UserRepository is responsible for saving and retrieving user data into db.

/*
	type User struct {
		FirstName string
		LastName  string
	}

	func (u User) GetFullName(firstName, lastName string) string {
		return u.FirstName + " " + u.LastName
	}

	type UserRepository struct{

	}

	func (u UserRepository) SaveIntoDb() {
		// saved user details into database
	}
*/
