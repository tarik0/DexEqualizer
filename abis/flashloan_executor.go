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
	Pairs      []common.Address
	Path       []common.Address
	AmountsOut []*big.Int
	PairTokens [][]common.Address
	BorrowFee  *big.Int
}

// FlashloanExecutorV1MetaData contains all meta data concerning the FlashloanExecutorV1 contract.
var FlashloanExecutorV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BSCswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BiswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DooarSwapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinsCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SaitaSwapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"babyCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"croDefiSwapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"definixCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"elkCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"Pairs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"Path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"AmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"address[][]\",\"name\":\"PairTokens\",\"type\":\"address[][]\"},{\"internalType\":\"uint256\",\"name\":\"BorrowFee\",\"type\":\"uint256\"}],\"internalType\":\"structFlashloanParameters\",\"name\":\"opts\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"encryptParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"encrypted\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"paramsBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"executeFlashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"jetswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"nomiswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pantherCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pinkswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeswapCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"stableXCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"wardenCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FlashloanExecutorV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashloanExecutorV1MetaData.ABI instead.
var FlashloanExecutorV1ABI = FlashloanExecutorV1MetaData.ABI

// FlashloanExecutorV1 is an auto generated Go binding around an Ethereum contract.
type FlashloanExecutorV1 struct {
	FlashloanExecutorV1Caller     // Read-only binding to the contract
	FlashloanExecutorV1Transactor // Write-only binding to the contract
	FlashloanExecutorV1Filterer   // Log filterer for contract events
}

// FlashloanExecutorV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FlashloanExecutorV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashloanExecutorV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashloanExecutorV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanExecutorV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashloanExecutorV1Session struct {
	Contract     *FlashloanExecutorV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlashloanExecutorV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashloanExecutorV1CallerSession struct {
	Contract *FlashloanExecutorV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FlashloanExecutorV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashloanExecutorV1TransactorSession struct {
	Contract     *FlashloanExecutorV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FlashloanExecutorV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FlashloanExecutorV1Raw struct {
	Contract *FlashloanExecutorV1 // Generic contract binding to access the raw methods on
}

// FlashloanExecutorV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashloanExecutorV1CallerRaw struct {
	Contract *FlashloanExecutorV1Caller // Generic read-only contract binding to access the raw methods on
}

// FlashloanExecutorV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashloanExecutorV1TransactorRaw struct {
	Contract *FlashloanExecutorV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashloanExecutorV1 creates a new instance of FlashloanExecutorV1, bound to a specific deployed contract.
func NewFlashloanExecutorV1(address common.Address, backend bind.ContractBackend) (*FlashloanExecutorV1, error) {
	contract, err := bindFlashloanExecutorV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV1{FlashloanExecutorV1Caller: FlashloanExecutorV1Caller{contract: contract}, FlashloanExecutorV1Transactor: FlashloanExecutorV1Transactor{contract: contract}, FlashloanExecutorV1Filterer: FlashloanExecutorV1Filterer{contract: contract}}, nil
}

// NewFlashloanExecutorV1Caller creates a new read-only instance of FlashloanExecutorV1, bound to a specific deployed contract.
func NewFlashloanExecutorV1Caller(address common.Address, caller bind.ContractCaller) (*FlashloanExecutorV1Caller, error) {
	contract, err := bindFlashloanExecutorV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV1Caller{contract: contract}, nil
}

// NewFlashloanExecutorV1Transactor creates a new write-only instance of FlashloanExecutorV1, bound to a specific deployed contract.
func NewFlashloanExecutorV1Transactor(address common.Address, transactor bind.ContractTransactor) (*FlashloanExecutorV1Transactor, error) {
	contract, err := bindFlashloanExecutorV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV1Transactor{contract: contract}, nil
}

// NewFlashloanExecutorV1Filterer creates a new log filterer instance of FlashloanExecutorV1, bound to a specific deployed contract.
func NewFlashloanExecutorV1Filterer(address common.Address, filterer bind.ContractFilterer) (*FlashloanExecutorV1Filterer, error) {
	contract, err := bindFlashloanExecutorV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashloanExecutorV1Filterer{contract: contract}, nil
}

// bindFlashloanExecutorV1 binds a generic wrapper to an already deployed contract.
func bindFlashloanExecutorV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashloanExecutorV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashloanExecutorV1 *FlashloanExecutorV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashloanExecutorV1.Contract.FlashloanExecutorV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashloanExecutorV1 *FlashloanExecutorV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.FlashloanExecutorV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashloanExecutorV1 *FlashloanExecutorV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.FlashloanExecutorV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashloanExecutorV1 *FlashloanExecutorV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashloanExecutorV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.contract.Transact(opts, method, params...)
}

// EncryptParams is a free data retrieval call binding the contract method 0xa33f388b.
//
// Solidity: function encryptParams((address[],address[],uint256[],address[][],uint256) opts, bytes key) pure returns(bytes encrypted)
func (_FlashloanExecutorV1 *FlashloanExecutorV1Caller) EncryptParams(opts *bind.CallOpts, params FlashloanParameters, key []byte) ([]byte, error) {
	var out []interface{}
	err := _FlashloanExecutorV1.contract.Call(opts, &out, "encryptParams", params, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncryptParams is a free data retrieval call binding the contract method 0xa33f388b.
//
// Solidity: function encryptParams((address[],address[],uint256[],address[][],uint256) opts, bytes key) pure returns(bytes encrypted)
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) EncryptParams(opts FlashloanParameters, key []byte) ([]byte, error) {
	return _FlashloanExecutorV1.Contract.EncryptParams(&_FlashloanExecutorV1.CallOpts, opts, key)
}

// EncryptParams is a free data retrieval call binding the contract method 0xa33f388b.
//
// Solidity: function encryptParams((address[],address[],uint256[],address[][],uint256) opts, bytes key) pure returns(bytes encrypted)
func (_FlashloanExecutorV1 *FlashloanExecutorV1CallerSession) EncryptParams(opts FlashloanParameters, key []byte) ([]byte, error) {
	return _FlashloanExecutorV1.Contract.EncryptParams(&_FlashloanExecutorV1.CallOpts, opts, key)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) BSCswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "BSCswapCall", sender, amount0, amount1, data)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) BSCswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BSCswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// BSCswapCall is a paid mutator transaction binding the contract method 0x75908f7c.
