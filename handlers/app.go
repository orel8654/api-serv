package handlers

import (
	"api/config"
	"api/currencies"
	"api/database"
	"api/ticker"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	app *fiber.App
	db  *database.Storage
	tk  *ticker.Tick
	ex  *currencies.Currency
}

func NewHandler(confDb config.ConfDB, confApi config.ConfAPI) (*Handler, error) {
	d, err := database.NewStorage(confDb)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		app: fiber.New(),
		db:  d,
		tk:  ticker.NewTick(),
		ex:  currencies.NewCurrency(confApi),
	}
	h.app.Get("/api/currency", h.GetRows)
	h.app.Post("/api/currency", h.CreateRowWell)
	h.app.Put("/api/currency", h.CreateRow)
	h.tk.Loop()
	return h, nil
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
	q, err := h.db.SelectAll()
	if err != nil {
		fmt.Println(err)
		return ctx.SendString("error")
	}
	r, err := json.Marshal(q)
	if err != nil {
		fmt.Println(err)
		return ctx.SendString("error")
	}
	return ctx.SendString(string(r))
}

func (h *Handler) UpdateTick(ctx *fiber.Ctx) error {
	data, err := h.ex.GetNewCurrency()
	if err != nil {
		fmt.Println(err)
		ctx.Status(500)
		return ctx.SendString("fail")
	}
	// update DB
	_ = data
	return ctx.SendString("updated")
}
