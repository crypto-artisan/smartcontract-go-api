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

// CustodianWalletLogicMetaData contains all meta data concerning the CustodianWalletLogic contract.
var CustodianWalletLogicMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"ApproveRejectedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"ClosedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"OpenOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"RejectedOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openOrderIndex\",\"type\":\"uint256\"}],\"name\":\"approveOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openOrderIndex\",\"type\":\"uint256\"}],\"name\":\"consentOrderRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenOrders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"newBuyOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"newSellOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"rejectOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506121a1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806391c024151161007157806391c0241514610178578063a85c38ef14610194578063bd4b8a04146101cd578063c45a0155146101e9578063dbe5bab514610207578063fdf61e6814610225576100a9565b806312b58349146100ae5780631971e03b146100cc5780631e7cea80146100fc57806324f638941461012c5780636282682a14610148575b600080fd5b6100b6610243565b6040516100c391906112de565b60405180910390f35b6100e660048036038101906100e19190611397565b610252565b6040516100f391906112de565b60405180910390f35b61011660048036038101906101119190611412565b61050e565b60405161012391906112de565b60405180910390f35b61014660048036038101906101419190611452565b61053f565b005b610162600480360381019061015d9190611397565b6106af565b60405161016f91906112de565b60405180910390f35b610192600480360381019061018d9190611452565b61074c565b005b6101ae60048036038101906101a99190611452565b6108e9565b6040516101c49a999897969594939291906114aa565b60405180910390f35b6101e760048036038101906101e29190611452565b6109c7565b005b6101f1610bf6565b6040516101fe9190611546565b60405180910390f35b61020f610c1c565b60405161021c919061161f565b60405180910390f35b61022d610c2b565b60405161023a91906112de565b60405180910390f35b600061024d610c3a565b905090565b60008061025d610d30565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cb9190611656565b9050600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff160361033c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610333906116e0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a29061174c565b60405180910390fd5b82856103b7919061179b565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231896040518263ffffffff1660e01b81526004016103f09190611546565b602060405180830381865afa15801561040d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043191906117e4565b1015610472576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104699061185d565b60405180910390fd5b61047a610d30565b73ffffffffffffffffffffffffffffffffffffffff1663b1fab21c88308989898960006040518863ffffffff1660e01b81526004016104bf97969594939291906118c2565b6020604051808303816000875af11580156104de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050291906117e4565b91505095945050505050565b6001602052816000526040600020818154811061052a57600080fd5b90600052602060002001600091509150505481565b6000610549610d30565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b815260040161058191906112de565b61014060405180830381865afa15801561059f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c39190611ae4565b90503073ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062e90611b5e565b60405180910390fd5b61063f610d30565b73ffffffffffffffffffffffffffffffffffffffff1663240f416030846040518363ffffffff1660e01b8152600401610679929190611b7e565b600060405180830381600087803b15801561069357600080fd5b505af11580156106a7573d6000803e3d6000fd5b505050505050565b60006106b9610d30565b73ffffffffffffffffffffffffffffffffffffffff1663b1fab21c30888888888860016040518863ffffffff1660e01b81526004016106fe9796959493929190611be2565b6020604051808303816000875af115801561071d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074191906117e4565b905095945050505050565b6000610756610dc8565b9050600081838151811061076d5761076c611c51565b5b602002602001015190506000610781610d30565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b81526004016107b991906112de565b61014060405180830381865afa1580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb9190611ae4565b90503073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161461086f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086690611ccc565b60405180910390fd5b610877610d30565b73ffffffffffffffffffffffffffffffffffffffff1663ef3bbbff30866040518363ffffffff1660e01b81526004016108b1929190611b7e565b600060405180830381600087803b1580156108cb57600080fd5b505af11580156108df573d6000803e3d6000fd5b5050505050505050565b600081815481106108f957600080fd5b90600052602060002090600902016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900460ff1690806007015490806008015490508a565b60006109d1610dc8565b90506000815111610a17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0e90611d38565b60405180910390fd5b6000818381518110610a2c57610a2b611c51565b5b602002602001015190506000610a40610d30565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b8152600401610a7891906112de565b61014060405180830381865afa158015610a96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aba9190611ae4565b90503073ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610b2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2590611da4565b60405180910390fd5b610b36610d30565b73ffffffffffffffffffffffffffffffffffffffff1663b00f215c30866040518363ffffffff1660e01b8152600401610b70929190611b7e565b600060405180830381600087803b158015610b8a57600080fd5b505af1158015610b9e573d6000803e3d6000fd5b50505050610bb9816040015182606001518360a00151610e55565b7fd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e82604051610be891906112de565b60405180910390a150505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060610c26610dc8565b905090565b6000610c356111b3565b905090565b6000610c44610d30565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb29190611656565b73ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610cea9190611546565b602060405180830381865afa158015610d07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2b91906117e4565b905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d226ae086040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc39190611656565b905090565b6060610dd2610d30565b73ffffffffffffffffffffffffffffffffffffffff16636dee2032306040518263ffffffff1660e01b8152600401610e0a9190611546565b600060405180830381865afa158015610e27573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610e509190611e91565b905090565b3073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ec3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eba90611f26565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2990611f92565b60405180910390fd5b60008211610f75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6c90611ffe565b60405180910390fd5b81610f7e6111b3565b1015610fbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb69061206a565b60405180910390fd5b610fc7610d30565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa158015611011573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110359190611656565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b815260040161106f929190611b7e565b6020604051808303816000875af115801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b291906120c2565b506110bb610d30565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa158015611105573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111299190611656565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb61114c610d30565b836040518363ffffffff1660e01b815260040161116a929190611b7e565b6020604051808303816000875af1158015611189573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ad91906120c2565b50505050565b6000806111be610dc8565b905060006111ca610c3a565b905060005b600183516111dd91906120ef565b8110156112bc5760006111ee610d30565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a85848151811061121c5761121b611c51565b5b60200260200101516040518263ffffffff1660e01b815260040161124091906112de565b61014060405180830381865afa15801561125e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112829190611ae4565b905080606001518361129491906120ef565b92508060a00151836112a691906120ef565b92505080806112b490612123565b9150506111cf565b50809250505090565b6000819050919050565b6112d8816112c5565b82525050565b60006020820190506112f360008301846112cf565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113388261130d565b9050919050565b6113488161132d565b811461135357600080fd5b50565b6000813590506113658161133f565b92915050565b611374816112c5565b811461137f57600080fd5b50565b6000813590506113918161136b565b92915050565b600080600080600060a086880312156113b3576113b2611303565b5b60006113c188828901611356565b95505060206113d288828901611356565b94505060406113e388828901611382565b93505060606113f488828901611382565b925050608061140588828901611382565b9150509295509295909350565b6000806040838503121561142957611428611303565b5b600061143785828601611356565b925050602061144885828601611382565b9150509250929050565b60006020828403121561146857611467611303565b5b600061147684828501611382565b91505092915050565b6114888161132d565b82525050565b600060ff82169050919050565b6114a48161148e565b82525050565b6000610140820190506114c0600083018d61147f565b6114cd602083018c61147f565b6114da604083018b61147f565b6114e7606083018a6112cf565b6114f460808301896112cf565b61150160a08301886112cf565b61150e60c083018761149b565b61151b60e083018661149b565b6115296101008301856112cf565b6115376101208301846112cf565b9b9a5050505050505050505050565b600060208201905061155b600083018461147f565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611596816112c5565b82525050565b60006115a8838361158d565b60208301905092915050565b6000602082019050919050565b60006115cc82611561565b6115d6818561156c565b93506115e18361157d565b8060005b838110156116125781516115f9888261159c565b9750611604836115b4565b9250506001810190506115e5565b5085935050505092915050565b6000602082019050818103600083015261163981846115c1565b905092915050565b6000815190506116508161133f565b92915050565b60006020828403121561166c5761166b611303565b5b600061167a84828501611641565b91505092915050565b600082825260208201905092915050565b7f43574c3a2073656c6c6572206e6f742073657400000000000000000000000000600082015250565b60006116ca601383611683565b91506116d582611694565b602082019050919050565b600060208201905081810360008301526116f9816116bd565b9050919050565b7f43574c3a207573646320746f6b656e206e6f7420736574000000000000000000600082015250565b6000611736601783611683565b915061174182611700565b602082019050919050565b6000602082019050818103600083015261176581611729565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117a6826112c5565b91506117b1836112c5565b92508282019050808211156117c9576117c861176c565b5b92915050565b6000815190506117de8161136b565b92915050565b6000602082840312156117fa576117f9611303565b5b6000611808848285016117cf565b91505092915050565b7f433a206e6f7420656e6f75676820555344000000000000000000000000000000600082015250565b6000611847601183611683565b915061185282611811565b602082019050919050565b600060208201905081810360008301526118768161183a565b9050919050565b6000819050919050565b6000819050919050565b60006118ac6118a76118a28461187d565b611887565b61148e565b9050919050565b6118bc81611891565b82525050565b600060e0820190506118d7600083018a61147f565b6118e4602083018961147f565b6118f1604083018861147f565b6118fe60608301876112cf565b61190b60808301866112cf565b61191860a08301856112cf565b61192560c08301846118b3565b98975050505050505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61197f82611936565b810181811067ffffffffffffffff8211171561199e5761199d611947565b5b80604052505050565b60006119b16112f9565b90506119bd8282611976565b919050565b6119cb8161148e565b81146119d657600080fd5b50565b6000815190506119e8816119c2565b92915050565b60006101408284031215611a0557611a04611931565b5b611a106101406119a7565b90506000611a2084828501611641565b6000830152506020611a3484828501611641565b6020830152506040611a4884828501611641565b6040830152506060611a5c848285016117cf565b6060830152506080611a70848285016117cf565b60808301525060a0611a84848285016117cf565b60a08301525060c0611a98848285016119d9565b60c08301525060e0611aac848285016119d9565b60e083015250610100611ac1848285016117cf565b61010083015250610120611ad7848285016117cf565b6101208301525092915050565b60006101408284031215611afb57611afa611303565b5b6000611b09848285016119ee565b91505092915050565b7f43574c3a206f6e6c792073656c6c657200000000000000000000000000000000600082015250565b6000611b48601083611683565b9150611b5382611b12565b602082019050919050565b60006020820190508181036000830152611b7781611b3b565b9050919050565b6000604082019050611b93600083018561147f565b611ba060208301846112cf565b9392505050565b6000819050919050565b6000611bcc611bc7611bc284611ba7565b611887565b61148e565b9050919050565b611bdc81611bb1565b82525050565b600060e082019050611bf7600083018a61147f565b611c04602083018961147f565b611c11604083018861147f565b611c1e60608301876112cf565b611c2b60808301866112cf565b611c3860a08301856112cf565b611c4560c0830184611bd3565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f43574c3a206f6e6c792062757965720000000000000000000000000000000000600082015250565b6000611cb6600f83611683565b9150611cc182611c80565b602082019050919050565b60006020820190508181036000830152611ce581611ca9565b9050919050565b7f43574c3a206e6f206f70656e206f726465727300000000000000000000000000600082015250565b6000611d22601383611683565b9150611d2d82611cec565b602082019050919050565b60006020820190508181036000830152611d5181611d15565b9050919050565b7f43574c3a20696e76616c6964206f726465720000000000000000000000000000600082015250565b6000611d8e601283611683565b9150611d9982611d58565b602082019050919050565b60006020820190508181036000830152611dbd81611d81565b9050919050565b600080fd5b600067ffffffffffffffff821115611de457611de3611947565b5b602082029050602081019050919050565b600080fd5b6000611e0d611e0884611dc9565b6119a7565b90508083825260208201905060208402830185811115611e3057611e2f611df5565b5b835b81811015611e595780611e4588826117cf565b845260208401935050602081019050611e32565b5050509392505050565b600082601f830112611e7857611e77611dc4565b5b8151611e88848260208601611dfa565b91505092915050565b600060208284031215611ea757611ea6611303565b5b600082015167ffffffffffffffff811115611ec557611ec4611308565b5b611ed184828501611e63565b91505092915050565b7f43574c3a2073656c6620666f7262696464656e00000000000000000000000000600082015250565b6000611f10601383611683565b9150611f1b82611eda565b602082019050919050565b60006020820190508181036000830152611f3f81611f03565b9050919050565b7f43574c3a20696e76616c696420746f2061646472657373000000000000000000600082015250565b6000611f7c601783611683565b9150611f8782611f46565b602082019050919050565b60006020820190508181036000830152611fab81611f6f565b9050919050565b7f43574c3a20616d6f756e742063616e6e6f7420657175616c2030000000000000600082015250565b6000611fe8601a83611683565b9150611ff382611fb2565b602082019050919050565b6000602082019050818103600083015261201781611fdb565b9050919050565b7f43574c3a20696e73756666696369656e742066756e6473000000000000000000600082015250565b6000612054601783611683565b915061205f8261201e565b602082019050919050565b6000602082019050818103600083015261208381612047565b9050919050565b60008115159050919050565b61209f8161208a565b81146120aa57600080fd5b50565b6000815190506120bc81612096565b92915050565b6000602082840312156120d8576120d7611303565b5b60006120e6848285016120ad565b91505092915050565b60006120fa826112c5565b9150612105836112c5565b925082820390508181111561211d5761211c61176c565b5b92915050565b600061212e826112c5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121605761215f61176c565b5b60018201905091905056fea2646970667358221220cc0d6ae1007d4d6762cfc96d9063e1f6eddfc4e55d4ab38e0c055e0d4dd1f05364736f6c63430008100033",
}

