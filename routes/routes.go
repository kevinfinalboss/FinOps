package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/kevinfinalboss/FinOps/api/controller"
	"github.com/kevinfinalboss/FinOps/api/middlewares"
	"github.com/kevinfinalboss/FinOps/internal/repository"
	"github.com/kevinfinalboss/FinOps/internal/services"
)

func RegisterRoutes(router *gin.Engine, userRepo *repository.UserRepository, spendingRepo *repository.SpendingRepository, incomeRepo *repository.IncomeRepository) {
	userService := services.NewUserService(userRepo)
	userController := api.NewUserController(userRepo, userService)

	spendingService := services.NewSpendingService(spendingRepo)
	spendingController := api.NewSpendingController(spendingService, userService)

	incomeService := services.NewIncomeService(incomeRepo)
	incomeController := api.NewIncomeController(incomeService, userService)

	authMiddleware := middlewares.AuthMiddleware(userService)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/healthcheck", healthCheck)
		apiGroup.POST("/user/login", userController.LoginUser)
		apiGroup.POST("/user/register", userController.RegisterUser)

		spendingGroup := apiGroup.Group("/spendings")
		{
			spendingGroup.Use(authMiddleware)
			spendingGroup.POST("/register", spendingController.CreateSpending)
			spendingGroup.GET("/recent", spendingController.GetRecentSpendings)
			spendingGroup.GET("/sumByMonth", spendingController.GetSpendingsSumByMonth)
		}

		incomeGroup := apiGroup.Group("/incomes")
		{
			incomeGroup.Use(authMiddleware)
			incomeGroup.POST("/register", incomeController.CreateIncome)
			incomeGroup.GET("/recent", incomeController.GetRecentIncomes)
			incomeGroup.GET("/sumByMonth", incomeController.GetIncomesSumByMonth)
		}
	}

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "Página de Login"})
	})

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	router.GET("/entradas", authMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "entradas.html", gin.H{"title": "Página de Entradas"})
	})

	router.GET("/saidas", authMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "saidas.html", gin.H{"title": "Página de Saídas"})
	})

	router.GET("/relatorios", authMiddleware, func(c *gin.Context) {
		c.HTML(http.StatusOK, "relatorios.html", gin.H{"title": "Página de Relátorios"})
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
