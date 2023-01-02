// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

import (
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
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

// SwapExecutorV1MetaData contains all meta data concerning the SwapExecutorV1 contract.
var SwapExecutorV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"pairTokens\",\"type\":\"address[][]\"}],\"name\":\"executeSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapExecutorV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapExecutorV1MetaData.ABI instead.
var SwapExecutorV1ABI = SwapExecutorV1MetaData.ABI

// SwapExecutorV1 is an auto generated Go binding around an Ethereum contract.
type SwapExecutorV1 struct {
	SwapExecutorV1Caller     // Read-only binding to the contract
	SwapExecutorV1Transactor // Write-only binding to the contract
	SwapExecutorV1Filterer   // Log filterer for contract events
}

// SwapExecutorV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapExecutorV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapExecutorV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapExecutorV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExecutorV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapExecutorV1Session struct {
	Contract     *SwapExecutorV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapExecutorV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapExecutorV1CallerSession struct {
	Contract *SwapExecutorV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SwapExecutorV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapExecutorV1TransactorSession struct {
	Contract     *SwapExecutorV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SwapExecutorV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapExecutorV1Raw struct {
	Contract *SwapExecutorV1 // Generic contract binding to access the raw methods on
}

// SwapExecutorV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapExecutorV1CallerRaw struct {
	Contract *SwapExecutorV1Caller // Generic read-only contract binding to access the raw methods on
}

// SwapExecutorV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapExecutorV1TransactorRaw struct {
	Contract *SwapExecutorV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapExecutorV1 creates a new instance of SwapExecutorV1, bound to a specific deployed contract.
func NewSwapExecutorV1(address common.Address, backend bind.ContractBackend) (*SwapExecutorV1, error) {
	contract, err := bindSwapExecutorV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV1{SwapExecutorV1Caller: SwapExecutorV1Caller{contract: contract}, SwapExecutorV1Transactor: SwapExecutorV1Transactor{contract: contract}, SwapExecutorV1Filterer: SwapExecutorV1Filterer{contract: contract}}, nil
}

// NewSwapExecutorV1Caller creates a new read-only instance of SwapExecutorV1, bound to a specific deployed contract.
func NewSwapExecutorV1Caller(address common.Address, caller bind.ContractCaller) (*SwapExecutorV1Caller, error) {
	contract, err := bindSwapExecutorV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV1Caller{contract: contract}, nil
}

// NewSwapExecutorV1Transactor creates a new write-only instance of SwapExecutorV1, bound to a specific deployed contract.
func NewSwapExecutorV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapExecutorV1Transactor, error) {
	contract, err := bindSwapExecutorV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV1Transactor{contract: contract}, nil
}

// NewSwapExecutorV1Filterer creates a new log filterer instance of SwapExecutorV1, bound to a specific deployed contract.
func NewSwapExecutorV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapExecutorV1Filterer, error) {
	contract, err := bindSwapExecutorV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapExecutorV1Filterer{contract: contract}, nil
}

// bindSwapExecutorV1 binds a generic wrapper to an already deployed contract.
func bindSwapExecutorV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapExecutorV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutorV1 *SwapExecutorV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutorV1.Contract.SwapExecutorV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutorV1 *SwapExecutorV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.SwapExecutorV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutorV1 *SwapExecutorV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.SwapExecutorV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExecutorV1 *SwapExecutorV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapExecutorV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExecutorV1 *SwapExecutorV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExecutorV1 *SwapExecutorV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.contract.Transact(opts, method, params...)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x180861fc.
//
// Solidity: function executeSwap(address[] pairs, address[] path, uint256[] amountsOut, address[][] pairTokens) returns()
func (_SwapExecutorV1 *SwapExecutorV1Transactor) ExecuteSwap(opts *bind.TransactOpts, pairs []common.Address, path []common.Address, amountsOut []*big.Int, pairTokens [][]common.Address) (*types.Transaction, error) {
	return _SwapExecutorV1.contract.Transact(opts, "executeSwap", pairs, path, amountsOut, pairTokens)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x180861fc.
//
// Solidity: function executeSwap(address[] pairs, address[] path, uint256[] amountsOut, address[][] pairTokens) returns()
func (_SwapExecutorV1 *SwapExecutorV1Session) ExecuteSwap(pairs []common.Address, path []common.Address, amountsOut []*big.Int, pairTokens [][]common.Address) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.ExecuteSwap(&_SwapExecutorV1.TransactOpts, pairs, path, amountsOut, pairTokens)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x180861fc.
//
// Solidity: function executeSwap(address[] pairs, address[] path, uint256[] amountsOut, address[][] pairTokens) returns()
func (_SwapExecutorV1 *SwapExecutorV1TransactorSession) ExecuteSwap(pairs []common.Address, path []common.Address, amountsOut []*big.Int, pairTokens [][]common.Address) (*types.Transaction, error) {
	return _SwapExecutorV1.Contract.ExecuteSwap(&_SwapExecutorV1.TransactOpts, pairs, path, amountsOut, pairTokens)
}
