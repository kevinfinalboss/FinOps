package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/healthcheck", healthCheck)
		//api.POST("/login", login)
		//api.POST("/register", register)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
