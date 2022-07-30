package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelgbenle/url-shortener/database"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	r:= database.CreateClient(0)
	defer r.Close()

}