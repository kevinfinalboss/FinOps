package services

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
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

func (s *UserService) GetUserFromToken(subject string) (*domain.User, error) {
	user, err := s.userRepo.FindUserByID(context.Background(), subject)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) SaveRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	return s.userRepo.SaveRefreshToken(ctx, userID, refreshToken)
}

func (s *UserService) RemoveRefreshToken(ctx context.Context, userID string) error {
	return s.userRepo.RemoveRefreshToken(ctx, userID)
}
