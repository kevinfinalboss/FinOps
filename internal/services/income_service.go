package services

import (
	"context"
	"time"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"github.com/kevinfinalboss/FinOps/internal/repository"
)

type IncomeService struct {
	incomeRepo *repository.IncomeRepository
}

func NewIncomeService(incomeRepo *repository.IncomeRepository) *IncomeService {
	return &IncomeService{
		incomeRepo: incomeRepo,
	}
}

func (s *IncomeService) CreateIncome(ctx context.Context, income *domain.Income) error {
	income.CreatedAt = time.Now()
	income.UpdatedAt = time.Now()
	return s.incomeRepo.CreateIncome(ctx, *income)
}
