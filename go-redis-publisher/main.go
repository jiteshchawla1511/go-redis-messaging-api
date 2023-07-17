package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "0.tcp.in.ngrok.io:16288",
})

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
	Location string `json:"location"`
}

var ctx = context.Background()

func main() {

	app := fiber.New()

	app.Post("/submit", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(&user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		err = redisClient.Publish(ctx, "send-user-data", payload).Err()

		if err != nil {
			panic(err)
		}

		log.Print(user)

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
