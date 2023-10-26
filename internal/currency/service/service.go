package service

import (
	"api/internal/types"
	"context"
)

// Repo - Контракт.
type Repo interface {
	UpdateWell(ctx context.Context, newData types.DataPut) error
	WriteRow(ctx context.Context, data types.DataPost) error
	SelectAll(ctx context.Context) ([]types.DatabaseFields, error)
}

type Service struct {
	repo Repo
}

// Сервис принимает какую-то реализацию. Которая будет соответствовать interface-y Repo.
func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) UpdateWellRepo(ctx context.Context, newData types.DataPut) error {
	// Вот тут может быть проверка: (Эта задача сервиса)
	// что newData.CurrencyFrom существует
	// что newData.CurrencyTo существует
	// что newData.Well > 0

	// s.repo.CurrencyExists(ctx, "a", "b")
	return s.repo.UpdateWell(ctx, newData)
}

func (s *Service) WriteRowRepo(ctx context.Context, data types.DataPost) error {
	return s.repo.WriteRow(ctx, data)
}

func (s *Service) SelectRowsRepo(ctx context.Context) ([]types.DatabaseFields, error) {
	q, err := s.repo.SelectAll(ctx)
	if err != nil {
		return []types.DatabaseFields{}, err
	}
	return q, nil
}
