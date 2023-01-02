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

// MulticallCall is an auto generated low-level Go binding around an user-defined struct.
type MulticallCall struct {
	Target   common.Address
	CallData []byte
}

// MulticallerMetaData contains all meta data concerning the Multicaller contract.
var MulticallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MulticallerABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallerMetaData.ABI instead.
var MulticallerABI = MulticallerMetaData.ABI

// Multicaller is an auto generated Go binding around an Ethereum contract.
type Multicaller struct {
	MulticallerCaller     // Read-only binding to the contract
	MulticallerTransactor // Write-only binding to the contract
	MulticallerFilterer   // Log filterer for contract events
}

// MulticallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallerSession struct {
	Contract     *Multicaller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallerCallerSession struct {
	Contract *MulticallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MulticallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallerTransactorSession struct {
	Contract     *MulticallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MulticallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallerRaw struct {
	Contract *Multicaller // Generic contract binding to access the raw methods on
}

// MulticallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallerCallerRaw struct {
	Contract *MulticallerCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallerTransactorRaw struct {
	Contract *MulticallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticaller creates a new instance of Multicaller, bound to a specific deployed contract.
func NewMulticaller(address common.Address, backend bind.ContractBackend) (*Multicaller, error) {
	contract, err := bindMulticaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicaller{MulticallerCaller: MulticallerCaller{contract: contract}, MulticallerTransactor: MulticallerTransactor{contract: contract}, MulticallerFilterer: MulticallerFilterer{contract: contract}}, nil
}

// NewMulticallerCaller creates a new read-only instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerCaller(address common.Address, caller bind.ContractCaller) (*MulticallerCaller, error) {
	contract, err := bindMulticaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallerCaller{contract: contract}, nil
}

// NewMulticallerTransactor creates a new write-only instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallerTransactor, error) {
	contract, err := bindMulticaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallerTransactor{contract: contract}, nil
}

// NewMulticallerFilterer creates a new log filterer instance of Multicaller, bound to a specific deployed contract.
func NewMulticallerFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallerFilterer, error) {
	contract, err := bindMulticaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallerFilterer{contract: contract}, nil
}

// bindMulticaller binds a generic wrapper to an already deployed contract.
func bindMulticaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicaller *MulticallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicaller.Contract.MulticallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicaller *MulticallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicaller.Contract.MulticallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicaller *MulticallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicaller.Contract.MulticallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicaller *MulticallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicaller *MulticallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicaller *MulticallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicaller.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicaller *MulticallerCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicaller *MulticallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicaller.Contract.GetBlockHash(&_Multicaller.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_Multicaller *MulticallerCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Multicaller.Contract.GetBlockHash(&_Multicaller.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicaller *MulticallerCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicaller *MulticallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicaller.Contract.GetCurrentBlockCoinbase(&_Multicaller.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_Multicaller *MulticallerCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Multicaller.Contract.GetCurrentBlockCoinbase(&_Multicaller.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicaller *MulticallerCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicaller *MulticallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockDifficulty(&_Multicaller.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_Multicaller *MulticallerCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockDifficulty(&_Multicaller.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicaller *MulticallerCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicaller *MulticallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockGasLimit(&_Multicaller.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_Multicaller *MulticallerCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockGasLimit(&_Multicaller.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicaller *MulticallerCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicaller *MulticallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockTimestamp(&_Multicaller.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_Multicaller *MulticallerCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Multicaller.Contract.GetCurrentBlockTimestamp(&_Multicaller.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicaller *MulticallerCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicaller *MulticallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicaller.Contract.GetEthBalance(&_Multicaller.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_Multicaller *MulticallerCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Multicaller.Contract.GetEthBalance(&_Multicaller.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicaller *MulticallerCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multicaller.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicaller *MulticallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicaller.Contract.GetLastBlockHash(&_Multicaller.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_Multicaller *MulticallerCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Multicaller.Contract.GetLastBlockHash(&_Multicaller.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicaller *MulticallerTransactor) Aggregate(opts *bind.TransactOpts, calls []MulticallCall) (*types.Transaction, error) {
	return _Multicaller.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicaller *MulticallerSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _Multicaller.Contract.Aggregate(&_Multicaller.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Multicaller *MulticallerTransactorSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _Multicaller.Contract.Aggregate(&_Multicaller.TransactOpts, calls)
}
