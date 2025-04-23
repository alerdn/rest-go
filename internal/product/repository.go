package product

import (
	"log"

	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProductRepository {
	return ProductRepository{
		connection,
	}
}

func (pr *ProductRepository) CreateProduct(product Product) (int, error) {
	result := pr.connection.Create(&product)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}

	return int(product.ID), nil
}

func (pr *ProductRepository) GetProducts() ([]Product, error) {
	var products []Product

	result := pr.connection.Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetProductById(id int) (*Product, error) {
	var product Product

	result := pr.connection.Find(&product, id)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		log.Println("Product not found")
		return nil, nil
	}

	return &product, nil
}
