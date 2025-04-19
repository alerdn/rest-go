package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase UserUsecase
}

func NewUserController(usecase UserUsecase) UserController {
	return UserController{
		usecase,
	}
}

func (u *UserController) Index(c *gin.Context) {
	users, err := u.usecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
