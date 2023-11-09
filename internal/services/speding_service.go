package services

import (
	"context"
	"time"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/repository"
)

type SpendingService struct {
	spendingRepo *repository.SpendingRepository
}

func NewSpendingService(spendingRepo *repository.SpendingRepository) *SpendingService {
	return &SpendingService{
		spendingRepo: spendingRepo,
	}
}

func (s *SpendingService) CreateSpending(ctx context.Context, spending *domain.Spending) error {
	spending.CreatedAt = time.Now()
	spending.UpdatedAt = time.Now()
	return s.spendingRepo.CreateSpending(ctx, *spending)
}
