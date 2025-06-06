package main

import (
	"Juan314/backend/reservas/config"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	config.InitDB()

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())
	go func() {
		if err := app.Listen(port); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	app.Get("/reservas", internal/handlers.GetReservas)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	app.Shutdown()

}
