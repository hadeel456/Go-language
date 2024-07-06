package main

import (
	"fmt"
	db "go-language/config"
	"go-language/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello world")
	db.Connect()
	app := fiber.New()

	// Middleware to log each request
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()
		log.Printf("[%s] %s %s %d %s\n", stop.Format(time.RFC3339), c.Method(), c.Path(), c.Response().StatusCode(), stop.Sub(start))
		return err
	})

	// Middleware to attach a header to all responses
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "This is a custom header")
		return c.Next()
	})

	routes.Book(app)
	routes.Author(app)

	app.Listen(":3000")
}
