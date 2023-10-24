package main

import (
	"github.com/jmoiron/sqlx"

	"api/internal/currency/delivery/http"
	"api/internal/currency/repo/postgres"
	"api/internal/currency/service"
	"api/ticker"
)

//TODO: ЗАПУСК нашего приложения

func main() {

}

func run() error {
	// Инициализация конфига

	// Инициализация базы данных
	db, err := sqlx.Open("", "")
	if err != nil {
		return err
	}

	// Repo - слой репозитория. Способ взаимодействия с данными конкретного хранилища.
	currencyRepo := postgres.New(
		db,
	)

	// Service - слой бизнес-логики. Способ получения информация или ее обратки в соответствии с заданным алгоритмом.
	currencyService := service.NewService(
		currencyRepo, // Contract
	)

	ticker.NewTick(0).LoopAccept(func() error {
		//return currencyService.UpdateWell()
	})

	// Delivery - слой доставки. Способ обратиться к бизнес логике (или способ обращения к сервису).
	currencyDelivery := http.New(
		currencyService, // Contract
	)

	return currencyDelivery.Listen("localhost:8090")
}
