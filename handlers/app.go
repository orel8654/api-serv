package handlers

import "github.com/gofiber/fiber/v2"

type Handler struct {
	app *fiber.App
}

func NewHandler() *Handler {
	h := &Handler{
		app: fiber.New(),
	}
	h.app.Get("/api/currency", h.GetRows)
	h.app.Post("/api/currency", h.CreateRowWell)
	h.app.Put("/api/currency", h.CreateRow)
	return h
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}

func (h *Handler) CreateRow(ctx *fiber.Ctx) error {
	return ctx.SendString("Start endpoint")
}

func (h *Handler) CreateRowWell(ctx *fiber.Ctx) error {
	return ctx.SendString("Start endpoint")
}

func (h *Handler) GetRows(ctx *fiber.Ctx) error {
	return ctx.SendString("Start endpoint")
}
