package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/services"
)

type SpendingController struct {
	spendingService *services.SpendingService
	userService     *services.UserService
}

func NewSpendingController(spendingService *services.SpendingService, userService *services.UserService) *SpendingController {
	return &SpendingController{
		spendingService: spendingService,
		userService:     userService,
	}
}

func (sc *SpendingController) CreateSpending(c *gin.Context) {
	var spending domain.Spending
	if err := c.ShouldBindJSON(&spending); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
		return
	}

	user, err := sc.userService.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não foi possível identificar o usuário"})
		return
	}

	spending.Author = user.FullName
	spending.CreatedAt = time.Now()
	spending.UpdatedAt = time.Now()

	if err := sc.spendingService.CreateSpending(c, &spending); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar gasto"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gasto criado com sucesso", "spending": spending})
}
