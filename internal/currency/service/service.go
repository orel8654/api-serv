package service

import (
	"context"
)

type Repo interface {
	CurrencyExists(ctx context.Context, to string) (bool, error)
}

type Service struct {
	repo Repo
	//repo *repo.Repo // Конкретная реализация
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCurrency(ctx context.Context, to string) (bool, error) {
	s.repo.CurrencyExists(context.Background(), to)
	return false, nil
}
