package bootstrap

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/blockchain"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/env"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	app := fiber.New(fiber.Config{
		AppName:       "Go Ethereum API",
		CaseSensitive: true,
	})

	//Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	blockchain.New()

	//Setup routes
	router.InstallRouter(app)

	return app
}
