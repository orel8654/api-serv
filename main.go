package main

import (
	"api/config"
	"api/database"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateRow(c *fiber.Ctx) error {
	return c.SendString("Start endpoint")
}

func Conversion(c *fiber.Ctx) error {
	return c.SendString("Start endpoint")
}

func Agg(c *fiber.Ctx) error {
	return c.SendString("Start endpoint")
}

func Start(c *fiber.Ctx) error {
	return c.SendString("Start endpoint")
}

func HandlerUpdateCurrency(confDB *config.ConfDB, confAPI *config.ConfAPI) {
	fmt.Println("here")
	connectDB, err := database.Connect(confDB)
	if err != nil {
		log.Fatal(err)
	}
	defer connectDB.Close()
	NC := database.NewDB(connectDB)
	results, err := NC.SelectRows()
	fmt.Println(results)
}

func main() {
	configDB, err := config.ConfigDB("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	configAPI, err := config.ConfigAPI("./config/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(6 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				go HandlerUpdateCurrency(configDB, configAPI)
			}
		}
	}()
	defer func() {
		ticker.Stop()
		done <- true
	}()

	app := fiber.New()
	app.Get("/api", Start)
	app.Post("/api/currency", CreateRow)
	app.Put("/api/curency", Conversion)
	app.Get("/api/currency", Agg)
	log.Fatal(app.Listen(":8001"))
}
