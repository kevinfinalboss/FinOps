package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/kevinfinalboss/FinOps/api/controller"
	"github.com/kevinfinalboss/FinOps/internal/repository"
)

func RegisterRoutes(router *gin.Engine, userRepo *repository.UserRepository) {
	userController := api.NewUserController(userRepo)

	api := router.Group("/api/v1")
	{
		api.GET("/healthcheck", healthCheck)
		api.POST("/user/login", userController.LoginUser)
		api.POST("/user/register", userController.RegisterUser)
	}

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login Page",
		})
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
