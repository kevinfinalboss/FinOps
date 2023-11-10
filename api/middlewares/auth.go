package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/internal/services"
	"github.com/kevinfinalboss/FinOps/api/token"
)

func AuthMiddleware(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token é necessário"})
			return
		}

		tkn, claims, err := token.ValidateToken(tokenString)
		if err != nil || !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}
		user, err := userService.GetUserFromToken(claims.Subject)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			return
		}

		c.Set("userID", user.ID)
		c.Next()
	}
}
