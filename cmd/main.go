package main

import (
	"arsip/config"
	"arsip/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.InitRedis()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
