package handlers

import (
	"api/internal/types"

	"github.com/gofiber/fiber/v2"
)

// Currency - Контракт
type Currency interface {
	UpdateWellRepo(ctx *fiber.Ctx, newData types.DataPut) error
}

type Handler struct {
	currency Currency
	app      *fiber.App
}

// Слой доставки принимает какую-то реализацию. Которая будет соответствовать interface-y Currency.
func New(currency Currency) *Handler {

	h := &Handler{
		currency: currency,
		app:      fiber.New(),
	}

	h.app.Put("/api/currency", h.UpdateWellFiber)

	return h
}

func (h *Handler) UpdateWellFiber(ctx *fiber.Ctx) error {
	// h.currency.UpdateWell(ctx)
	return nil
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}
