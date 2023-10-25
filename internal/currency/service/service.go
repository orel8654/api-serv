package service

import (
	"context"

	"api/internal/types"
)

// Repo - Контракт.
type Repo interface {
	CurrencyExists(ctx context.Context, from, to string) (bool, error)
	UpdateWell(ctx context.Context, newData types.DataPut) error
	WriteRow(ctx context.Context, data types.DataPost) error
	UpdateRows(ctx context.Context, newData types.CurrencyLatest) error
	SelectRow(ctx context.Context, data types.DatabaseFields) (types.DatabaseFields, error)
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

func (s *Service) GetCurrency(ctx context.Context, from, to string) (any, error) {
	return nil, nil
}

func (s *Service) UpdateWell(ctx context.Context, newData types.DataPut) error {
	// Вот тут может быть проверка: (Эта задача сервиса)
	// что newData.CurrencyFrom существует
	// что newData.CurrencyTo существует
	// что newData.Well > 0

	// s.repo.CurrencyExists(ctx, "a", "b")

	return s.repo.UpdateWell(
		ctx,
		newData,
	)
}

func exists() {

}
