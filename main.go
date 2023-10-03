package main

import (
	"api/config"
	"api/handlers"
	"fmt"
	"log"
	"os"
)

func main() {
	configDB, err := config.ConfigDB("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	configAPI, err := config.ConfigAPI("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	h, err := handlers.NewHandler(*configDB, *configAPI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h.Listen(":3000")
}
