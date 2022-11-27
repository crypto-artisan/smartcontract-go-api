package blockchain

import (
	"github.com/alofeoluwafemi/go-ethereum-api/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

var WalletInstance *Wallet

var (
	WalletAddress string
)

type Order struct {
	Seller   string `json:"seller"`   // P2P Seller Address
	Receiver string `json:"receiver"` // P2P Buyer Address
	Amount   int64  `json:"amount"`   // Amount in USDC Buyer wants from Seller for Fiat
	Rate     int64  `json:"rate"`     // Rate e.g 702 NGN / 1 USD
	Fee      int64  `json:"fee"`      // 1 USD e.g 1% of 100USD
}

type Wallet struct {
	Address  common.Address
	Instance *ethereum.CustodianWalletLogic
}

func (clientCon ClientConnection) newWallet(address string) *ethereum.CustodianWalletLogic {
	WalletInstance = new(Wallet)

	contractAddress := common.HexToAddress(address)

	WalletInstance.Address = contractAddress

	instance, err := ethereum.NewCustodianWalletLogic(contractAddress, clientCon.Client)
	if err != nil {
		log.Fatalln("Cannot get wallet address at ", address, " due to: ", err)
	}

	return instance
}

func (clientCon ClientConnection) OpenBuyOrder(order Order) (*types.Transaction, error) {

	trx, err := getWalletLogic().NewBuyOrder(
		clientCon.trxOpts,
		common.HexToAddress(order.Seller),
		common.HexToAddress(order.Receiver),
		big.NewInt(order.Amount),
		big.NewInt(order.Rate),
		big.NewInt(order.Fee),
	)

	if err != nil {
		log.Printf("Cannot open order : %v", err)

		return new(types.Transaction), err
	}

	clientCon.postTransact()

	return trx, nil
}

func (clientCon ClientConnection) GetUSDCBalance() (*big.Int, error) {

	balance, err := getWalletLogic().GetTotalBalance(clientCon.callOpts)

	if err != nil {
		log.Printf("Cannot open order : %v", err)

		return new(big.Int), err
	}

	return balance, nil
}

func getWalletLogic() *ethereum.CustodianWalletLogic {
	return CurrentConnection.newWallet(WalletAddress)
}
