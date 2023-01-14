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

// FlashloanParameters is an auto generated low-level Go binding around an user-defined struct.
type FlashloanParameters struct {
	Pairs                 []common.Address
	Reserves              [][]*big.Int
	PairTokens            [][]common.Address
	Path                  []common.Address
	AmountsOut            []*big.Int
	BorrowFee             *big.Int
	RevertOnReserveChange bool
}

// FlashloanExecutorV2MetaData contains all meta data concerning the FlashloanExecutorV2 contract.
var FlashloanExecutorV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BSCswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BiswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DooarSwapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinsCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SaitaSwapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TRUTHCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"W3swapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"annexCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"babyCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"croDefiSwapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"definixCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"elkCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"Pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"Reserves\",\"type\":\"uint256[][]\"},{\"internalType\":\"address[][]\",\"name\":\"PairTokens\",\"type\":\"address[][]\"},{\"internalType\":\"address[]\",\"name\":\"Path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"AmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"BorrowFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"RevertOnReserveChange\",\"type\":\"bool\"}],\"internalType\":\"structFlashloanParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"executeFlashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"fstswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"jetswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"nomiswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pantherCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pinkswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"stableXCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"wardenCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FlashloanExecutorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashloanExecutorV2MetaData.ABI instead.
var FlashloanExecutorV2ABI = FlashloanExecutorV2MetaData.ABI

// FlashloanExecutorV2 is an auto generated Go binding around an Ethereum contract.
type FlashloanExecutorV2 struct {
	FlashloanExecutorV2Caller     // Read-only binding to the contract
	FlashloanExecutorV2Transactor // Write-only binding to the contract
	FlashloanExecutorV2Filterer   // Log filterer for contract events
}

// FlashloanExecutorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FlashloanExecutorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashloanExecutorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashloanExecutorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashloanExecutorV2Session struct {
	Contract     *FlashloanExecutorV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlashloanExecutorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashloanExecutorV2CallerSession struct {
	Contract *FlashloanExecutorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FlashloanExecutorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashloanExecutorV2TransactorSession struct {
	Contract     *FlashloanExecutorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FlashloanExecutorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FlashloanExecutorV2Raw struct {
	Contract *FlashloanExecutorV2 // Generic contract binding to access the raw methods on
}

// FlashloanExecutorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashloanExecutorV2CallerRaw struct {
	Contract *FlashloanExecutorV2Caller // Generic read-only contract binding to access the raw methods on
}

// FlashloanExecutorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashloanExecutorV2TransactorRaw struct {
	Contract *FlashloanExecutorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashloanExecutorV2 creates a new instance of FlashloanExecutorV2, bound to a specific deployed contract.
func NewFlashloanExecutorV2(address common.Address, backend bind.ContractBackend) (*FlashloanExecutorV2, error) {
	contract, err := bindFlashloanExecutorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV2{FlashloanExecutorV2Caller: FlashloanExecutorV2Caller{contract: contract}, FlashloanExecutorV2Transactor: FlashloanExecutorV2Transactor{contract: contract}, FlashloanExecutorV2Filterer: FlashloanExecutorV2Filterer{contract: contract}}, nil
}

// NewFlashloanExecutorV2Caller creates a new read-only instance of FlashloanExecutorV2, bound to a specific deployed contract.
func NewFlashloanExecutorV2Caller(address common.Address, caller bind.ContractCaller) (*FlashloanExecutorV2Caller, error) {
	contract, err := bindFlashloanExecutorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV2Caller{contract: contract}, nil
}

// NewFlashloanExecutorV2Transactor creates a new write-only instance of FlashloanExecutorV2, bound to a specific deployed contract.
func NewFlashloanExecutorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FlashloanExecutorV2Transactor, error) {
	contract, err := bindFlashloanExecutorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV2Transactor{contract: contract}, nil
}

// NewFlashloanExecutorV2Filterer creates a new log filterer instance of FlashloanExecutorV2, bound to a specific deployed contract.
func NewFlashloanExecutorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FlashloanExecutorV2Filterer, error) {
	contract, err := bindFlashloanExecutorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV2Filterer{contract: contract}, nil
}

// bindFlashloanExecutorV2 binds a generic wrapper to an already deployed contract.
func bindFlashloanExecutorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashloanExecutorV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashloanExecutorV2 *FlashloanExecutorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashloanExecutorV2.Contract.FlashloanExecutorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashloanExecutorV2 *FlashloanExecutorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FlashloanExecutorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashloanExecutorV2 *FlashloanExecutorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FlashloanExecutorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashloanExecutorV2 *FlashloanExecutorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashloanExecutorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.contract.Transact(opts, method, params...)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) BSCswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "BSCswapCall", sender, amount0, amount1, data)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) BSCswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BSCswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) BSCswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BSCswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) BiswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "BiswapCall", sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) BiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BiswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) BiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BiswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) DooarSwapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "DooarSwapV2Call", sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) DooarSwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.DooarSwapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) DooarSwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.DooarSwapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) FinsCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "FinsCall", sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) FinsCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FinsCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) FinsCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FinsCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) SaitaSwapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "SaitaSwapCall", sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) SaitaSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SaitaSwapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) SaitaSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SaitaSwapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// TRUTHCall is a paid mutator transaction binding the contract method 0x48f22ad9.
//
// Solidity: function TRUTHCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) TRUTHCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "TRUTHCall", sender, amount0, amount1, data)
}

