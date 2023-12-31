package handlers

import (
	"fmt"
	"net/http"
	"time"

	"api/currencies"
	"api/database"
	"api/ticker"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

const UrlsSelf = "http://127.0.0.1:3000/api/tick/update"

type Handler struct {
	app *fiber.App
	db  *database.Storage
	tk  *ticker.Tick
	ex  *currencies.Currency
}

func NewHandler(confDb types.ConfDB, confApi types.ConfAPI) (*Handler, error) {
	d, err := database.NewStorage(confDb)
	if err != nil {
		return nil, err
	}
	fn := func() error {
		_, err := http.Get(UrlsSelf)
		if err != nil {
			return err
		}
		return nil
	}
	h := &Handler{
		app: fiber.New(),
		db:  d,
		tk:  ticker.NewTick(1 * time.Minute),
		ex:  currencies.NewCurrency(confApi),
	}
	h.app.Get("/api/currency", h.GetRows)
	h.app.Post("/api/currency", h.CreateRow)
	h.app.Put("/api/currency", h.UpdateRowWell)
	h.app.Get("/api/tick/update", h.UpdateTick)
	h.tk.LoopAccept(fn)
	return h, nil
}

func (h *Handler) Listen(host string) error {
	return h.app.Listen(host)
}

func (h *Handler) UpdateRowWell(ctx *fiber.Ctx) error {
	var payload types.DataPut
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	err := h.db.Exists(
		types.DataPost{
			CurrencyTo: payload.CurrencyTo,
		},
	)
	if err != nil {
		return err
	}
	if err := h.db.UpdateWell(payload); err != nil {
		return err
	}
	return ctx.JSON(payload)
}

func (h *Handler) CreateRow(ctx *fiber.Ctx) error {
	var payload types.DataPost
	if err := ctx.BodyParser(&payload); err != nil {
		fmt.Println(err)
		return err
	}
	if err := h.db.Exists(payload); err == nil {
		fmt.Println(err)
		return fmt.Errorf("row exist")
	}
	if err := h.db.WriteRow(ctx.Context(), payload); err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.JSON(payload)
}

func (h *Handler) GetRows(ctx *fiber.Ctx) error {
	q, err := h.db.SelectAll()
	if err != nil {
		fmt.Println(err)
		return ctx.SendString("error")
	}
	return ctx.JSON(q)
}

func (h *Handler) UpdateTick(ctx *fiber.Ctx) error {
	data, err := h.ex.GetNewCurrency()
	if err != nil {
		fmt.Println(err)
		ctx.Status(500)
		return ctx.SendString("fail")
	}
	if err = h.db.UpdateRows(data); err != nil {
		fmt.Println(err)
		ctx.Status(500)
		return ctx.SendString("fail")
	}
	return ctx.SendString("updated")
}
