package services

import (
	"github.com/alerdn/rest-go/internal/models"
	"github.com/alerdn/rest-go/internal/repositories"
)

type ProductService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return ProductService{
		repository,
	}
}

func (pu *ProductService) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = id

	return product, nil
}

func (pu *ProductService) GetProductById(id int) (*models.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
