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

// CustodianWalletProxyMetaData contains all meta data concerning the CustodianWalletProxy contract.
var CustodianWalletProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ochestrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"ApproveRejectedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"ClosedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"OpenOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"RejectedOrder\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ochestrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161091938038061091983398181016040528101906100329190610146565b8173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610199565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610113826100e8565b9050919050565b61012381610108565b811461012e57600080fd5b50565b6000815190506101408161011a565b92915050565b60008060006060848603121561015f5761015e6100e3565b5b600061016d86828701610131565b935050602061017e86828701610131565b925050604061018f86828701610131565b9150509250925092565b60805160a05161074f6101ca60003960008181605701526102d401526000818160e701526103fc015261074f6000f3fe60806040526004361061004e5760003560e01c80631e7cea801461019d5780639c811d43146101da578063a85c38ef14610205578063c45a01551461024b578063d7dfa0dd1461027657610055565b3661005557005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100da9061047b565b60405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000009050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610177576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016e906104e7565b60405180910390fd5b3660008037600080366000845af43d6000803e8060008114610198573d6000f35b3d6000fd5b3480156101a957600080fd5b506101c460048036038101906101bf91906105a0565b6102a1565b6040516101d191906105ef565b60405180910390f35b3480156101e657600080fd5b506101ef6102d2565b6040516101fc9190610619565b60405180910390f35b34801561021157600080fd5b5061022c60048036038101906102279190610634565b6102f6565b6040516102429a9998979695949392919061067d565b60405180910390f35b34801561025757600080fd5b506102606103d4565b60405161026d9190610619565b60405180910390f35b34801561028257600080fd5b5061028b6103fa565b6040516102989190610619565b60405180910390f35b600160205281600052604060002081815481106102bd57600080fd5b90600052602060002001600091509150505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000818154811061030657600080fd5b90600052602060002090600902016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900460ff1690806007015490806008015490508a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600082825260208201905092915050565b7f57503a206465706c6f796572206f6e6c79000000000000000000000000000000600082015250565b600061046560118361041e565b91506104708261042f565b602082019050919050565b6000602082019050818103600083015261049481610458565b9050919050565b7f4c6f67696320636f6e7472616374206e6f742073657400000000000000000000600082015250565b60006104d160168361041e565b91506104dc8261049b565b602082019050919050565b60006020820190508181036000830152610500816104c4565b9050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105378261050c565b9050919050565b6105478161052c565b811461055257600080fd5b50565b6000813590506105648161053e565b92915050565b6000819050919050565b61057d8161056a565b811461058857600080fd5b50565b60008135905061059a81610574565b92915050565b600080604083850312156105b7576105b6610507565b5b60006105c585828601610555565b92505060206105d68582860161058b565b9150509250929050565b6105e98161056a565b82525050565b600060208201905061060460008301846105e0565b92915050565b6106138161052c565b82525050565b600060208201905061062e600083018461060a565b92915050565b60006020828403121561064a57610649610507565b5b60006106588482850161058b565b91505092915050565b600060ff82169050919050565b61067781610661565b82525050565b600061014082019050610693600083018d61060a565b6106a0602083018c61060a565b6106ad604083018b61060a565b6106ba606083018a6105e0565b6106c760808301896105e0565b6106d460a08301886105e0565b6106e160c083018761066e565b6106ee60e083018661066e565b6106fc6101008301856105e0565b61070a6101208301846105e0565b9b9a505050505050505050505056fea26469706673582212206310d4c702078f951b16bc3595bdbf9bd0e4f5e5134f883d6f261bc5db6eebf164736f6c63430008100033",
}

// CustodianWalletProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use CustodianWalletProxyMetaData.ABI instead.
var CustodianWalletProxyABI = CustodianWalletProxyMetaData.ABI

// CustodianWalletProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustodianWalletProxyMetaData.Bin instead.
var CustodianWalletProxyBin = CustodianWalletProxyMetaData.Bin

