package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/url-shortener/routes"
)

func main() {
	err:= godotenv.Load()
app:= fiber.New()
}

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1/", routes.ShortenUrl)


}