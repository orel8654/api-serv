package main

import (
	"api/internal/config"
	"api/internal/currency/delivery/handlers"
	"api/internal/currency/repo/postgres"
	"api/internal/currency/service"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	// Инициализация конфига
	conf, err := config.NewConfig("./configs/some.yaml")
	if err != nil {
		return err
	}
	// Инициализация базы
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.Username, conf.Password, conf.Database, conf.Host, conf.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return err
	}

	// Repo - слой репозитория. Способ взаимодействия с данными конкретного хранилища.
	currencyRepo := postgres.New(
		db,
	)

	// Service - слой бизнес-логики. Способ получения информация или ее обратки в соответствии с заданным алгоритмом.
	currencyService := service.NewService(
		currencyRepo,
	)

	// Delivery - слой доставки. Способ обратиться к бизнес логике (или способ обращения к сервису).
	currencyDelivery := handlers.New(
		currencyService,
	)

	return currencyDelivery.Listen(fmt.Sprintf("%s:%s", conf.HostApp, conf.PortApp))
}