//
// Solidity: function BSCswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) BSCswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BSCswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) BiswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "BiswapCall", sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) BiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BiswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// BiswapCall is a paid mutator transaction binding the contract method 0x5b3bc4fe.
//
// Solidity: function BiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) BiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BiswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) DooarSwapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "DooarSwapV2Call", sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) DooarSwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.DooarSwapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// DooarSwapV2Call is a paid mutator transaction binding the contract method 0x330f9b41.
//
// Solidity: function DooarSwapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) DooarSwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.DooarSwapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) FinsCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "FinsCall", sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) FinsCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.FinsCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// FinsCall is a paid mutator transaction binding the contract method 0xf17194aa.
//
// Solidity: function FinsCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) FinsCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.FinsCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) SaitaSwapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "SaitaSwapCall", sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) SaitaSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SaitaSwapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SaitaSwapCall is a paid mutator transaction binding the contract method 0xfe1b3a66.
//
// Solidity: function SaitaSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) SaitaSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SaitaSwapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) BabyCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "babyCall", sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) BabyCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BabyCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// BabyCall is a paid mutator transaction binding the contract method 0x110c03de.
//
// Solidity: function babyCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) BabyCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.BabyCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) CroDefiSwapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "croDefiSwapCall", sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) CroDefiSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.CroDefiSwapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// CroDefiSwapCall is a paid mutator transaction binding the contract method 0x6c813d29.
//
// Solidity: function croDefiSwapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) CroDefiSwapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.CroDefiSwapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) DefinixCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "definixCall", sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) DefinixCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.DefinixCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// DefinixCall is a paid mutator transaction binding the contract method 0x0c33efd3.
//
// Solidity: function definixCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) DefinixCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.DefinixCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) ElkCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "elkCall", sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) ElkCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.ElkCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// ElkCall is a paid mutator transaction binding the contract method 0x07d3513a.
//
// Solidity: function elkCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) ElkCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.ElkCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x9382ad06.
//
// Solidity: function executeFlashloan(bytes paramsBytes, bytes key) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) ExecuteFlashloan(opts *bind.TransactOpts, paramsBytes []byte, key []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "executeFlashloan", paramsBytes, key)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x9382ad06.
//
// Solidity: function executeFlashloan(bytes paramsBytes, bytes key) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) ExecuteFlashloan(paramsBytes []byte, key []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.ExecuteFlashloan(&_FlashloanExecutorV1.TransactOpts, paramsBytes, key)
}

// ExecuteFlashloan is a paid mutator transaction binding the contract method 0x9382ad06.
//
// Solidity: function executeFlashloan(bytes paramsBytes, bytes key) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) ExecuteFlashloan(paramsBytes []byte, key []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.ExecuteFlashloan(&_FlashloanExecutorV1.TransactOpts, paramsBytes, key)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) JetswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "jetswapCall", sender, amount0, amount1, data)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) JetswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.JetswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// JetswapCall is a paid mutator transaction binding the contract method 0x3fc01685.
//
// Solidity: function jetswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) JetswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.JetswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) NomiswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "nomiswapCall", sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) NomiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.NomiswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// NomiswapCall is a paid mutator transaction binding the contract method 0x22109682.
//
// Solidity: function nomiswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) NomiswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.NomiswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) PancakeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "pancakeCall", sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PancakeCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PancakeCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) PantherCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "pantherCall", sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) PantherCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PantherCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PantherCall is a paid mutator transaction binding the contract method 0x1c8f37b3.
//
// Solidity: function pantherCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) PantherCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PantherCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) PinkswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "pinkswapCall", sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) PinkswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PinkswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// PinkswapCall is a paid mutator transaction binding the contract method 0x5ddd1198.
//
// Solidity: function pinkswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) PinkswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.PinkswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) SafeswapCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "safeswapCall", sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) SafeswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SafeswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SafeswapCall is a paid mutator transaction binding the contract method 0x3cc9c6b4.
//
// Solidity: function safeswapCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) SafeswapCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SafeswapCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) StableXCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "stableXCall", sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) StableXCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.StableXCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// StableXCall is a paid mutator transaction binding the contract method 0x5aec284b.
//
// Solidity: function stableXCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) StableXCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.StableXCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) SwapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "swapV2Call", sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) SwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SwapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// SwapV2Call is a paid mutator transaction binding the contract method 0xb2ff9f26.
//
// Solidity: function swapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) SwapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.SwapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.UniswapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.UniswapV2Call(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) WardenCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.Transact(opts, "wardenCall", sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) WardenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.WardenCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// WardenCall is a paid mutator transaction binding the contract method 0x46337f3a.
//
// Solidity: function wardenCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) WardenCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.WardenCall(&_FlashloanExecutorV1.TransactOpts, sender, amount0, amount1, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashloanExecutorV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1Session) Receive() (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.Receive(&_FlashloanExecutorV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashloanExecutorV1 *FlashloanExecutorV1TransactorSession) Receive() (*types.Transaction, error) {
	return _FlashloanExecutorV1.Contract.Receive(&_FlashloanExecutorV1.TransactOpts)
}

