package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mikefmeyer/catchphrase-go-mongodb-rest-api/config"
	"github.com/mikefmeyer/catchphrase-go-mongodb-rest-api/routes"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.CatchphrasesRoute(api.Group("/catchphrases"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	setupRoutes(app)

	err = app.Listen(":8000")

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}