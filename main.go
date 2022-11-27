package main

import (
	"fmt"
	"github.com/alofeoluwafemi/go-ethereum-api/bootstrap"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/blockchain"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplication()

	defer func() {
		fmt.Println("Closing connection...")
		 blockchain.CurrentConnection.Client.Close()
	}()

	log.Fatal(app.Listen(
		fmt.Sprintf("%s:%s",
			env.GetEnv("APP_HOST", "localhost"),
			env.GetEnv("APP_PORT", "3000"),
		)))
}
