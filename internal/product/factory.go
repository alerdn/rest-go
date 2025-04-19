package product

import (
	"github.com/alerdn/rest-go/config"
)

func CreateProductController() ProductController {
	ProductRepository := NewProductRepository(config.DB)
	ProductUsecase := NewProductUsecase(ProductRepository)
	ProductController := NewProductController(ProductUsecase)

	return ProductController
}
