package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	usecase ProductUsecase
}

func NewProductController(usecase ProductUsecase) ProductController {
	return ProductController{
		usecase,
	}
}

func (p *ProductController) GetProduct(c *gin.Context) {
	products, err := p.usecase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	}

	newProduct, err := p.usecase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (p *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Informe o ID do produto."})
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID do produto deve ser um numero"})
		return
	}

	product, err := p.usecase.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produto não foi encontrado."})
		return
	}

	c.JSON(http.StatusOK, product)
}