// TRUTHCall is a paid mutator transaction binding the contract method 0x48f22ad9.
//
// Solidity: function TRUTHCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) TRUTHCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.TRUTHCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// TRUTHCall is a paid mutator transaction binding the contract method 0x48f22ad9.
//
// Solidity: function TRUTHCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) TRUTHCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.TRUTHCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// W3swapCall is a paid mutator transaction binding the contract method 0xced54c51.
//
// Solidity: function W3swapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) W3swapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "W3swapCall", sender, amount0, amount1, data)
}

// W3swapCall is a paid mutator transaction binding the contract method 0xced54c51.
//
// Solidity: function W3swapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) W3swapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.W3swapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// W3swapCall is a paid mutator transaction binding the contract method 0xced54c51.
//
// Solidity: function W3swapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) W3swapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.W3swapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// AnnexCall is a paid mutator transaction binding the contract method 0x40c77747.
//
// Solidity: function annexCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) AnnexCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "annexCall", sender, amount0, amount1, data)
}

// AnnexCall is a paid mutator transaction binding the contract method 0x40c77747.
//
// Solidity: function annexCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) AnnexCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.AnnexCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// AnnexCall is a paid mutator transaction binding the contract method 0x40c77747.
//
// Solidity: function annexCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) AnnexCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.AnnexCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) BabyCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "babyCall", sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) BabyCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BabyCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) BabyCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.BabyCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) CroDefiSwapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "croDefiSwapCall", sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) CroDefiSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.CroDefiSwapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) CroDefiSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.CroDefiSwapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) DefinixCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "definixCall", sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) DefinixCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.DefinixCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) DefinixCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.DefinixCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) ElkCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "elkCall", sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) ElkCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.ElkCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) ElkCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.ElkCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x46cd836f.
//
// Solidity: function executeFlashloan((address[],uint256[][],address[][],address[],uint256[],uint256,bool) params) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) ExecuteFlashloan(opts *bind.TransactOpts, params FlashloanParameters) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "executeFlashloan", params)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x46cd836f.
//
// Solidity: function executeFlashloan((address[],uint256[][],address[][],address[],uint256[],uint256,bool) params) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) ExecuteFlashloan(params FlashloanParameters) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.ExecuteFlashloan(&_FlashloanExecutorV2.TransactOpts, params)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x46cd836f.
//
// Solidity: function executeFlashloan((address[],uint256[][],address[][],address[],uint256[],uint256,bool) params) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) ExecuteFlashloan(params FlashloanParameters) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.ExecuteFlashloan(&_FlashloanExecutorV2.TransactOpts, params)
}

// FstswapCall is a paid mutator transaction binding the contract method 0x88f9eddd.
//
// Solidity: function fstswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) FstswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "fstswapCall", sender, amount0, amount1, data)
}

// FstswapCall is a paid mutator transaction binding the contract method 0x88f9eddd.
//
// Solidity: function fstswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) FstswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FstswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// FstswapCall is a paid mutator transaction binding the contract method 0x88f9eddd.
//
// Solidity: function fstswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) FstswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.FstswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) JetswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "jetswapCall", sender, amount0, amount1, data)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) JetswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.JetswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) JetswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.JetswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) NomiswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "nomiswapCall", sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) NomiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.NomiswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) NomiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.NomiswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) PancakeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "pancakeCall", sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PancakeCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PancakeCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) PantherCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "pantherCall", sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) PantherCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PantherCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) PantherCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PantherCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) PinkswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "pinkswapCall", sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) PinkswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PinkswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) PinkswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.PinkswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) SafeswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "safeswapCall", sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) SafeswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SafeswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) SafeswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SafeswapCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) StableXCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "stableXCall", sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) StableXCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.StableXCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) StableXCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.StableXCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) SwapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "swapV2Call", sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) SwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SwapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) SwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.SwapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.UniswapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.UniswapV2Call(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) WardenCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.Transact(opts, "wardenCall", sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) WardenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.WardenCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) WardenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.WardenCall(&_FlashloanExecutorV2.TransactOpts, sender, amount0, amount1, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2Session) Receive() (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.Receive(&_FlashloanExecutorV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV2 *FlashloanExecutorV2TransactorSession) Receive() (*types.Transaction, error) {
	return _FlashloanExecutorV2.Contract.Receive(&_FlashloanExecutorV2.TransactOpts)
}
