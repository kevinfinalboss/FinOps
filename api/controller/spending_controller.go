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
	parsedDate, err := time.Parse(time.RFC3339, spending.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido"})
		return
	}
	spending.Date = parsedDate.Format("02/01/2006")

	if err := sc.spendingService.CreateSpending(c, &spending); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar gasto"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gasto criado com sucesso", "spending": spending})
}

func (sc *SpendingController) GetRecentSpendings(c *gin.Context) {
	spendings, err := sc.spendingService.GetRecentSpendings(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gastos recentes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"recent_spendings": spendings})
}

func (sc *SpendingController) GetSpendingsSumByMonth(c *gin.Context) {
	month := c.Query("month")
	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mês não especificado"})
		return
	}

	spendings, err := sc.spendingService.GetSpendingsByMonth(c.Request.Context(), month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gastos"})
		return
	}

	var total float64
	for _, spending := range spendings {
		total += spending.Value
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