// DeployCustodianWalletProxy deploys a new Ethereum contract, binding an instance of CustodianWalletProxy to it.
func DeployCustodianWalletProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _ochestrator common.Address, _factory common.Address) (common.Address, *types.Transaction, *CustodianWalletProxy, error) {
	parsed, err := CustodianWalletProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustodianWalletProxyBin), backend, _logic, _ochestrator, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustodianWalletProxy{CustodianWalletProxyCaller: CustodianWalletProxyCaller{contract: contract}, CustodianWalletProxyTransactor: CustodianWalletProxyTransactor{contract: contract}, CustodianWalletProxyFilterer: CustodianWalletProxyFilterer{contract: contract}}, nil
}

// CustodianWalletProxy is an auto generated Go binding around an Ethereum contract.
type CustodianWalletProxy struct {
	CustodianWalletProxyCaller     // Read-only binding to the contract
	CustodianWalletProxyTransactor // Write-only binding to the contract
	CustodianWalletProxyFilterer   // Log filterer for contract events
}

// CustodianWalletProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodianWalletProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodianWalletProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodianWalletProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodianWalletProxySession struct {
	Contract     *CustodianWalletProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CustodianWalletProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodianWalletProxyCallerSession struct {
	Contract *CustodianWalletProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CustodianWalletProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodianWalletProxyTransactorSession struct {
	Contract     *CustodianWalletProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CustodianWalletProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodianWalletProxyRaw struct {
	Contract *CustodianWalletProxy // Generic contract binding to access the raw methods on
}

// CustodianWalletProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodianWalletProxyCallerRaw struct {
	Contract *CustodianWalletProxyCaller // Generic read-only contract binding to access the raw methods on
}

// CustodianWalletProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodianWalletProxyTransactorRaw struct {
	Contract *CustodianWalletProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustodianWalletProxy creates a new instance of CustodianWalletProxy, bound to a specific deployed contract.
func NewCustodianWalletProxy(address common.Address, backend bind.ContractBackend) (*CustodianWalletProxy, error) {
	contract, err := bindCustodianWalletProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxy{CustodianWalletProxyCaller: CustodianWalletProxyCaller{contract: contract}, CustodianWalletProxyTransactor: CustodianWalletProxyTransactor{contract: contract}, CustodianWalletProxyFilterer: CustodianWalletProxyFilterer{contract: contract}}, nil
}

// NewCustodianWalletProxyCaller creates a new read-only instance of CustodianWalletProxy, bound to a specific deployed contract.
func NewCustodianWalletProxyCaller(address common.Address, caller bind.ContractCaller) (*CustodianWalletProxyCaller, error) {
	contract, err := bindCustodianWalletProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyCaller{contract: contract}, nil
}

// NewCustodianWalletProxyTransactor creates a new write-only instance of CustodianWalletProxy, bound to a specific deployed contract.
func NewCustodianWalletProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodianWalletProxyTransactor, error) {
	contract, err := bindCustodianWalletProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyTransactor{contract: contract}, nil
}

// NewCustodianWalletProxyFilterer creates a new log filterer instance of CustodianWalletProxy, bound to a specific deployed contract.
func NewCustodianWalletProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodianWalletProxyFilterer, error) {
	contract, err := bindCustodianWalletProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyFilterer{contract: contract}, nil
}

