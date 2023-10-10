package service

import (
	"context"
)

type Repo interface {
	CurrencyExists(ctx context.Context, from, to string) (bool, error)
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

func (s *Service) GetCurrency(ctx context.Context, from, to string) (any, error) {
	s.repo.CurrencyExists(context.Background(), "a", "b")

	return nil, nil
}
