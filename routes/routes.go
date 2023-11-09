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
		//api.POST("/login", userController.LoginUser) er
		api.POST("/user/register", userController.RegisterUser)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
