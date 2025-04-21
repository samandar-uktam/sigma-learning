package solid

// This is a good example because the usecase is only connected by interfaces.

/*

	// repository.go
	type ProductRepository interface {
		GetProductByID(id int) string
	}

	// db.go
	type ProductDB struct{}

	func (db ProductDB) GetProductByID(id int) string {
		return "Product from DB"
	}

	// usecase.go
	type ProductUseCase struct {
		repo ProductRepository
	}

	func NewProductUseCase(r ProductRepository) *ProductUseCase {
		return &ProductUseCase{repo: r}
	}

	func (uc *ProductUseCase) GetProduct(id int) string {
		return uc.repo.GetProductByID(id)
	}
*/
