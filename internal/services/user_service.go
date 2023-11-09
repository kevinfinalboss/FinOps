package services

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kevinfinalboss/FinOps/api/middlewares"
	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserFromContext(ctx *gin.Context) (*domain.User, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return nil, errors.New("usuário não encontrado no contexto")
	}

	user, err := s.userRepo.FindUserByID(context.Background(), userID.(string))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserFromToken(tokenString string) (*domain.User, error) {
	token, err := middlewares.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*middlewares.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("claims inválidos no token")
	}

	user, err := s.userRepo.FindUserByID(context.Background(), claims.Subject)
	if err != nil {
		return nil, err
	}

	return user, nil
}
