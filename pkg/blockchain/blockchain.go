package blockchain

import (
	"context"
	standardcryto "crypto"
	"crypto/ecdsa"
	"fmt"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/env"
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var CurrentConnection *ClientConnection

type ClientConnection struct {
	RpcURL      string
	SignerKey   string
	SignerPublicAddress common.Address
	Client      *ethclient.Client
	NetworkId   *big.Int
	NetworkName string
	ctx         context.Context
	callOpts    *bind.CallOpts
	trxOpts     *bind.TransactOpts
}

func New() {
	clientCon := new(ClientConnection)
	clientCon.RpcURL = env.GetEnv("RPC_URL", "")

	client, err := ethclient.Dial(clientCon.RpcURL)

	if err != nil {
		log.Fatalln("Could not connect to client: ", err)
	}

	clientCon.Client = client
	clientCon.SignerKey = env.GetEnv("DEPLOYER_PRIVATE_KEY", "")
	clientCon.NetworkName = "Goerli"
	clientCon.NetworkId = new(big.Int).SetInt64(5)
	clientCon.ctx = context.Background()

	clientCon.setTransactOpts()
	clientCon.setCallOpts()

	CurrentConnection = clientCon
}

func (clientCon *ClientConnection) nonceAt(address common.Address) uint64 {

	nonce, err := clientCon.Client.PendingNonceAt(clientCon.ctx, address)

	fmt.Println("nonce: ", nonce)
	if err != nil {
		log.Fatalln("Cannot get nonce: ", err)
	}

	return nonce
}

func (clientCon *ClientConnection) setTransactOpts() *bind.TransactOpts {
	signer, _, publicKeyECDSA, _ := clientCon.extractSignerKeys()

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	opts, err := bind.NewKeyedTransactorWithChainID(signer, clientCon.NetworkId)
	if err != nil {
		log.Fatalln("Error with NewKeyedTransactorWithChainID ", err)
	}

	priorityFee, err := clientCon.Client.SuggestGasTipCap(clientCon.ctx)
	if err != nil {
		log.Fatalln("Cannot suggest Tip: ", err)
	}
	currentNonce := new(big.Int).SetUint64(clientCon.nonceAt(from))
	opts.Nonce = currentNonce
	opts.GasTipCap = priorityFee

	clientCon.SignerPublicAddress = from
	clientCon.trxOpts = opts

	return opts
}

func (clientCon *ClientConnection) setCallOpts() *bind.CallOpts {
	opts := new(bind.CallOpts)
	_, _, publicKeyEcdsa, err := clientCon.extractSignerKeys()

	if err != nil {
		log.Fatalln("Unable to extract private key: ", err)
	}

	opts.From = crypto.PubkeyToAddress(*publicKeyEcdsa)

	clientCon.callOpts = opts

	return opts
}

func (clientCon *ClientConnection) extractSignerKeys() (*ecdsa.PrivateKey, standardcryto.PublicKey, *ecdsa.PublicKey, error) {
	privateKey, err := crypto.HexToECDSA(clientCon.SignerKey)
	if err != nil {
		log.Fatalln("Invalid private key: ", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return privateKey, publicKey, publicKeyECDSA, err
}

func (clientCon *ClientConnection) ExtractSignerKeys() (*ecdsa.PrivateKey, standardcryto.PublicKey, *ecdsa.PublicKey, error) {
	return clientCon.extractSignerKeys()
}

func (clientCon *ClientConnection) DeployUSDC() (common.Address, *types.Transaction) {
	contract, trx, _, err := ethereum.DeployUSDC(clientCon.trxOpts, clientCon.Client)

	if err != nil {
		log.Fatalln("Could not deploy USDC contract: ", err)
	}

	clientCon.postTransact()

	return contract, trx
}

func (clientCon *ClientConnection) DeployFactory() (common.Address, *types.Transaction) {
	contract, trx, _, err := ethereum.DeployFactory(clientCon.trxOpts, clientCon.Client)

	if err != nil {
		log.Fatalln("Could not deploy Factory contract: ", err)
	}

	clientCon.postTransact()

	return contract, trx
}

func (clientCon ClientConnection) postTransact() {
	clientCon.trxOpts.Nonce = new(big.Int).SetUint64(clientCon.nonceAt(clientCon.SignerPublicAddress))
}