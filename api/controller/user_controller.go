package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, _ := uc.repo.FindUserByEmail(c, user.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email já está em uso"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar senha"})
		return
	}
	user.PasswordHash = string(hashedPassword)

	err = uc.repo.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}
