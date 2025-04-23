package sale

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaleController struct {
	usecase SaleUsecase
}

func NewSaleController(usecase SaleUsecase) SaleController {
	return SaleController{
		usecase,
	}
}

func (sc *SaleController) GetSales(c *gin.Context) {
	userId := c.GetInt("userId")


	sales, err := sc.usecase.GetSalesByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}
