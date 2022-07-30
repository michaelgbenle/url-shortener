package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/internal/uuid"
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
	}else{
		val,_=r2.Get(database.Ctx,c.IP()).Result()
		valInt,_ := strconv.Atoi(val)
		if valInt <= 0{
			limit,_:=r2.TTL(database.Ctx,c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":"Rate limit exceeded",
				"rate_limit_rest": limit/time.Nanosecond/time.Minute
			})
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
	var id string
	if body.CustomShort==""{
		id = uuid.New().String()[:6]
	} else {
		id= body.CustomShort
	}
r:= database.CreateClient(0)
defer r.Close()
val,_= r.Get(database.Ctx,id).Result()
if val != ""{
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error":"url custom short is already in use",
	})
}
if body.Expiry ==0{
	body.Expiry=24
}

err = r.Set(database.Ctx,id,body.URL,body.Expiry*3600*time.Second).Err()
		if err != nil{
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":"unable to connect to server",
			})
		}


	r2.Decr(database.Ctx,c.IP())

}