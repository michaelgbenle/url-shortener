package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/michaelgbenle/url-shortener/helpers"
)

type request struct {
	URL         	string `json:"url"`
	CustomShort 	string `json:"short"`
	Expiry      	time.Duration `json:"expiry"`
}

type response struct {
	URL				string	`json:"url"`
	CustomShort		string	`json:"custom_short"`
	Expiry			time.Duration	`json:"expiry"`
	XRateRemaining	int				`json:"rate_limit"`
	XRateLimitRest	time.Duration	`json:"rate_limit_reset"`
}

func ShortenUrl(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err!= nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JS"})
	}
	// implement rate limiting
	//check if input is an actual url
	if !govalidator .IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid url"})
	}
	//check for domain error
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":""})
	}
	//enforce https,SSL

}