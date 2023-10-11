package handlers

import "context"

type Service interface {
	GetCurrency(ctx context.Context, from, to string) (any, error)
}

type Handler struct {
	s Service
}

// New - Функция принимает контракт
// этот контракт будет реализовывать какой-нибудь тип
func New(s Service) *Handler {
	// Тут можно делать инициализацию fiber

	return &Handler{
		s: s,
	}
}

// HTTP: GET /api/currency/v1/...

func (h *Handler) GetCurrency() {
	h.s.GetCurrency(context.Background(), "a", "b")
	//... some logic
}
