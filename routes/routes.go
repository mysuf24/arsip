package routes

import (
	"arsip/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/arsip", controllers.GetArsip)
	}
}
