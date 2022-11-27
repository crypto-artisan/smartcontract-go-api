// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TypesOrder is an auto generated low-level Go binding around an user-defined struct.
type TypesOrder struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ochestrator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"ApproveRejectedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"ClosedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"OpenOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"RejectedOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"closeOpenOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"consentOrderRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getOpenOrdersOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"getOrderById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_orderType\",\"type\":\"uint8\"}],\"name\":\"newOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"numberOfOpenOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ochestrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"rejectOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdcContractAddress\",\"type\":\"address\"}],\"name\":\"setUsdcTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeesEarned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620020b5380380620020b58339818101604052810190620000379190620002fd565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000604051806101400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600060ff168152602001600160ff168152602001428152602001428152509050600081908060018154018082558091505060019003906000526020600020906009020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff021916908360ff16021790555060e08201518160060160016101000a81548160ff021916908360ff16021790555061010082015181600701556101208201518160080155505050506200032f565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002c58262000298565b9050919050565b620002d781620002b8565b8114620002e357600080fd5b50565b600081519050620002f781620002cc565b92915050565b60006020828403121562000316576200031562000293565b5b60006200032684828501620002e6565b91505092915050565b611d76806200033f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636dee20321161008c578063b00f215c11610066578063b00f215c1461022f578063b1fab21c1461024b578063ef3bbbff1461027b578063f3fd622f14610297576100cf565b80636dee2032146101a85780639c811d43146101d8578063a85c38ef146101f6576100cf565b806311a21b2c146100d457806311eac855146100de5780631e7cea80146100fc578063240f41601461012c57806332c011171461014857806342b50a7a14610178575b600080fd5b6100dc6102b3565b005b6100e66104a9565b6040516100f3919061137e565b60405180910390f35b61011660048036038101906101119190611400565b6104cf565b604051610123919061144f565b60405180910390f35b61014660048036038101906101419190611400565b610500565b005b610162600480360381019061015d919061146a565b6105ed565b60405161016f919061144f565b60405180910390f35b610192600480360381019061018d9190611497565b610639565b60405161019f91906115ca565b60405180910390f35b6101c260048036038101906101bd919061146a565b6107e4565b6040516101cf9190611695565b60405180910390f35b6101e061087b565b6040516101ed919061137e565b60405180910390f35b610210600480360381019061020b9190611497565b6108a1565b6040516102269a999897969594939291906116c6565b60405180910390f35b61024960048036038101906102449190611400565b61097f565b005b6102656004803603810190610260919061178e565b610c24565b604051610272919061144f565b60405180910390f35b61029560048036038101906102909190611400565b610ff6565b005b6102b160048036038101906102ac919061146a565b61115f565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610343576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033a9061188d565b60405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103a0919061137e565b602060405180830381865afa1580156103bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e191906118c2565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016104629291906118ef565b6020604051808303816000875af1158015610481573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a59190611950565b5050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160205281600052604060002081815481106104eb57600080fd5b90600052602060002001600091509150505481565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461056e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610565906119c9565b60405180910390fd5b6000808281548110610583576105826119e9565b5b9060005260206000209060090201905060048160060160016101000a81548160ff021916908360ff1602179055507fbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd826040516105e0919061144f565b60405180910390a1505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b6106416112a2565b60008281548110610655576106546119e9565b5b9060005260206000209060090201604051806101400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff1660ff1660ff1681526020016006820160019054906101000a900460ff1660ff1660ff168152602001600782015481526020016008820154815250509050919050565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561086f57602002820191906000526020600020905b81548152602001906001019080831161085b575b50505050509050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081815481106108b157600080fd5b90600052602060002090600902016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900460ff1690806007015490806008015490508a565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e4906119c9565b60405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110610a4057610a3f6119e9565b5b90600052602060002001549050600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110610a9e57610a9d6119e9565b5b90600052602060002001600090556000808281548110610ac157610ac06119e9565b5b9060005260206000209060090201905042816008018190555060018160060160016101000a81548160ff021916908360ff1602179055508060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482858560030154866004015487600501548860060160009054906101000a900460ff1689600801548a60060160019054906101000a900460ff16604051610c169796959493929190611a18565b60405180910390a450505050565b60008673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8b90611ad3565b60405180910390fd5b60008511610cd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cce90611b3f565b60405180910390fd5b60006040518061014001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020018581526020018460ff168152602001600060ff16815260200142815260200160008152509050600081908060018154018082558091505060019003906000526020600020906009020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff021916908360ff16021790555060e08201518160060160016101000a81548160ff021916908360ff16021790555061010082015181600701556101208201518160080155505060006001600080549050610ef89190611b8e565b9050600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150558773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167e23c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2848b8b8b8b6000604051610fde96959493929190611c07565b60405180910390a48092505050979650505050505050565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611064576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105b90611cb4565b60405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106110b7576110b66119e9565b5b90600052602060002001549050600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110611115576111146119e9565b5b90600052602060002001600090557f6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac881604051611152919061144f565b60405180910390a1505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e69061188d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361125e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125590611d20565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b604051806101400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600060ff168152602001600060ff16815260200160008152602001600081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113688261133d565b9050919050565b6113788161135d565b82525050565b6000602082019050611393600083018461136f565b92915050565b600080fd5b6113a78161135d565b81146113b257600080fd5b50565b6000813590506113c48161139e565b92915050565b6000819050919050565b6113dd816113ca565b81146113e857600080fd5b50565b6000813590506113fa816113d4565b92915050565b6000806040838503121561141757611416611399565b5b6000611425858286016113b5565b9250506020611436858286016113eb565b9150509250929050565b611449816113ca565b82525050565b60006020820190506114646000830184611440565b92915050565b6000602082840312156114805761147f611399565b5b600061148e848285016113b5565b91505092915050565b6000602082840312156114ad576114ac611399565b5b60006114bb848285016113eb565b91505092915050565b6114cd8161135d565b82525050565b6114dc816113ca565b82525050565b600060ff82169050919050565b6114f8816114e2565b82525050565b6101408201600082015161151560008501826114c4565b50602082015161152860208501826114c4565b50604082015161153b60408501826114c4565b50606082015161154e60608501826114d3565b50608082015161156160808501826114d3565b5060a082015161157460a08501826114d3565b5060c082015161158760c08501826114ef565b5060e082015161159a60e08501826114ef565b506101008201516115af6101008501826114d3565b506101208201516115c46101208501826114d3565b50505050565b6000610140820190506115e060008301846114fe565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061161e83836114d3565b60208301905092915050565b6000602082019050919050565b6000611642826115e6565b61164c81856115f1565b935061165783611602565b8060005b8381101561168857815161166f8882611612565b975061167a8361162a565b92505060018101905061165b565b5085935050505092915050565b600060208201905081810360008301526116af8184611637565b905092915050565b6116c0816114e2565b82525050565b6000610140820190506116dc600083018d61136f565b6116e9602083018c61136f565b6116f6604083018b61136f565b611703606083018a611440565b6117106080830189611440565b61171d60a0830188611440565b61172a60c08301876116b7565b61173760e08301866116b7565b611745610100830185611440565b611753610120830184611440565b9b9a5050505050505050505050565b61176b816114e2565b811461177657600080fd5b50565b60008135905061178881611762565b92915050565b600080600080600080600060e0888a0312156117ad576117ac611399565b5b60006117bb8a828b016113b5565b97505060206117cc8a828b016113b5565b96505060406117dd8a828b016113b5565b95505060606117ee8a828b016113eb565b94505060806117ff8a828b016113eb565b93505060a06118108a828b016113eb565b92505060c06118218a828b01611779565b91505092959891949750929550565b600082825260208201905092915050565b7f463a206f6e6c79206f63686573747261746f7200000000000000000000000000600082015250565b6000611877601383611830565b915061188282611841565b602082019050919050565b600060208201905081810360008301526118a68161186a565b9050919050565b6000815190506118bc816113d4565b92915050565b6000602082840312156118d8576118d7611399565b5b60006118e6848285016118ad565b91505092915050565b6000604082019050611904600083018561136f565b6119116020830184611440565b9392505050565b60008115159050919050565b61192d81611918565b811461193857600080fd5b50565b60008151905061194a81611924565b92915050565b60006020828403121561196657611965611399565b5b60006119748482850161193b565b91505092915050565b7f433a206f6e6c792073656c6c6572000000000000000000000000000000000000600082015250565b60006119b3600e83611830565b91506119be8261197d565b602082019050919050565b600060208201905081810360008301526119e2816119a6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060e082019050611a2d600083018a611440565b611a3a6020830189611440565b611a476040830188611440565b611a546060830187611440565b611a6160808301866116b7565b611a6e60a0830185611440565b611a7b60c08301846116b7565b98975050505050505050565b7f433a20637573746f6d6572206f6e6c7900000000000000000000000000000000600082015250565b6000611abd601083611830565b9150611ac882611a87565b602082019050919050565b60006020820190508181036000830152611aec81611ab0565b9050919050565b7f433a20696e76616c6964206f7264657200000000000000000000000000000000600082015250565b6000611b29601083611830565b9150611b3482611af3565b602082019050919050565b60006020820190508181036000830152611b5881611b1c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b99826113ca565b9150611ba4836113ca565b9250828203905081811115611bbc57611bbb611b5f565b5b92915050565b6000819050919050565b6000819050919050565b6000611bf1611bec611be784611bc2565b611bcc565b6114e2565b9050919050565b611c0181611bd6565b82525050565b600060c082019050611c1c6000830189611440565b611c296020830188611440565b611c366040830187611440565b611c436060830186611440565b611c5060808301856116b7565b611c5d60a0830184611bf8565b979650505050505050565b7f433a206f6e6c7920627579657200000000000000000000000000000000000000600082015250565b6000611c9e600d83611830565b9150611ca982611c68565b602082019050919050565b60006020820190508181036000830152611ccd81611c91565b9050919050565b7f463a20696e76616c696420616464726573730000000000000000000000000000600082015250565b6000611d0a601283611830565b9150611d1582611cd4565b602082019050919050565b60006020820190508181036000830152611d3981611cfd565b905091905056fea264697066735822122036b90ca459dcf4f0bd8177369ebf87b988a395621ed523ffe562bcd32134dce964736f6c63430008100033",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// EscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EscrowMetaData.Bin instead.
var EscrowBin = EscrowMetaData.Bin

