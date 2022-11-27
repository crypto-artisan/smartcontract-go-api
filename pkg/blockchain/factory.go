package blockchain

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

var FactoryInstance *Factory

const (
	FactoryAddress = "0x85ecb996f88ebe774a4769e14b1c958e3f8f8c63"
)

type Factory struct {
	Address  common.Address
	Instance *ethereum.Factory
}

func (clientCon ClientConnection) newFactory(address string) *ethereum.Factory {
	FactoryInstance = new(Factory)

	contractAddress := common.HexToAddress(address)

	FactoryInstance.Address = contractAddress

	instance, err := ethereum.NewFactory(contractAddress, clientCon.Client)
	if err != nil {
		log.Fatalln("Cannot get Factory contract at address ", address, " due to: ", err)
	}

	return instance
}

// GetLogicAddress 0x505A066E89Be22D3e56f16e1666de31f9328572e
func (clientCon ClientConnection) GetLogicAddress() common.Address {

	address, err := getFactory().CustodianWalletLogic(clientCon.callOpts)

	if err != nil {
		log.Fatalln("Cannot make call to Factory at ", FactoryInstance.Address, "due to: ", err)
	}

	return address
}

// GetEscrowAddress 0xd27adc3848dE1324AF87e5C235355e4a017Aa1CF
func (clientCon ClientConnection) GetEscrowAddress() common.Address {

	address, err := getFactory().EscrowContractAddress(clientCon.callOpts)

	if err != nil {
		log.Fatalln("Cannot make call to Factory at ", FactoryInstance.Address, "due to: ", err)
	}

	clientCon.postTransact()

	return address
}

func (clientCon ClientConnection) NewWallet(uuid string) (*types.Transaction, error) {

	trx, err := getFactory().NewCustodian(clientCon.trxOpts, uuid)

	if err != nil {
		log.Printf("Cannot create new wallet %v due to error: %v", uuid, err)

		return new(types.Transaction), err
	}

	clientCon.postTransact()

	return trx, nil
}

func (clientCon ClientConnection) GetAccountByUUID(uuid string) (*common.Address, error) {

	address, err := getFactory().Accounts(clientCon.callOpts, uuid)

	if err != nil {
		log.Printf("Cannot get account: %v due to error %v", uuid, err)

		return new(common.Address), err
	}

	return &address, nil
}

func getFactory() *ethereum.Factory {
	return CurrentConnection.newFactory(FactoryAddress)
}
