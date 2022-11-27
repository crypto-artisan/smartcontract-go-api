package controllers

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/blockchain"
	"github.com/gofiber/fiber/v2"
)

func SetUSDCTokenAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	type Request struct {
		Address string `json:"address"`
	}

	request := new(Request)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Malformed data",
			"data":    err,
		})
	}

	err := conn.SetUSDCAddress(request.Address)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   request.Address,
	})
}

func GetUSDCAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection

	address, err := conn.GetUSDCAddress()

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
	})
}

