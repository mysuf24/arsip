package controllers

import (
	"arsip/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArsip(c *gin.Context) {
	data, err := services.GetAllArsip()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, data)
}