// DeployEscrow deploys a new Ethereum contract, binding an instance of Escrow to it.
func DeployEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, _ochestrator common.Address) (common.Address, *types.Transaction, *Escrow, error) {
	parsed, err := EscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EscrowBin), backend, _ochestrator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowCaller) GetOpenOrdersOf(opts *bind.CallOpts, _seller common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getOpenOrdersOf", _seller)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowSession) GetOpenOrdersOf(_seller common.Address) ([]*big.Int, error) {
	return _Escrow.Contract.GetOpenOrdersOf(&_Escrow.CallOpts, _seller)
}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowCallerSession) GetOpenOrdersOf(_seller common.Address) ([]*big.Int, error) {
	return _Escrow.Contract.GetOpenOrdersOf(&_Escrow.CallOpts, _seller)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowCaller) GetOrderById(opts *bind.CallOpts, _orderId *big.Int) (TypesOrder, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getOrderById", _orderId)

	if err != nil {
		return *new(TypesOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOrder)).(*TypesOrder)

	return out0, err

}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowSession) GetOrderById(_orderId *big.Int) (TypesOrder, error) {
	return _Escrow.Contract.GetOrderById(&_Escrow.CallOpts, _orderId)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowCallerSession) GetOrderById(_orderId *big.Int) (TypesOrder, error) {
	return _Escrow.Contract.GetOrderById(&_Escrow.CallOpts, _orderId)
}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowCaller) NumberOfOpenOrders(opts *bind.CallOpts, _seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "numberOfOpenOrders", _seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowSession) NumberOfOpenOrders(_seller common.Address) (*big.Int, error) {
	return _Escrow.Contract.NumberOfOpenOrders(&_Escrow.CallOpts, _seller)
}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowCallerSession) NumberOfOpenOrders(_seller common.Address) (*big.Int, error) {
	return _Escrow.Contract.NumberOfOpenOrders(&_Escrow.CallOpts, _seller)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowCaller) Ochestrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "ochestrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowSession) Ochestrator() (common.Address, error) {
	return _Escrow.Contract.Ochestrator(&_Escrow.CallOpts)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowCallerSession) Ochestrator() (common.Address, error) {
	return _Escrow.Contract.Ochestrator(&_Escrow.CallOpts)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowCaller) OpenOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "openOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.OpenOrders(&_Escrow.CallOpts, arg0, arg1)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowCallerSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.OpenOrders(&_Escrow.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_Escrow *EscrowCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Seller       common.Address
		Buyer        common.Address
		Receiver     common.Address
		Amount       *big.Int
		Rate         *big.Int
		Fee          *big.Int
		OrderType    uint8
		OrderStatus  uint8
		StartTime    *big.Int
		FulfiledTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OrderType = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.OrderStatus = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.StartTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FulfiledTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_Escrow *EscrowSession) Orders(arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	return _Escrow.Contract.Orders(&_Escrow.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_Escrow *EscrowCallerSession) Orders(arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	return _Escrow.Contract.Orders(&_Escrow.CallOpts, arg0)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowSession) UsdcToken() (common.Address, error) {
	return _Escrow.Contract.UsdcToken(&_Escrow.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowCallerSession) UsdcToken() (common.Address, error) {
	return _Escrow.Contract.UsdcToken(&_Escrow.CallOpts)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactor) CloseOpenOrder(opts *bind.TransactOpts, _seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "closeOpenOrder", _seller, _orderIndex)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowSession) CloseOpenOrder(_seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CloseOpenOrder(&_Escrow.TransactOpts, _seller, _orderIndex)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactorSession) CloseOpenOrder(_seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CloseOpenOrder(&_Escrow.TransactOpts, _seller, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactor) ConsentOrderRejected(opts *bind.TransactOpts, _buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "consentOrderRejected", _buyer, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowSession) ConsentOrderRejected(_buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ConsentOrderRejected(&_Escrow.TransactOpts, _buyer, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactorSession) ConsentOrderRejected(_buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ConsentOrderRejected(&_Escrow.TransactOpts, _buyer, _orderIndex)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowTransactor) NewOrder(opts *bind.TransactOpts, _seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "newOrder", _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowSession) NewOrder(_seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.Contract.NewOrder(&_Escrow.TransactOpts, _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowTransactorSession) NewOrder(_seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.Contract.NewOrder(&_Escrow.TransactOpts, _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowTransactor) RejectOrder(opts *bind.TransactOpts, _seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "rejectOrder", _seller, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowSession) RejectOrder(_seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.RejectOrder(&_Escrow.TransactOpts, _seller, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowTransactorSession) RejectOrder(_seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.RejectOrder(&_Escrow.TransactOpts, _seller, _orderId)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowTransactor) SetUsdcTokenAddress(opts *bind.TransactOpts, usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setUsdcTokenAddress", usdcContractAddress)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowSession) SetUsdcTokenAddress(usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetUsdcTokenAddress(&_Escrow.TransactOpts, usdcContractAddress)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowTransactorSession) SetUsdcTokenAddress(usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetUsdcTokenAddress(&_Escrow.TransactOpts, usdcContractAddress)
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowTransactor) WithdrawFeesEarned(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdrawFeesEarned")
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowSession) WithdrawFeesEarned() (*types.Transaction, error) {
	return _Escrow.Contract.WithdrawFeesEarned(&_Escrow.TransactOpts)
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowTransactorSession) WithdrawFeesEarned() (*types.Transaction, error) {
	return _Escrow.Contract.WithdrawFeesEarned(&_Escrow.TransactOpts)
}

// EscrowApproveRejectedOrderIterator is returned from FilterApproveRejectedOrder and is used to iterate over the raw logs and unpacked data for ApproveRejectedOrder events raised by the Escrow contract.
type EscrowApproveRejectedOrderIterator struct {
	Event *EscrowApproveRejectedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EscrowApproveRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowApproveRejectedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EscrowApproveRejectedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EscrowApproveRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowApproveRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowApproveRejectedOrder represents a ApproveRejectedOrder event raised by the Escrow contract.
type EscrowApproveRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproveRejectedOrder is a free log retrieval operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterApproveRejectedOrder(opts *bind.FilterOpts) (*EscrowApproveRejectedOrderIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return &EscrowApproveRejectedOrderIterator{contract: _Escrow.contract, event: "ApproveRejectedOrder", logs: logs, sub: sub}, nil
}

// WatchApproveRejectedOrder is a free log subscription operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchApproveRejectedOrder(opts *bind.WatchOpts, sink chan<- *EscrowApproveRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowApproveRejectedOrder)
				if err := _Escrow.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveRejectedOrder is a log parse operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) ParseApproveRejectedOrder(log types.Log) (*EscrowApproveRejectedOrder, error) {
	event := new(EscrowApproveRejectedOrder)
	if err := _Escrow.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowClosedOrderIterator is returned from FilterClosedOrder and is used to iterate over the raw logs and unpacked data for ClosedOrder events raised by the Escrow contract.
type EscrowClosedOrderIterator struct {
	Event *EscrowClosedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EscrowClosedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowClosedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EscrowClosedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EscrowClosedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowClosedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowClosedOrder represents a ClosedOrder event raised by the Escrow contract.
type EscrowClosedOrder struct {
	OrderId      *big.Int
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	FulfiledTime *big.Int
	OrderStatus  uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClosedOrder is a free log retrieval operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_Escrow *EscrowFilterer) FilterClosedOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*EscrowClosedOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &EscrowClosedOrderIterator{contract: _Escrow.contract, event: "ClosedOrder", logs: logs, sub: sub}, nil
}

// WatchClosedOrder is a free log subscription operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_Escrow *EscrowFilterer) WatchClosedOrder(opts *bind.WatchOpts, sink chan<- *EscrowClosedOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowClosedOrder)
				if err := _Escrow.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosedOrder is a log parse operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_Escrow *EscrowFilterer) ParseClosedOrder(log types.Log) (*EscrowClosedOrder, error) {
	event := new(EscrowClosedOrder)
	if err := _Escrow.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowOpenOrderIterator is returned from FilterOpenOrder and is used to iterate over the raw logs and unpacked data for OpenOrder events raised by the Escrow contract.
type EscrowOpenOrderIterator struct {
	Event *EscrowOpenOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EscrowOpenOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOpenOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EscrowOpenOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EscrowOpenOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOpenOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOpenOrder represents a OpenOrder event raised by the Escrow contract.
type EscrowOpenOrder struct {
	OrderId     *big.Int
	Seller      common.Address
	Buyer       common.Address
	Receiver    common.Address
	Amount      *big.Int
	Rate        *big.Int
	Fee         *big.Int
	OrderType   uint8
	OrderStatus uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpenOrder is a free log retrieval operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_Escrow *EscrowFilterer) FilterOpenOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*EscrowOpenOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &EscrowOpenOrderIterator{contract: _Escrow.contract, event: "OpenOrder", logs: logs, sub: sub}, nil
}

// WatchOpenOrder is a free log subscription operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_Escrow *EscrowFilterer) WatchOpenOrder(opts *bind.WatchOpts, sink chan<- *EscrowOpenOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOpenOrder)
				if err := _Escrow.contract.UnpackLog(event, "OpenOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpenOrder is a log parse operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_Escrow *EscrowFilterer) ParseOpenOrder(log types.Log) (*EscrowOpenOrder, error) {
	event := new(EscrowOpenOrder)
	if err := _Escrow.contract.UnpackLog(event, "OpenOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Escrow contract.
type EscrowOrderFulfilledIterator struct {
	Event *EscrowOrderFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EscrowOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOrderFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EscrowOrderFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EscrowOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOrderFulfilled represents a OrderFulfilled event raised by the Escrow contract.
type EscrowOrderFulfilled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*EscrowOrderFulfilledIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &EscrowOrderFulfilledIterator{contract: _Escrow.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *EscrowOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOrderFulfilled)
				if err := _Escrow.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFulfilled is a log parse operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_Escrow *EscrowFilterer) ParseOrderFulfilled(log types.Log) (*EscrowOrderFulfilled, error) {
	event := new(EscrowOrderFulfilled)
	if err := _Escrow.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowRejectedOrderIterator is returned from FilterRejectedOrder and is used to iterate over the raw logs and unpacked data for RejectedOrder events raised by the Escrow contract.
type EscrowRejectedOrderIterator struct {
	Event *EscrowRejectedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EscrowRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRejectedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EscrowRejectedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EscrowRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRejectedOrder represents a RejectedOrder event raised by the Escrow contract.
type EscrowRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRejectedOrder is a free log retrieval operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterRejectedOrder(opts *bind.FilterOpts) (*EscrowRejectedOrderIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return &EscrowRejectedOrderIterator{contract: _Escrow.contract, event: "RejectedOrder", logs: logs, sub: sub}, nil
}

// WatchRejectedOrder is a free log subscription operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchRejectedOrder(opts *bind.WatchOpts, sink chan<- *EscrowRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRejectedOrder)
				if err := _Escrow.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRejectedOrder is a log parse operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) ParseRejectedOrder(log types.Log) (*EscrowRejectedOrder, error) {
	event := new(EscrowRejectedOrder)
	if err := _Escrow.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
