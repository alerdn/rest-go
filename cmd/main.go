package main

import (
	"github.com/alerdn/rest-go/config"
	"github.com/alerdn/rest-go/router"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	server := router.RegisterRoutes()
	server.Run(":3000")
}
