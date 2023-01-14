// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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

// SwapParameters is an auto generated low-level Go binding around an user-defined struct.
type SwapParameters struct {
	Pairs                 []common.Address
	Reserves              [][]*big.Int
	PairTokens            [][]common.Address
	Path                  []common.Address
	AmountsOut            []*big.Int
	RevertOnReserveChange bool
	GasToken              common.Address
	UseGasToken           bool
}

// SwapExecutorV2MetaData contains all meta data concerning the SwapExecutorV2 contract.
var SwapExecutorV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"Pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"Reserves\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"PairTokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[]\",\"name\":\"Path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"AmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"RevertOnReserveChange\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"GasToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"UseGasToken\",\"type\":\"bool\"}],\"internalType\":\"structSwapParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"executeSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapExecutorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapExecutorV2MetaData.ABI instead.
var SwapExecutorV2ABI = SwapExecutorV2MetaData.ABI

// SwapExecutorV2 is an auto generated Go binding around an Ethereum contract.
type SwapExecutorV2 struct {
	SwapExecutorV2Caller     // Read-only binding to the contract
	SwapExecutorV2Transactor // Write-only binding to the contract
	SwapExecutorV2Filterer   // Log filterer for contract events
}

// SwapExecutorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapExecutorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapExecutorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapExecutorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapExecutorV2Session struct {
	Contract     *SwapExecutorV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapExecutorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapExecutorV2CallerSession struct {
	Contract *SwapExecutorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SwapExecutorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapExecutorV2TransactorSession struct {
	Contract     *SwapExecutorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SwapExecutorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapExecutorV2Raw struct {
	Contract *SwapExecutorV2 // Generic contract binding to access the raw methods on
}

// SwapExecutorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapExecutorV2CallerRaw struct {
	Contract *SwapExecutorV2Caller // Generic read-only contract binding to access the raw methods on
}

// SwapExecutorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapExecutorV2TransactorRaw struct {
	Contract *SwapExecutorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapExecutorV2 creates a new instance of SwapExecutorV2, bound to a specific deployed contract.
func NewSwapExecutorV2(address common.Address, backend bind.ContractBackend) (*SwapExecutorV2, error) {
	contract, err := bindSwapExecutorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV2{SwapExecutorV2Caller: SwapExecutorV2Caller{contract: contract}, SwapExecutorV2Transactor: SwapExecutorV2Transactor{contract: contract}, SwapExecutorV2Filterer: SwapExecutorV2Filterer{contract: contract}}, nil
}

// NewSwapExecutorV2Caller creates a new read-only instance of SwapExecutorV2, bound to a specific deployed contract.
func NewSwapExecutorV2Caller(address common.Address, caller bind.ContractCaller) (*SwapExecutorV2Caller, error) {
	contract, err := bindSwapExecutorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV2Caller{contract: contract}, nil
}

// NewSwapExecutorV2Transactor creates a new write-only instance of SwapExecutorV2, bound to a specific deployed contract.
func NewSwapExecutorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapExecutorV2Transactor, error) {
	contract, err := bindSwapExecutorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV2Transactor{contract: contract}, nil
}

// NewSwapExecutorV2Filterer creates a new log filterer instance of SwapExecutorV2, bound to a specific deployed contract.
func NewSwapExecutorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapExecutorV2Filterer, error) {
	contract, err := bindSwapExecutorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV2Filterer{contract: contract}, nil
}

// bindSwapExecutorV2 binds a generic wrapper to an already deployed contract.
func bindSwapExecutorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapExecutorV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutorV2 *SwapExecutorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutorV2.Contract.SwapExecutorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutorV2 *SwapExecutorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.SwapExecutorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutorV2 *SwapExecutorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.SwapExecutorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutorV2 *SwapExecutorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutorV2 *SwapExecutorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutorV2 *SwapExecutorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.contract.Transact(opts, method, params...)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0xcae89ff8.
//
// Solidity: function executeSwap((address[],uint256[][],address[][],address[],uint256[],bool,address,bool) params) returns()
func (_SwapExecutorV2 *SwapExecutorV2Transactor) ExecuteSwap(opts *bind.TransactOpts, params SwapParameters) (*types.Transaction, error) {
	return _SwapExecutorV2.contract.Transact(opts, "executeSwap", params)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0xcae89ff8.
//
// Solidity: function executeSwap((address[],uint256[][],address[][],address[],uint256[],bool,address,bool) params) returns()
func (_SwapExecutorV2 *SwapExecutorV2Session) ExecuteSwap(params SwapParameters) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.ExecuteSwap(&_SwapExecutorV2.TransactOpts, params)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0xcae89ff8.
//
// Solidity: function executeSwap((address[],uint256[][],address[][],address[],uint256[],bool,address,bool) params) returns()
func (_SwapExecutorV2 *SwapExecutorV2TransactorSession) ExecuteSwap(params SwapParameters) (*types.Transaction, error) {
	return _SwapExecutorV2.Contract.ExecuteSwap(&_SwapExecutorV2.TransactOpts, params)
}
