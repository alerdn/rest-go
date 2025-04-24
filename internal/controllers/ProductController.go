package controllers

import (
	"net/http"
	"strconv"

	"github.com/alerdn/rest-go/internal/models"
	"github.com/alerdn/rest-go/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return ProductController{
		service,
	}
}

func (p *ProductController) GetProduct(c *gin.Context) {
	products, err := p.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	newProduct, err := p.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
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

	product, err := p.service.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produto não foi encontrado."})
		return
	}

	c.JSON(http.StatusOK, product)
}
