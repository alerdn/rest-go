package user

import (
	"github.com/alerdn/rest-go/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	UserController := CreateUserController()

	g := api.Group("/users")
	{
		g.POST("/login", UserController.Login)
		g.POST("/", UserController.Register)

		protected := g.Use(auth.Middleware())
		{
			protected.GET("/", UserController.Perfil)
		}
	}
}
