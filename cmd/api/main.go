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

	handler := handlers.New( // delivey
		service.NewService( //  service
			repo.New(), //   repo
		),
	)

}