// CustodianWalletLogicABI is the input ABI used to generate the binding from.
// Deprecated: Use CustodianWalletLogicMetaData.ABI instead.
var CustodianWalletLogicABI = CustodianWalletLogicMetaData.ABI

// CustodianWalletLogicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustodianWalletLogicMetaData.Bin instead.
var CustodianWalletLogicBin = CustodianWalletLogicMetaData.Bin

// DeployCustodianWalletLogic deploys a new Ethereum contract, binding an instance of CustodianWalletLogic to it.
func DeployCustodianWalletLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CustodianWalletLogic, error) {
	parsed, err := CustodianWalletLogicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustodianWalletLogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustodianWalletLogic{CustodianWalletLogicCaller: CustodianWalletLogicCaller{contract: contract}, CustodianWalletLogicTransactor: CustodianWalletLogicTransactor{contract: contract}, CustodianWalletLogicFilterer: CustodianWalletLogicFilterer{contract: contract}}, nil
}

// CustodianWalletLogic is an auto generated Go binding around an Ethereum contract.
type CustodianWalletLogic struct {
	CustodianWalletLogicCaller     // Read-only binding to the contract
	CustodianWalletLogicTransactor // Write-only binding to the contract
	CustodianWalletLogicFilterer   // Log filterer for contract events
}

// CustodianWalletLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodianWalletLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodianWalletLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodianWalletLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodianWalletLogicSession struct {
	Contract     *CustodianWalletLogic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CustodianWalletLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodianWalletLogicCallerSession struct {
	Contract *CustodianWalletLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CustodianWalletLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodianWalletLogicTransactorSession struct {
	Contract     *CustodianWalletLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CustodianWalletLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodianWalletLogicRaw struct {
	Contract *CustodianWalletLogic // Generic contract binding to access the raw methods on
}

// CustodianWalletLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodianWalletLogicCallerRaw struct {
	Contract *CustodianWalletLogicCaller // Generic read-only contract binding to access the raw methods on
}

// CustodianWalletLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodianWalletLogicTransactorRaw struct {
	Contract *CustodianWalletLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustodianWalletLogic creates a new instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogic(address common.Address, backend bind.ContractBackend) (*CustodianWalletLogic, error) {
	contract, err := bindCustodianWalletLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogic{CustodianWalletLogicCaller: CustodianWalletLogicCaller{contract: contract}, CustodianWalletLogicTransactor: CustodianWalletLogicTransactor{contract: contract}, CustodianWalletLogicFilterer: CustodianWalletLogicFilterer{contract: contract}}, nil
}

// NewCustodianWalletLogicCaller creates a new read-only instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicCaller(address common.Address, caller bind.ContractCaller) (*CustodianWalletLogicCaller, error) {
	contract, err := bindCustodianWalletLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicCaller{contract: contract}, nil
}

