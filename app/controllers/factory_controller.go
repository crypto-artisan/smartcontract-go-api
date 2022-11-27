package controllers

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/blockchain"
	"github.com/gofiber/fiber/v2"
)

func GetCustodianWalletLogicAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection

	address := conn.GetLogicAddress()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
	})
}

func GetEscrowAddress(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection

	address := conn.GetEscrowAddress()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"address": address.String(),
	})
}

func NewWallet(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	uuid := c.Params("uuid")

	trx, err := conn.NewWallet(uuid)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"hash":   trx.Hash(),
	})
}

func GetWallet(c *fiber.Ctx) error {
	conn := blockchain.CurrentConnection
	uuid := c.Params("uuid")

	address, err := conn.GetAccountByUUID(uuid)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"hash":   address.String(),
	})
}

//UUID									| Account Address
//b93e42b0-33a2-11ed-a261-0242ac120002 	| 0xC1F07Db647Aa3002c12BbaF8D598F0ef19c4ddd3
//cec3df26-339a-11ed-a261-0242ac120002 	| 0x63cc2dd4d1836ba66b40b3debfca7896256c24d0
//cec3dd14-339a-11ed-a261-0242ac120002 	| 0x61fe251efdcb04cf95d3f75ffa7d9e9d466eae23
