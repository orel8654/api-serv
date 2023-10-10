package handlers

import "context"

type Service interface {
	GetCurrency(ctx context.Context, from, to string) (any, error)
}

type Handler struct {
	s Service
}

func New(s Service) *Handler {
	return &Handler{
		s: s,
	}
}

// HTTP: GET /api/currency/v1/...

func (h *Handler) GetCurrency() {
	h.s.GetCurrency(context.Background(), "a", "b")
	//... some logic
}
