package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/url-shortener/routes"
)

func main() {
	err:= godotenv.Load()
	if err != nil {
		panic(err)
	}
		app:= fiber.New()
		app.use(logger.New())
}

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1/", routes.ShortenUrl)


}