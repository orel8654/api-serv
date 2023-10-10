package handlers

import "context"

type Service interface {
	GetCurrency(ctx context.Context, to string) (bool, error)
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
	h.s.GetCurrency(context.Background(), "a")
	//... some logic
}
