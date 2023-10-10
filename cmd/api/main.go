package main

import (
	"api/internal/config"
	"api/internal/currency/handlers"
	"api/internal/currency/repo"
	"api/internal/currency/service"
	"fmt"
	"os"
)

// TODO: ЗАПУСК нашего приложения
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	// log
	config, err := config.NewConfig("./configs/some.yaml")
	if err != nil {
		return err
	}
	fmt.Println(config)
	handler := handlers.New(
		service.NewService(
			repo.New(*config),
		),
	)
	return handler.Listen(":3000")
}
