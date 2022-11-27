package controllers

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/blockchain"
	"github.com/gofiber/fiber/v2"
)

// DeployUSDC Deployed at 0xb3f2504110eeea6a522218048D0519B829788DDD
func DeployUSDC(c *fiber.Ctx) error {

	conn := blockchain.CurrentConnection

	address, transaction := conn.DeployUSDC()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
		"hash":    transaction.Hash(),
	})
}

//DeployFactory Deployed at  0x85ecb996f88ebe774a4769e14b1c958e3f8f8c63
func DeployFactory(c *fiber.Ctx) error {

	conn := blockchain.CurrentConnection

	address, transaction := conn.DeployFactory()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
		"hash":    transaction.Hash(),
	})
}
