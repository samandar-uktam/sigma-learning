package solid

/*
	// db.go - Low-level module
	type ProductDB struct{}

	func (db ProductDB) GetProductByID(id int) string {
		// Imagine this hits a real database
		return "Product from DB"
	}
*/

/*
	// usecase.go - High-level module
	type ProductUseCase struct {
		db ProductDB // directly depends on a low-level detail!
	}

	func (uc *ProductUseCase) GetProduct(id int) string {
		return uc.db.GetProductByID(id)
	}
*/

// Here in usecase it is violated that the usecase is directly dependant on the db layer's object.
// To follow DIP principle, it should be connected with only interfaces.
