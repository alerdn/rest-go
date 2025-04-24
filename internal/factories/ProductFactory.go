package factories

import (
	"github.com/alerdn/rest-go/config"
	"github.com/alerdn/rest-go/internal/controllers"
	"github.com/alerdn/rest-go/internal/repositories"
	"github.com/alerdn/rest-go/internal/services"
)

func CreateProductController() controllers.ProductController {
	ProductRepository := repositories.NewProductRepository(config.DB)
	ProductUsecase := services.NewProductService(ProductRepository)
	ProductController := controllers.NewProductController(ProductUsecase)

	return ProductController
}
