package handlers

import (
	"api/internal/types"

	"github.com/gofiber/fiber/v2"
)

// Currency - Контракт
type Currency interface {
	UpdateWellRepo(newData types.DataPut) error
	SelectRowsRepo() ([]types.DatabaseFields, error)
	WriteRowRepo(data types.DataPost) error
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
	if err := h.currency.UpdateWellRepo(payload); err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(payload)
}

func (h *Handler) GetRows(ctx *fiber.Ctx) error {
	data, err := h.currency.SelectRowsRepo()
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
	if err := h.currency.WriteRowRepo(payload); err != nil {
		return ctx.SendString(err.Error())
	}
	return ctx.JSON(payload)
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}
