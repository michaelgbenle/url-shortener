package main

import "github.com/michaelgbenle/url-shortener/routes"

func main() {

}

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1/", routes.ShortenUrl)

}