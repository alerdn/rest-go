package product

type ProductUsecase struct {
	repository ProductRepository
}

func NewProductUsecase(repository ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product Product) (Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}

	product.ID = id

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id int) (*Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
