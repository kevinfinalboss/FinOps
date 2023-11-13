package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/services"
)

type IncomeController struct {
	incomeService *services.IncomeService
	userService   *services.UserService
}

func NewIncomeController(incomeService *services.IncomeService, userService *services.UserService) *IncomeController {
	return &IncomeController{
		incomeService: incomeService,
		userService:   userService,
	}
}

func (ic *IncomeController) CreateIncome(c *gin.Context) {
	var income domain.Income
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
		return
	}

	user, err := ic.userService.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não foi possível identificar o usuário"})
		return
	}

	income.Author = user.FullName
	income.CreatedAt = time.Now()
	income.UpdatedAt = time.Now()

	parsedDate, err := time.Parse("02/01/2006", income.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido"})
		return
	}
	income.Date = parsedDate.Format(time.RFC3339)

	if err := ic.incomeService.CreateIncome(c, &income); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar entrada"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Entrada registrada com sucesso", "income": income})
}
