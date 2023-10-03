package main

import (
	"api/config"
	"log"
)

func main() {
	// postgresql юзать ==> sqlx, pgx
	configDB, err := config.ConfigDB("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	configAPI, err := config.ConfigAPI("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

}
