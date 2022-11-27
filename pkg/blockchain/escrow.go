package blockchain

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

var EscrowInstance *Escrow

const (
	EscrowAddress = "0xd27adc3848dE1324AF87e5C235355e4a017Aa1CF"
)

type Escrow struct {
	Address  common.Address
	Instance *ethereum.Escrow
}

func (clientCon ClientConnection) newEscrow(address string) *ethereum.Escrow {
	EscrowInstance = new(Escrow)

	contractAddress := common.HexToAddress(address)

	EscrowInstance.Address = contractAddress

	instance, err := ethereum.NewEscrow(contractAddress, clientCon.Client)
	if err != nil {
		log.Fatalln("Cannot get Factory contract at address ", address, " due to: ", err)
	}

	return instance
}

func (clientCon ClientConnection) SetUSDCAddress(address string) error {

	_, err := getEscrow().SetUsdcTokenAddress(clientCon.trxOpts, common.HexToAddress(address))

	if err != nil {
		log.Printf("Cannot set new USDC token on Escrow due to: %v", err)

		return err
	}

	clientCon.postTransact()

	return nil
}

func (clientCon ClientConnection) GetUSDCAddress() (*common.Address,error) {

	address, err := getEscrow().UsdcToken(clientCon.callOpts)

	if err != nil {
		log.Printf("Cannot get USDC token on Escrow due to: %v", err)

		return new(common.Address),err
	}

	return &address, nil
}

func getEscrow() *ethereum.Escrow {
	return CurrentConnection.newEscrow(EscrowAddress)
}
