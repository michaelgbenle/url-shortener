package routes

import (
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/michaelgbenle/url-shortener/database"
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
	r2 := database.CreateClient(1)
	defer r2.Close()
	val,err:=r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil{
		_=r2.Set(database.Ctx,c.IP(),os.Getenv("API_QUOTA"), 30*time.Minute).Err()
	}
	}


	//check if input is an actual url
	if !govalidator .IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid url"})
	}
	//check for domain error
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":""})
	}
	//enforce https,SSL
	body.URL= helpers.EnforceHTTP(body.URL)

}