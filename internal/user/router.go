package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup) {
	UserController := CreateUserController()

	g := api.Group("/users")
	{
		g.GET("/", UserController.Index)
	}
}
