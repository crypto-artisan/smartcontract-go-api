package router

import (
	"github.com/gofiber/fiber/v2"
)

func InstallRouter(app *fiber.App) {
	setup(app, NewApiRouter())

	//If no route was matched
	app.Use(func(c *fiber.Ctx) error {
		err := c.SendStatus(404)
		panic(err)
	})
}
func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}