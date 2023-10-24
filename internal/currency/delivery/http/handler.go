package http

import (
	"context"
	"net/http"

	"api/types"
)

// Currency - Контракт
type Currency interface {
	GetCurrency(ctx context.Context, from, to string) (any, error)

	UpdateWell(ctx context.Context, newData types.DataPut) error
}

type Handler struct {
	currency Currency
}

// Слой доставки принимает какую-то реализацию. Которая будет соответствовать interface-y Currency.

func New(currency Currency) *Handler {
	// Инициализироваться framework (для http, grpc ...)
	h := &Handler{
		currency: currency,
	}

	http.HandleFunc("/api/currency/v1/", h.GetCurrency)
	http.HandleFunc("/api/currencny/....", h.UpdateWell)

	return h
}

// HTTP: GET /api/currency/v1/...

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	h.currency.GetCurrency(context.Background(), "a", "b")
}

func (h *Handler) UpdateWell(w http.ResponseWriter, r *http.Request) {
	h.currency.UpdateWell(r.Context(), types.DataPut{})
}

func (h *Handler) Listen(host string) error {
	return http.ListenAndServe(host, nil)
}
