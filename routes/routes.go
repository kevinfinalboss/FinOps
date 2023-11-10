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
	userService := services.NewUserService(userRepo)
	userController := api.NewUserController(userRepo, userService)

	spendingService := services.NewSpendingService(spendingRepo)
	spendingController := api.NewSpendingController(spendingService, userService)

	authMiddleware := middlewares.AuthMiddleware(userService)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/healthcheck", healthCheck)
		apiGroup.POST("/user/login", userController.LoginUser)
		apiGroup.POST("/user/register", userController.RegisterUser)

		apiGroup.Use(authMiddleware)
		apiGroup.POST("/user/register/spendings", spendingController.CreateSpending)
	}

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Página de Login",
		})
	})

	router.GET("/entradas", authMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "entradas.html", gin.H{
			"title": "Página de Entradas",
		})
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