// NewCustodianWalletLogicTransactor creates a new write-only instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodianWalletLogicTransactor, error) {
	contract, err := bindCustodianWalletLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicTransactor{contract: contract}, nil
}

// NewCustodianWalletLogicFilterer creates a new log filterer instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodianWalletLogicFilterer, error) {
	contract, err := bindCustodianWalletLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicFilterer{contract: contract}, nil
}

// bindCustodianWalletLogic binds a generic wrapper to an already deployed contract.
func bindCustodianWalletLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianWalletLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletLogic *CustodianWalletLogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletLogic *CustodianWalletLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletLogic *CustodianWalletLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.contract.Transact(opts, method, params...)
}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) AvailBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "availBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) AvailBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.AvailBalance(&_CustodianWalletLogic.CallOpts)
}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) AvailBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.AvailBalance(&_CustodianWalletLogic.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicSession) Factory() (common.Address, error) {
	return _CustodianWalletLogic.Contract.Factory(&_CustodianWalletLogic.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) Factory() (common.Address, error) {
	return _CustodianWalletLogic.Contract.Factory(&_CustodianWalletLogic.CallOpts)
}

// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicCaller) GetOpenOrders(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "getOpenOrders")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) Ochestrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "ochestrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}


// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicSession) GetOpenOrders() ([]*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetOpenOrders(&_CustodianWalletLogic.CallOpts)
}

// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) GetOpenOrders() ([]*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetOpenOrders(&_CustodianWalletLogic.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) GetTotalBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetTotalBalance(&_CustodianWalletLogic.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) GetTotalBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetTotalBalance(&_CustodianWalletLogic.CallOpts)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) OpenOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "openOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletLogic.Contract.OpenOrders(&_CustodianWalletLogic.CallOpts, arg0, arg1)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletLogic.Contract.OpenOrders(&_CustodianWalletLogic.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _CustodianWalletLogic.contract.Call(opts, &out, "orders", arg0)

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
func (_CustodianWalletLogic *CustodianWalletLogicSession) Orders(arg0 *big.Int) (struct {
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
	return _CustodianWalletLogic.Contract.Orders(&_CustodianWalletLogic.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) Orders(arg0 *big.Int) (struct {
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
	return _CustodianWalletLogic.Contract.Orders(&_CustodianWalletLogic.CallOpts, arg0)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) ApproveOrder(opts *bind.TransactOpts, _openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "approveOrder", _openOrderIndex)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) ApproveOrder(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ApproveOrder(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) ApproveOrder(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ApproveOrder(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) ConsentOrderRejected(opts *bind.TransactOpts, _openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "consentOrderRejected", _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) ConsentOrderRejected(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ConsentOrderRejected(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) ConsentOrderRejected(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ConsentOrderRejected(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) NewBuyOrder(opts *bind.TransactOpts, _seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "newBuyOrder", _seller, _receiver, _amount, _rate, _fee)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) NewBuyOrder(_seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewBuyOrder(&_CustodianWalletLogic.TransactOpts, _seller, _receiver, _amount, _rate, _fee)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) NewBuyOrder(_seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewBuyOrder(&_CustodianWalletLogic.TransactOpts, _seller, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) NewSellOrder(opts *bind.TransactOpts, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "newSellOrder", _buyer, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) NewSellOrder(_buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewSellOrder(&_CustodianWalletLogic.TransactOpts, _buyer, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) NewSellOrder(_buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewSellOrder(&_CustodianWalletLogic.TransactOpts, _buyer, _receiver, _amount, _rate, _fee)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) RejectOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "rejectOrder", _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) RejectOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.RejectOrder(&_CustodianWalletLogic.TransactOpts, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) RejectOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.RejectOrder(&_CustodianWalletLogic.TransactOpts, _orderId)
}

// CustodianWalletLogicApproveRejectedOrderIterator is returned from FilterApproveRejectedOrder and is used to iterate over the raw logs and unpacked data for ApproveRejectedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicApproveRejectedOrderIterator struct {
	Event *CustodianWalletLogicApproveRejectedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicApproveRejectedOrder)
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
		it.Event = new(CustodianWalletLogicApproveRejectedOrder)
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
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicApproveRejectedOrder represents a ApproveRejectedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicApproveRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproveRejectedOrder is a free log retrieval operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterApproveRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletLogicApproveRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicApproveRejectedOrderIterator{contract: _CustodianWalletLogic.contract, event: "ApproveRejectedOrder", logs: logs, sub: sub}, nil
}

// WatchApproveRejectedOrder is a free log subscription operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchApproveRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicApproveRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicApproveRejectedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseApproveRejectedOrder(log types.Log) (*CustodianWalletLogicApproveRejectedOrder, error) {
	event := new(CustodianWalletLogicApproveRejectedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicClosedOrderIterator is returned from FilterClosedOrder and is used to iterate over the raw logs and unpacked data for ClosedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicClosedOrderIterator struct {
	Event *CustodianWalletLogicClosedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletLogicClosedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicClosedOrder)
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
		it.Event = new(CustodianWalletLogicClosedOrder)
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
func (it *CustodianWalletLogicClosedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicClosedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicClosedOrder represents a ClosedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicClosedOrder struct {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterClosedOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletLogicClosedOrderIterator, error) {

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

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicClosedOrderIterator{contract: _CustodianWalletLogic.contract, event: "ClosedOrder", logs: logs, sub: sub}, nil
}

// WatchClosedOrder is a free log subscription operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchClosedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicClosedOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicClosedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseClosedOrder(log types.Log) (*CustodianWalletLogicClosedOrder, error) {
	event := new(CustodianWalletLogicClosedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicOpenOrderIterator is returned from FilterOpenOrder and is used to iterate over the raw logs and unpacked data for OpenOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOpenOrderIterator struct {
	Event *CustodianWalletLogicOpenOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletLogicOpenOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicOpenOrder)
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
		it.Event = new(CustodianWalletLogicOpenOrder)
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
func (it *CustodianWalletLogicOpenOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicOpenOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicOpenOrder represents a OpenOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOpenOrder struct {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterOpenOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletLogicOpenOrderIterator, error) {

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

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicOpenOrderIterator{contract: _CustodianWalletLogic.contract, event: "OpenOrder", logs: logs, sub: sub}, nil
}

// WatchOpenOrder is a free log subscription operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchOpenOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicOpenOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicOpenOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "OpenOrder", log); err != nil {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseOpenOrder(log types.Log) (*CustodianWalletLogicOpenOrder, error) {
	event := new(CustodianWalletLogicOpenOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "OpenOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOrderFulfilledIterator struct {
	Event *CustodianWalletLogicOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *CustodianWalletLogicOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicOrderFulfilled)
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
		it.Event = new(CustodianWalletLogicOrderFulfilled)
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
func (it *CustodianWalletLogicOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicOrderFulfilled represents a OrderFulfilled event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOrderFulfilled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*CustodianWalletLogicOrderFulfilledIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicOrderFulfilledIterator{contract: _CustodianWalletLogic.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicOrderFulfilled)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseOrderFulfilled(log types.Log) (*CustodianWalletLogicOrderFulfilled, error) {
	event := new(CustodianWalletLogicOrderFulfilled)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicRejectedOrderIterator is returned from FilterRejectedOrder and is used to iterate over the raw logs and unpacked data for RejectedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicRejectedOrderIterator struct {
	Event *CustodianWalletLogicRejectedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletLogicRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicRejectedOrder)
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
		it.Event = new(CustodianWalletLogicRejectedOrder)
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
func (it *CustodianWalletLogicRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicRejectedOrder represents a RejectedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRejectedOrder is a free log retrieval operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletLogicRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicRejectedOrderIterator{contract: _CustodianWalletLogic.contract, event: "RejectedOrder", logs: logs, sub: sub}, nil
}

// WatchRejectedOrder is a free log subscription operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicRejectedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
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
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseRejectedOrder(log types.Log) (*CustodianWalletLogicRejectedOrder, error) {
	event := new(CustodianWalletLogicRejectedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
