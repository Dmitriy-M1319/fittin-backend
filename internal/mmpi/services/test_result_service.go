package services

import (
	"context"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
)

type TestResultRepository interface {
	Create(ctx context.Context, result *models.TestResult) error
	GetByUUID(ctx context.Context, uuid string) (*models.TestResult, error)
	GetAll(ctx context.Context) ([]*models.TestResult, error)
	Delete(ctx context.Context, uuid string) error
}

type TestResultService struct {
	repository TestResultRepository
}

func NewTestResultService(r TestResultRepository) *TestResultService {
	return &TestResultService{repository: r}
}

func (s *TestResultService) AddNewResult(ctx context.Context, result *models.TestResult) error {
	return s.repository.Create(ctx, result)
}

func (s *TestResultService) GetByUUID(ctx context.Context, uuid string) (*models.TestResult, error) {
	return s.repository.GetByUUID(ctx, uuid)
}

func (s *TestResultService) GetAll(ctx context.Context) ([]*models.TestResult, error) {
	return s.repository.GetAll(ctx)
}

func (s *TestResultService) Delete(ctx context.Context, uuid string) error {
	return s.repository.Delete(ctx, uuid)
}
