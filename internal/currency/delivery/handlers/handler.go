package handlers

import (
	"api/internal/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

// Currency - Контракт
type Currency interface {
	UpdateWellRepo(ctx context.Context, newData types.DataPut) error
	SelectRowsRepo(ctx context.Context) ([]types.DatabaseFields, error)
	WriteRowRepo(ctx context.Context, data types.DataPost) error
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

	h.app.Get("/api/currency", h.GetRows)
	h.app.Post("/api/currency", h.CreateRow)
	h.app.Put("/api/currency", h.UpdateWellFiber)

	return h
}

func (h *Handler) UpdateWellFiber(ctx *fiber.Ctx) error {
	var payload types.DataPut
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	if err := h.currency.UpdateWellRepo(ctx.Context(), payload); err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(payload)
}

func (h *Handler) GetRows(ctx *fiber.Ctx) error {
	data, err := h.currency.SelectRowsRepo(ctx.Context())
	if err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(data)
}

func (h *Handler) CreateRow(ctx *fiber.Ctx) error {
	var payload types.DataPost
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.SendString(err.Error())
	}
	if err := h.currency.WriteRowRepo(ctx.Context(), payload); err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(payload)
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}
