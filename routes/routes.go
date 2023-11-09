package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/kevinfinalboss/FinOps/api/controller"
	"github.com/kevinfinalboss/FinOps/api/middlewares"
	"github.com/kevinfinalboss/FinOps/internal/repository"
	"github.com/kevinfinalboss/FinOps/internal/services"
)

func RegisterRoutes(router *gin.Engine, userRepo *repository.UserRepository, spendingRepo *repository.SpendingRepository) {
	userController := api.NewUserController(userRepo)
	userService := services.NewUserService(userRepo)
	spendingService := services.NewSpendingService(spendingRepo)
	spendingController := api.NewSpendingController(spendingService, userService)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/healthcheck", healthCheck)
		apiGroup.POST("/user/login", userController.LoginUser)
		apiGroup.POST("/user/register", userController.RegisterUser)
		apiGroup.POST("/user/register/spendings", spendingController.CreateSpending)
	}

	authRoutes := router.Group("/")
	authRoutes.Use(middlewares.AuthMiddleware())
	{
		authRoutes.GET("/entradas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "entradas.html", gin.H{
				"title": "Entradas Page",
			})
		})
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
