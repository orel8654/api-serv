package main

import (
	"api/internal/currency/handlers"
	"api/internal/currency/repo"
	"api/internal/currency/service"
)

//TODO: ЗАПУСК нашего приложения

func main() {

}

func run() error {
	// log
	// config
	//

	// Инициализируем структуру Repo
	currencyRepo := repo.New() //   repo

	// Инициализируем структуру Service.
	// конструктор NewService принимает контракт Contract, этот контракт реализирует наш currencyRepo
	currencyService := service.NewService(
		currencyRepo,
	)

	// Инициализируем структуру Handler
	// Она принимает контракт Service - данный контракт реализует currencyService
	currencyDelivery := handlers.New(
		currencyService,
	)

	_ = currencyDelivery

	// У currencyDelivery можно вызывать различные методу тут.
	// В частности для него можно создать метод Listen который вызовешь и у тебя будет запускаться твой сервак.

	return nil
}
