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

func (s *SpendingService) GetRecentSpendings(ctx context.Context) ([]domain.Spending, error) {
	return s.spendingRepo.GetRecentSpendings(ctx)
}

func (s *SpendingService) GetSpendingsByMonth(ctx context.Context, month string) ([]domain.Spending, error) {
	return s.spendingRepo.GetSpendingsByMonth(ctx, month)
}

func (s *IncomeService) GetRecentIncomes(ctx context.Context) ([]domain.Income, error) {
	return s.incomeRepo.GetRecentIncomes(ctx)
}

func (s *IncomeService) GetIncomesSum(ctx context.Context) (float64, error) {
	return s.incomeRepo.GetIncomesSum(ctx)
}