// bindCustodianWalletProxy binds a generic wrapper to an already deployed contract.
func bindCustodianWalletProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianWalletProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletProxy *CustodianWalletProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletProxy.Contract.CustodianWalletProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletProxy *CustodianWalletProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.CustodianWalletProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletProxy *CustodianWalletProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.CustodianWalletProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletProxy *CustodianWalletProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletProxy *CustodianWalletProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletProxy *CustodianWalletProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletProxy.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxySession) Factory() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Factory(&_CustodianWalletProxy.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCallerSession) Factory() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Factory(&_CustodianWalletProxy.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletProxy.contract.Call(opts, &out, "logic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxySession) Logic() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Logic(&_CustodianWalletProxy.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCallerSession) Logic() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Logic(&_CustodianWalletProxy.CallOpts)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCaller) Ochestrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletProxy.contract.Call(opts, &out, "ochestrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxySession) Ochestrator() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Ochestrator(&_CustodianWalletProxy.CallOpts)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_CustodianWalletProxy *CustodianWalletProxyCallerSession) Ochestrator() (common.Address, error) {
	return _CustodianWalletProxy.Contract.Ochestrator(&_CustodianWalletProxy.CallOpts)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletProxy *CustodianWalletProxyCaller) OpenOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletProxy.contract.Call(opts, &out, "openOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletProxy *CustodianWalletProxySession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletProxy.Contract.OpenOrders(&_CustodianWalletProxy.CallOpts, arg0, arg1)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletProxy *CustodianWalletProxyCallerSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletProxy.Contract.OpenOrders(&_CustodianWalletProxy.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletProxy *CustodianWalletProxyCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _CustodianWalletProxy.contract.Call(opts, &out, "orders", arg0)

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
func (_CustodianWalletProxy *CustodianWalletProxySession) Orders(arg0 *big.Int) (struct {
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
	return _CustodianWalletProxy.Contract.Orders(&_CustodianWalletProxy.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletProxy *CustodianWalletProxyCallerSession) Orders(arg0 *big.Int) (struct {
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
	return _CustodianWalletProxy.Contract.Orders(&_CustodianWalletProxy.CallOpts, arg0)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CustodianWalletProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.Fallback(&_CustodianWalletProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.Fallback(&_CustodianWalletProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxySession) Receive() (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.Receive(&_CustodianWalletProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CustodianWalletProxy *CustodianWalletProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _CustodianWalletProxy.Contract.Receive(&_CustodianWalletProxy.TransactOpts)
}

// CustodianWalletProxyApproveRejectedOrderIterator is returned from FilterApproveRejectedOrder and is used to iterate over the raw logs and unpacked data for ApproveRejectedOrder events raised by the CustodianWalletProxy contract.
type CustodianWalletProxyApproveRejectedOrderIterator struct {
	Event *CustodianWalletProxyApproveRejectedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletProxyApproveRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletProxyApproveRejectedOrder)
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
		it.Event = new(CustodianWalletProxyApproveRejectedOrder)
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
func (it *CustodianWalletProxyApproveRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletProxyApproveRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletProxyApproveRejectedOrder represents a ApproveRejectedOrder event raised by the CustodianWalletProxy contract.
type CustodianWalletProxyApproveRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproveRejectedOrder is a free log retrieval operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) FilterApproveRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletProxyApproveRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletProxy.contract.FilterLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyApproveRejectedOrderIterator{contract: _CustodianWalletProxy.contract, event: "ApproveRejectedOrder", logs: logs, sub: sub}, nil
}

// WatchApproveRejectedOrder is a free log subscription operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) WatchApproveRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletProxyApproveRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletProxy.contract.WatchLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletProxyApproveRejectedOrder)
				if err := _CustodianWalletProxy.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) ParseApproveRejectedOrder(log types.Log) (*CustodianWalletProxyApproveRejectedOrder, error) {
	event := new(CustodianWalletProxyApproveRejectedOrder)
	if err := _CustodianWalletProxy.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletProxyClosedOrderIterator is returned from FilterClosedOrder and is used to iterate over the raw logs and unpacked data for ClosedOrder events raised by the CustodianWalletProxy contract.
type CustodianWalletProxyClosedOrderIterator struct {
	Event *CustodianWalletProxyClosedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletProxyClosedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletProxyClosedOrder)
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
		it.Event = new(CustodianWalletProxyClosedOrder)
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
func (it *CustodianWalletProxyClosedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletProxyClosedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletProxyClosedOrder represents a ClosedOrder event raised by the CustodianWalletProxy contract.
type CustodianWalletProxyClosedOrder struct {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) FilterClosedOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletProxyClosedOrderIterator, error) {

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

	logs, sub, err := _CustodianWalletProxy.contract.FilterLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyClosedOrderIterator{contract: _CustodianWalletProxy.contract, event: "ClosedOrder", logs: logs, sub: sub}, nil
}

// WatchClosedOrder is a free log subscription operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) WatchClosedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletProxyClosedOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CustodianWalletProxy.contract.WatchLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletProxyClosedOrder)
				if err := _CustodianWalletProxy.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) ParseClosedOrder(log types.Log) (*CustodianWalletProxyClosedOrder, error) {
	event := new(CustodianWalletProxyClosedOrder)
	if err := _CustodianWalletProxy.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletProxyOpenOrderIterator is returned from FilterOpenOrder and is used to iterate over the raw logs and unpacked data for OpenOrder events raised by the CustodianWalletProxy contract.
type CustodianWalletProxyOpenOrderIterator struct {
	Event *CustodianWalletProxyOpenOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletProxyOpenOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletProxyOpenOrder)
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
		it.Event = new(CustodianWalletProxyOpenOrder)
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
func (it *CustodianWalletProxyOpenOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletProxyOpenOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletProxyOpenOrder represents a OpenOrder event raised by the CustodianWalletProxy contract.
type CustodianWalletProxyOpenOrder struct {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) FilterOpenOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletProxyOpenOrderIterator, error) {

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

	logs, sub, err := _CustodianWalletProxy.contract.FilterLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyOpenOrderIterator{contract: _CustodianWalletProxy.contract, event: "OpenOrder", logs: logs, sub: sub}, nil
}

// WatchOpenOrder is a free log subscription operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) WatchOpenOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletProxyOpenOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CustodianWalletProxy.contract.WatchLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletProxyOpenOrder)
				if err := _CustodianWalletProxy.contract.UnpackLog(event, "OpenOrder", log); err != nil {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) ParseOpenOrder(log types.Log) (*CustodianWalletProxyOpenOrder, error) {
	event := new(CustodianWalletProxyOpenOrder)
	if err := _CustodianWalletProxy.contract.UnpackLog(event, "OpenOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletProxyOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the CustodianWalletProxy contract.
type CustodianWalletProxyOrderFulfilledIterator struct {
	Event *CustodianWalletProxyOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *CustodianWalletProxyOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletProxyOrderFulfilled)
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
		it.Event = new(CustodianWalletProxyOrderFulfilled)
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
func (it *CustodianWalletProxyOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletProxyOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletProxyOrderFulfilled represents a OrderFulfilled event raised by the CustodianWalletProxy contract.
type CustodianWalletProxyOrderFulfilled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*CustodianWalletProxyOrderFulfilledIterator, error) {

	logs, sub, err := _CustodianWalletProxy.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyOrderFulfilledIterator{contract: _CustodianWalletProxy.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *CustodianWalletProxyOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletProxy.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletProxyOrderFulfilled)
				if err := _CustodianWalletProxy.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) ParseOrderFulfilled(log types.Log) (*CustodianWalletProxyOrderFulfilled, error) {
	event := new(CustodianWalletProxyOrderFulfilled)
	if err := _CustodianWalletProxy.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletProxyRejectedOrderIterator is returned from FilterRejectedOrder and is used to iterate over the raw logs and unpacked data for RejectedOrder events raised by the CustodianWalletProxy contract.
type CustodianWalletProxyRejectedOrderIterator struct {
	Event *CustodianWalletProxyRejectedOrder // Event containing the contract specifics and raw log

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
func (it *CustodianWalletProxyRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletProxyRejectedOrder)
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
		it.Event = new(CustodianWalletProxyRejectedOrder)
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
func (it *CustodianWalletProxyRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletProxyRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletProxyRejectedOrder represents a RejectedOrder event raised by the CustodianWalletProxy contract.
type CustodianWalletProxyRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRejectedOrder is a free log retrieval operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) FilterRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletProxyRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletProxy.contract.FilterLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletProxyRejectedOrderIterator{contract: _CustodianWalletProxy.contract, event: "RejectedOrder", logs: logs, sub: sub}, nil
}

// WatchRejectedOrder is a free log subscription operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) WatchRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletProxyRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletProxy.contract.WatchLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletProxyRejectedOrder)
				if err := _CustodianWalletProxy.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
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
func (_CustodianWalletProxy *CustodianWalletProxyFilterer) ParseRejectedOrder(log types.Log) (*CustodianWalletProxyRejectedOrder, error) {
	event := new(CustodianWalletProxyRejectedOrder)
	if err := _CustodianWalletProxy.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
