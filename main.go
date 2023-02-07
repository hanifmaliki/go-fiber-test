// https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hanifmaliki/go-fiber-test/config"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	listener_port := config.Config("LISTENER_PORT")

	if listener_port == "" {
		listener_port = "3000"
	}

	// Listen on PORT 3000
	err := app.Listen(":" + listener_port)

	if err != nil {
		fmt.Print(err)
	}
}
