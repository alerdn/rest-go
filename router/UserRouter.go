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

		g.POST("/forgot-password", UserController.RequestPassRecovery)

		protected := g.Use(middlewares.Middleware())
		{
			protected.GET("/", UserController.Profile)
			protected.POST("/avatar", UserController.UploadAvatar)
			protected.GET("/avatar", UserController.DownloadAvatar)
			protected.DELETE("/avatar", UserController.DeleteAvatar)
		}
	}
}
