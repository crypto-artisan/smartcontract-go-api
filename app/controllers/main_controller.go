package controllers

import "github.com/gofiber/fiber/v2"

func RenderHello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello from api",
	})
}