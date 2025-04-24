package controllers

import (
	"net/http"

	"github.com/alerdn/rest-go/internal/services"
	"github.com/gin-gonic/gin"
)

type SaleController struct {
	service services.SaleService
}

type CreateSaleRequest struct {
	ProductId     int    `json:"product_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required,oneof=credit pix"`
}

func NewSaleController(service services.SaleService) SaleController {
	return SaleController{
		service,
	}
}

func (sc *SaleController) CreateSale(c *gin.Context) {
	var req CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, err := sc.service.Create(services.CreateRequest{
		ProductId:    req.ProductId,
		Quantity:     req.Quantity,
		PaymentMethod: req.PaymentMethod,
		UserID:            c.GetInt("userId"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": sale})
}

func (sc *SaleController) GetSales(c *gin.Context) {
	userId := c.GetInt("userId")

	sales, err := sc.service.GetSalesByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}
