package router

import (
	"github.com/alerdn/rest-go/internal/factories"
	"github.com/alerdn/rest-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup) {
	UserController := factories.CreateUserController()

	g := api.Group("/users")
	{
		g.POST("/login", UserController.Login)
		g.POST("/", UserController.Register)

		protected := g.Use(middlewares.Middleware())
		{
			protected.GET("/", UserController.Perfil)
		}
	}
}
