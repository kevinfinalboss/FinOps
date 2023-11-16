package api

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/api/token"
	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/repository"
	"github.com/kevinfinalboss/FinOps/internal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo        *repository.UserRepository
	userService *services.UserService
}

func NewUserController(repo *repository.UserRepository, userService *services.UserService) *UserController {
	return &UserController{
		repo:        repo,
		userService: userService,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
		return
	}

	if !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de email inválido"})
		return
	}

	if uc.repo.IsEmailInUse(c, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email já está em uso"})
		return
	}

	user.ID = primitive.NewObjectID().Hex()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar senha"})
		return
	}
	user.PasswordHash = string(hashedPassword)

	if err := uc.repo.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var loginCreds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginCreds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
		return
	}

	user, err := uc.repo.FindUserByEmail(c.Request.Context(), loginCreds.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginCreds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	tokenString, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o token"})
		return
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o refresh token"})
		return
	}

	err = uc.userService.SaveRefreshToken(c.Request.Context(), user.ID, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o refresh token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	go services.SendLoginWebhook(user.FullName, user.Email, c.ClientIP())

	c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido", "token": tokenString})
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.repo.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) LogoutUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível identificar o usuário para logout"})
		return
	}

	err := uc.repo.RemoveRefreshToken(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover o refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout bem-sucedido"})
}
