package main

import (
	"api/internal/currency/handlers"
	"api/internal/currency/repo"
	"api/internal/currency/service"
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
	// config
	//

	handler := handlers.New( // delivey
		service.NewService( //  service
			repo.New(), //   repo
		),
	)
}
