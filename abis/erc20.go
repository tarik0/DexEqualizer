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

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"LogRebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoLiquidityReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyBackFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyBackFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkFeeExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkMaxWalletExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkMaxWalletToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkSwapThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPercentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"clearStuckBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecosystemFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecosystemFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"enableTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCirculatingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accuracy\",\"type\":\"uint256\"}],\"name\":\"getLiquidityBacking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gonMaxWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialDistributionFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNotInSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accuracy\",\"type\":\"uint256\"}],\"name\":\"isOverLiquified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manualSync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairContract\",\"outputs\":[{\"internalType\":\"contractInterfaceLP\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"supplyDelta\",\"type\":\"int256\"}],\"name\":\"rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"rescueToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIDEXRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"sendPresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeExempt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_autoLiquidityReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ecosystemFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketingFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyBackFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceivers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ecosystemFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_buyBackFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeDenominator\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setInitialDistributionFinished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_master\",\"type\":\"address\"}],\"name\":\"setMaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaxWalletExempt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denom\",\"type\":\"uint256\"}],\"name\":\"setMaxWalletToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denom\",\"type\":\"uint256\"}],\"name\":\"setSwapBackSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accuracy\",\"type\":\"uint256\"}],\"name\":\"setTargetLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner_ common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner_, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner_ common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner_, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner_ common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner_, spender)
}

// AutoLiquidityReceiver is a free data retrieval call binding the contract method 0xca33e64c.
//
// Solidity: function autoLiquidityReceiver() view returns(address)
func (_ERC20 *ERC20Caller) AutoLiquidityReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "autoLiquidityReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AutoLiquidityReceiver is a free data retrieval call binding the contract method 0xca33e64c.
//
// Solidity: function autoLiquidityReceiver() view returns(address)
func (_ERC20 *ERC20Session) AutoLiquidityReceiver() (common.Address, error) {
	return _ERC20.Contract.AutoLiquidityReceiver(&_ERC20.CallOpts)
}

// AutoLiquidityReceiver is a free data retrieval call binding the contract method 0xca33e64c.
//
// Solidity: function autoLiquidityReceiver() view returns(address)
func (_ERC20 *ERC20CallerSession) AutoLiquidityReceiver() (common.Address, error) {
	return _ERC20.Contract.AutoLiquidityReceiver(&_ERC20.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BuyBackFee is a free data retrieval call binding the contract method 0x4be8f8b1.
//
// Solidity: function buyBackFee() view returns(uint256)
func (_ERC20 *ERC20Caller) BuyBackFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "buyBackFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyBackFee is a free data retrieval call binding the contract method 0x4be8f8b1.
//
// Solidity: function buyBackFee() view returns(uint256)
func (_ERC20 *ERC20Session) BuyBackFee() (*big.Int, error) {
	return _ERC20.Contract.BuyBackFee(&_ERC20.CallOpts)
}

// BuyBackFee is a free data retrieval call binding the contract method 0x4be8f8b1.
//
// Solidity: function buyBackFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) BuyBackFee() (*big.Int, error) {
	return _ERC20.Contract.BuyBackFee(&_ERC20.CallOpts)
}

// BuyBackFeeReceiver is a free data retrieval call binding the contract method 0x50271226.
//
// Solidity: function buyBackFeeReceiver() view returns(address)
func (_ERC20 *ERC20Caller) BuyBackFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "buyBackFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BuyBackFeeReceiver is a free data retrieval call binding the contract method 0x50271226.
//
// Solidity: function buyBackFeeReceiver() view returns(address)
func (_ERC20 *ERC20Session) BuyBackFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.BuyBackFeeReceiver(&_ERC20.CallOpts)
}

// BuyBackFeeReceiver is a free data retrieval call binding the contract method 0x50271226.
//
// Solidity: function buyBackFeeReceiver() view returns(address)
func (_ERC20 *ERC20CallerSession) BuyBackFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.BuyBackFeeReceiver(&_ERC20.CallOpts)
}

// CheckFeeExempt is a free data retrieval call binding the contract method 0xd4399790.
//
// Solidity: function checkFeeExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20Caller) CheckFeeExempt(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "checkFeeExempt", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFeeExempt is a free data retrieval call binding the contract method 0xd4399790.
//
// Solidity: function checkFeeExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20Session) CheckFeeExempt(_addr common.Address) (bool, error) {
	return _ERC20.Contract.CheckFeeExempt(&_ERC20.CallOpts, _addr)
}

// CheckFeeExempt is a free data retrieval call binding the contract method 0xd4399790.
//
// Solidity: function checkFeeExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20CallerSession) CheckFeeExempt(_addr common.Address) (bool, error) {
	return _ERC20.Contract.CheckFeeExempt(&_ERC20.CallOpts, _addr)
}

// CheckMaxWalletExempt is a free data retrieval call binding the contract method 0xddff51a8.
//
// Solidity: function checkMaxWalletExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20Caller) CheckMaxWalletExempt(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "checkMaxWalletExempt", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMaxWalletExempt is a free data retrieval call binding the contract method 0xddff51a8.
//
// Solidity: function checkMaxWalletExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20Session) CheckMaxWalletExempt(_addr common.Address) (bool, error) {
	return _ERC20.Contract.CheckMaxWalletExempt(&_ERC20.CallOpts, _addr)
}

// CheckMaxWalletExempt is a free data retrieval call binding the contract method 0xddff51a8.
//
// Solidity: function checkMaxWalletExempt(address _addr) view returns(bool)
func (_ERC20 *ERC20CallerSession) CheckMaxWalletExempt(_addr common.Address) (bool, error) {
	return _ERC20.Contract.CheckMaxWalletExempt(&_ERC20.CallOpts, _addr)
}

// CheckMaxWalletToken is a free data retrieval call binding the contract method 0xb43b7835.
//
// Solidity: function checkMaxWalletToken() view returns(uint256)
func (_ERC20 *ERC20Caller) CheckMaxWalletToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "checkMaxWalletToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckMaxWalletToken is a free data retrieval call binding the contract method 0xb43b7835.
//
// Solidity: function checkMaxWalletToken() view returns(uint256)
func (_ERC20 *ERC20Session) CheckMaxWalletToken() (*big.Int, error) {
	return _ERC20.Contract.CheckMaxWalletToken(&_ERC20.CallOpts)
}

// CheckMaxWalletToken is a free data retrieval call binding the contract method 0xb43b7835.
//
// Solidity: function checkMaxWalletToken() view returns(uint256)
func (_ERC20 *ERC20CallerSession) CheckMaxWalletToken() (*big.Int, error) {
	return _ERC20.Contract.CheckMaxWalletToken(&_ERC20.CallOpts)
}

// CheckSwapThreshold is a free data retrieval call binding the contract method 0x6d351d1a.
//
// Solidity: function checkSwapThreshold() view returns(uint256)
func (_ERC20 *ERC20Caller) CheckSwapThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "checkSwapThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckSwapThreshold is a free data retrieval call binding the contract method 0x6d351d1a.
//
// Solidity: function checkSwapThreshold() view returns(uint256)
func (_ERC20 *ERC20Session) CheckSwapThreshold() (*big.Int, error) {
	return _ERC20.Contract.CheckSwapThreshold(&_ERC20.CallOpts)
}

// CheckSwapThreshold is a free data retrieval call binding the contract method 0x6d351d1a.
//
// Solidity: function checkSwapThreshold() view returns(uint256)
func (_ERC20 *ERC20CallerSession) CheckSwapThreshold() (*big.Int, error) {
	return _ERC20.Contract.CheckSwapThreshold(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// EcosystemFee is a free data retrieval call binding the contract method 0xb1fb0e97.
//
// Solidity: function ecosystemFee() view returns(uint256)
func (_ERC20 *ERC20Caller) EcosystemFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "ecosystemFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EcosystemFee is a free data retrieval call binding the contract method 0xb1fb0e97.
//
// Solidity: function ecosystemFee() view returns(uint256)
func (_ERC20 *ERC20Session) EcosystemFee() (*big.Int, error) {
	return _ERC20.Contract.EcosystemFee(&_ERC20.CallOpts)
}

// EcosystemFee is a free data retrieval call binding the contract method 0xb1fb0e97.
//
// Solidity: function ecosystemFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) EcosystemFee() (*big.Int, error) {
	return _ERC20.Contract.EcosystemFee(&_ERC20.CallOpts)
}

// EcosystemFeeReceiver is a free data retrieval call binding the contract method 0xb20ae67f.
//
// Solidity: function ecosystemFeeReceiver() view returns(address)
func (_ERC20 *ERC20Caller) EcosystemFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "ecosystemFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcosystemFeeReceiver is a free data retrieval call binding the contract method 0xb20ae67f.
//
// Solidity: function ecosystemFeeReceiver() view returns(address)
func (_ERC20 *ERC20Session) EcosystemFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.EcosystemFeeReceiver(&_ERC20.CallOpts)
}

// EcosystemFeeReceiver is a free data retrieval call binding the contract method 0xb20ae67f.
//
// Solidity: function ecosystemFeeReceiver() view returns(address)
func (_ERC20 *ERC20CallerSession) EcosystemFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.EcosystemFeeReceiver(&_ERC20.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_ERC20 *ERC20Caller) FeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "feeDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_ERC20 *ERC20Session) FeeDenominator() (*big.Int, error) {
	return _ERC20.Contract.FeeDenominator(&_ERC20.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_ERC20 *ERC20CallerSession) FeeDenominator() (*big.Int, error) {
	return _ERC20.Contract.FeeDenominator(&_ERC20.CallOpts)
}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) GetCirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getCirculatingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_ERC20 *ERC20Session) GetCirculatingSupply() (*big.Int, error) {
	return _ERC20.Contract.GetCirculatingSupply(&_ERC20.CallOpts)
}

// GetCirculatingSupply is a free data retrieval call binding the contract method 0x2b112e49.
//
// Solidity: function getCirculatingSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) GetCirculatingSupply() (*big.Int, error) {
	return _ERC20.Contract.GetCirculatingSupply(&_ERC20.CallOpts)
}

// GetLiquidityBacking is a free data retrieval call binding the contract method 0xd51ed1c8.
//
// Solidity: function getLiquidityBacking(uint256 accuracy) view returns(uint256)
func (_ERC20 *ERC20Caller) GetLiquidityBacking(opts *bind.CallOpts, accuracy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getLiquidityBacking", accuracy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidityBacking is a free data retrieval call binding the contract method 0xd51ed1c8.
//
// Solidity: function getLiquidityBacking(uint256 accuracy) view returns(uint256)
func (_ERC20 *ERC20Session) GetLiquidityBacking(accuracy *big.Int) (*big.Int, error) {
	return _ERC20.Contract.GetLiquidityBacking(&_ERC20.CallOpts, accuracy)
}

// GetLiquidityBacking is a free data retrieval call binding the contract method 0xd51ed1c8.
//
// Solidity: function getLiquidityBacking(uint256 accuracy) view returns(uint256)
func (_ERC20 *ERC20CallerSession) GetLiquidityBacking(accuracy *big.Int) (*big.Int, error) {
	return _ERC20.Contract.GetLiquidityBacking(&_ERC20.CallOpts, accuracy)
}

// GonMaxWallet is a free data retrieval call binding the contract method 0x7d5a3fa0.
//
// Solidity: function gonMaxWallet() view returns(uint256)
func (_ERC20 *ERC20Caller) GonMaxWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "gonMaxWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GonMaxWallet is a free data retrieval call binding the contract method 0x7d5a3fa0.
//
// Solidity: function gonMaxWallet() view returns(uint256)
func (_ERC20 *ERC20Session) GonMaxWallet() (*big.Int, error) {
	return _ERC20.Contract.GonMaxWallet(&_ERC20.CallOpts)
}

// GonMaxWallet is a free data retrieval call binding the contract method 0x7d5a3fa0.
//
// Solidity: function gonMaxWallet() view returns(uint256)
func (_ERC20 *ERC20CallerSession) GonMaxWallet() (*big.Int, error) {
	return _ERC20.Contract.GonMaxWallet(&_ERC20.CallOpts)
}

// InitialDistributionFinished is a free data retrieval call binding the contract method 0xd1fce264.
//
// Solidity: function initialDistributionFinished() view returns(bool)
func (_ERC20 *ERC20Caller) InitialDistributionFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "initialDistributionFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InitialDistributionFinished is a free data retrieval call binding the contract method 0xd1fce264.
//
// Solidity: function initialDistributionFinished() view returns(bool)
func (_ERC20 *ERC20Session) InitialDistributionFinished() (bool, error) {
	return _ERC20.Contract.InitialDistributionFinished(&_ERC20.CallOpts)
}

// InitialDistributionFinished is a free data retrieval call binding the contract method 0xd1fce264.
//
// Solidity: function initialDistributionFinished() view returns(bool)
func (_ERC20 *ERC20CallerSession) InitialDistributionFinished() (bool, error) {
	return _ERC20.Contract.InitialDistributionFinished(&_ERC20.CallOpts)
}

// IsNotInSwap is a free data retrieval call binding the contract method 0x83b4ac68.
//
// Solidity: function isNotInSwap() view returns(bool)
func (_ERC20 *ERC20Caller) IsNotInSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "isNotInSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotInSwap is a free data retrieval call binding the contract method 0x83b4ac68.
//
// Solidity: function isNotInSwap() view returns(bool)
func (_ERC20 *ERC20Session) IsNotInSwap() (bool, error) {
	return _ERC20.Contract.IsNotInSwap(&_ERC20.CallOpts)
}

// IsNotInSwap is a free data retrieval call binding the contract method 0x83b4ac68.
//
// Solidity: function isNotInSwap() view returns(bool)
func (_ERC20 *ERC20CallerSession) IsNotInSwap() (bool, error) {
	return _ERC20.Contract.IsNotInSwap(&_ERC20.CallOpts)
}

// IsOverLiquified is a free data retrieval call binding the contract method 0x1161ae39.
//
// Solidity: function isOverLiquified(uint256 target, uint256 accuracy) view returns(bool)
func (_ERC20 *ERC20Caller) IsOverLiquified(opts *bind.CallOpts, target *big.Int, accuracy *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "isOverLiquified", target, accuracy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOverLiquified is a free data retrieval call binding the contract method 0x1161ae39.
//
// Solidity: function isOverLiquified(uint256 target, uint256 accuracy) view returns(bool)
func (_ERC20 *ERC20Session) IsOverLiquified(target *big.Int, accuracy *big.Int) (bool, error) {
	return _ERC20.Contract.IsOverLiquified(&_ERC20.CallOpts, target, accuracy)
}

// IsOverLiquified is a free data retrieval call binding the contract method 0x1161ae39.
//
// Solidity: function isOverLiquified(uint256 target, uint256 accuracy) view returns(bool)
func (_ERC20 *ERC20CallerSession) IsOverLiquified(target *big.Int, accuracy *big.Int) (bool, error) {
	return _ERC20.Contract.IsOverLiquified(&_ERC20.CallOpts, target, accuracy)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20 *ERC20Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20 *ERC20Session) IsOwner() (bool, error) {
	return _ERC20.Contract.IsOwner(&_ERC20.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20 *ERC20CallerSession) IsOwner() (bool, error) {
	return _ERC20.Contract.IsOwner(&_ERC20.CallOpts)
}

// LiquidityFee is a free data retrieval call binding the contract method 0x98118cb4.
//
// Solidity: function liquidityFee() view returns(uint256)
func (_ERC20 *ERC20Caller) LiquidityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "liquidityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityFee is a free data retrieval call binding the contract method 0x98118cb4.
//
// Solidity: function liquidityFee() view returns(uint256)
func (_ERC20 *ERC20Session) LiquidityFee() (*big.Int, error) {
	return _ERC20.Contract.LiquidityFee(&_ERC20.CallOpts)
}

// LiquidityFee is a free data retrieval call binding the contract method 0x98118cb4.
//
// Solidity: function liquidityFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) LiquidityFee() (*big.Int, error) {
	return _ERC20.Contract.LiquidityFee(&_ERC20.CallOpts)
}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_ERC20 *ERC20Caller) MarketingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "marketingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_ERC20 *ERC20Session) MarketingFee() (*big.Int, error) {
	return _ERC20.Contract.MarketingFee(&_ERC20.CallOpts)
}

// MarketingFee is a free data retrieval call binding the contract method 0x6b67c4df.
//
// Solidity: function marketingFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) MarketingFee() (*big.Int, error) {
	return _ERC20.Contract.MarketingFee(&_ERC20.CallOpts)
}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_ERC20 *ERC20Caller) MarketingFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "marketingFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_ERC20 *ERC20Session) MarketingFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.MarketingFeeReceiver(&_ERC20.CallOpts)
}

// MarketingFeeReceiver is a free data retrieval call binding the contract method 0xe96fada2.
//
// Solidity: function marketingFeeReceiver() view returns(address)
func (_ERC20 *ERC20CallerSession) MarketingFeeReceiver() (common.Address, error) {
	return _ERC20.Contract.MarketingFeeReceiver(&_ERC20.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ERC20 *ERC20Caller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ERC20 *ERC20Session) Master() (common.Address, error) {
	return _ERC20.Contract.Master(&_ERC20.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ERC20 *ERC20CallerSession) Master() (common.Address, error) {
	return _ERC20.Contract.Master(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Session) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20CallerSession) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_ERC20 *ERC20Caller) Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_ERC20 *ERC20Session) Pair() (common.Address, error) {
	return _ERC20.Contract.Pair(&_ERC20.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_ERC20 *ERC20CallerSession) Pair() (common.Address, error) {
	return _ERC20.Contract.Pair(&_ERC20.CallOpts)
}

// PairContract is a free data retrieval call binding the contract method 0x4d709adf.
//
// Solidity: function pairContract() view returns(address)
func (_ERC20 *ERC20Caller) PairContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "pairContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairContract is a free data retrieval call binding the contract method 0x4d709adf.
//
// Solidity: function pairContract() view returns(address)
func (_ERC20 *ERC20Session) PairContract() (common.Address, error) {
	return _ERC20.Contract.PairContract(&_ERC20.CallOpts)
}

// PairContract is a free data retrieval call binding the contract method 0x4d709adf.
//
// Solidity: function pairContract() view returns(address)
func (_ERC20 *ERC20CallerSession) PairContract() (common.Address, error) {
	return _ERC20.Contract.PairContract(&_ERC20.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ERC20 *ERC20Caller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ERC20 *ERC20Session) Router() (common.Address, error) {
	return _ERC20.Contract.Router(&_ERC20.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ERC20 *ERC20CallerSession) Router() (common.Address, error) {
	return _ERC20.Contract.Router(&_ERC20.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_ERC20 *ERC20Caller) SwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "swapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_ERC20 *ERC20Session) SwapEnabled() (bool, error) {
	return _ERC20.Contract.SwapEnabled(&_ERC20.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() view returns(bool)
func (_ERC20 *ERC20CallerSession) SwapEnabled() (bool, error) {
	return _ERC20.Contract.SwapEnabled(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_ERC20 *ERC20Session) TotalFee() (*big.Int, error) {
	return _ERC20.Contract.TotalFee(&_ERC20.CallOpts)
}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalFee() (*big.Int, error) {
	return _ERC20.Contract.TotalFee(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x56a227f2.
//
// Solidity: function clearStuckBalance(uint256 amountPercentage, address adr) returns()
func (_ERC20 *ERC20Transactor) ClearStuckBalance(opts *bind.TransactOpts, amountPercentage *big.Int, adr common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "clearStuckBalance", amountPercentage, adr)
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x56a227f2.
//
// Solidity: function clearStuckBalance(uint256 amountPercentage, address adr) returns()
func (_ERC20 *ERC20Session) ClearStuckBalance(amountPercentage *big.Int, adr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.ClearStuckBalance(&_ERC20.TransactOpts, amountPercentage, adr)
}

// ClearStuckBalance is a paid mutator transaction binding the contract method 0x56a227f2.
//
// Solidity: function clearStuckBalance(uint256 amountPercentage, address adr) returns()
func (_ERC20 *ERC20TransactorSession) ClearStuckBalance(amountPercentage *big.Int, adr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.ClearStuckBalance(&_ERC20.TransactOpts, amountPercentage, adr)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// EnableTransfer is a paid mutator transaction binding the contract method 0xd5938aac.
//
// Solidity: function enableTransfer(address _addr) returns()
func (_ERC20 *ERC20Transactor) EnableTransfer(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "enableTransfer", _addr)
}

// EnableTransfer is a paid mutator transaction binding the contract method 0xd5938aac.
//
// Solidity: function enableTransfer(address _addr) returns()
func (_ERC20 *ERC20Session) EnableTransfer(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.EnableTransfer(&_ERC20.TransactOpts, _addr)
}

// EnableTransfer is a paid mutator transaction binding the contract method 0xd5938aac.
//
// Solidity: function enableTransfer(address _addr) returns()
func (_ERC20 *ERC20TransactorSession) EnableTransfer(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.EnableTransfer(&_ERC20.TransactOpts, _addr)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// ManualSync is a paid mutator transaction binding the contract method 0x753d02a1.
//
// Solidity: function manualSync() returns()
func (_ERC20 *ERC20Transactor) ManualSync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "manualSync")
}

// ManualSync is a paid mutator transaction binding the contract method 0x753d02a1.
//
// Solidity: function manualSync() returns()
func (_ERC20 *ERC20Session) ManualSync() (*types.Transaction, error) {
	return _ERC20.Contract.ManualSync(&_ERC20.TransactOpts)
}

// ManualSync is a paid mutator transaction binding the contract method 0x753d02a1.
//
// Solidity: function manualSync() returns()
func (_ERC20 *ERC20TransactorSession) ManualSync() (*types.Transaction, error) {
	return _ERC20.Contract.ManualSync(&_ERC20.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_ERC20 *ERC20Transactor) Rebase(opts *bind.TransactOpts, epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "rebase", epoch, supplyDelta)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_ERC20 *ERC20Session) Rebase(epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Rebase(&_ERC20.TransactOpts, epoch, supplyDelta)
}

// Rebase is a paid mutator transaction binding the contract method 0x7a43e23f.
//
// Solidity: function rebase(uint256 epoch, int256 supplyDelta) returns(uint256)
func (_ERC20 *ERC20TransactorSession) Rebase(epoch *big.Int, supplyDelta *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Rebase(&_ERC20.TransactOpts, epoch, supplyDelta)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20.Contract.RenounceOwnership(&_ERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20 *ERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20.Contract.RenounceOwnership(&_ERC20.TransactOpts)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_ERC20 *ERC20Transactor) RescueToken(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "rescueToken", tokenAddress, tokens)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_ERC20 *ERC20Session) RescueToken(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.RescueToken(&_ERC20.TransactOpts, tokenAddress, tokens)
}

// RescueToken is a paid mutator transaction binding the contract method 0x33f3d628.
//
// Solidity: function rescueToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_ERC20 *ERC20TransactorSession) RescueToken(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.RescueToken(&_ERC20.TransactOpts, tokenAddress, tokens)
}

// SendPresale is a paid mutator transaction binding the contract method 0xd37e219d.
//
// Solidity: function sendPresale(address[] recipients, uint256[] values) returns()
func (_ERC20 *ERC20Transactor) SendPresale(opts *bind.TransactOpts, recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "sendPresale", recipients, values)
}

// SendPresale is a paid mutator transaction binding the contract method 0xd37e219d.
//
// Solidity: function sendPresale(address[] recipients, uint256[] values) returns()
func (_ERC20 *ERC20Session) SendPresale(recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SendPresale(&_ERC20.TransactOpts, recipients, values)
}

// SendPresale is a paid mutator transaction binding the contract method 0xd37e219d.
//
// Solidity: function sendPresale(address[] recipients, uint256[] values) returns()
func (_ERC20 *ERC20TransactorSession) SendPresale(recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SendPresale(&_ERC20.TransactOpts, recipients, values)
}

// SetFeeExempt is a paid mutator transaction binding the contract method 0x749796a5.
//
// Solidity: function setFeeExempt(address _addr) returns()
func (_ERC20 *ERC20Transactor) SetFeeExempt(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setFeeExempt", _addr)
}

// SetFeeExempt is a paid mutator transaction binding the contract method 0x749796a5.
//
// Solidity: function setFeeExempt(address _addr) returns()
func (_ERC20 *ERC20Session) SetFeeExempt(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetFeeExempt(&_ERC20.TransactOpts, _addr)
}

// SetFeeExempt is a paid mutator transaction binding the contract method 0x749796a5.
//
// Solidity: function setFeeExempt(address _addr) returns()
func (_ERC20 *ERC20TransactorSession) SetFeeExempt(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetFeeExempt(&_ERC20.TransactOpts, _addr)
}

// SetFeeReceivers is a paid mutator transaction binding the contract method 0x3c8e556d.
//
// Solidity: function setFeeReceivers(address _autoLiquidityReceiver, address _ecosystemFeeReceiver, address _marketingFeeReceiver, address _buyBackFeeReceiver) returns()
func (_ERC20 *ERC20Transactor) SetFeeReceivers(opts *bind.TransactOpts, _autoLiquidityReceiver common.Address, _ecosystemFeeReceiver common.Address, _marketingFeeReceiver common.Address, _buyBackFeeReceiver common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setFeeReceivers", _autoLiquidityReceiver, _ecosystemFeeReceiver, _marketingFeeReceiver, _buyBackFeeReceiver)
}

// SetFeeReceivers is a paid mutator transaction binding the contract method 0x3c8e556d.
//
// Solidity: function setFeeReceivers(address _autoLiquidityReceiver, address _ecosystemFeeReceiver, address _marketingFeeReceiver, address _buyBackFeeReceiver) returns()
func (_ERC20 *ERC20Session) SetFeeReceivers(_autoLiquidityReceiver common.Address, _ecosystemFeeReceiver common.Address, _marketingFeeReceiver common.Address, _buyBackFeeReceiver common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetFeeReceivers(&_ERC20.TransactOpts, _autoLiquidityReceiver, _ecosystemFeeReceiver, _marketingFeeReceiver, _buyBackFeeReceiver)
}

// SetFeeReceivers is a paid mutator transaction binding the contract method 0x3c8e556d.
//
// Solidity: function setFeeReceivers(address _autoLiquidityReceiver, address _ecosystemFeeReceiver, address _marketingFeeReceiver, address _buyBackFeeReceiver) returns()
func (_ERC20 *ERC20TransactorSession) SetFeeReceivers(_autoLiquidityReceiver common.Address, _ecosystemFeeReceiver common.Address, _marketingFeeReceiver common.Address, _buyBackFeeReceiver common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetFeeReceivers(&_ERC20.TransactOpts, _autoLiquidityReceiver, _ecosystemFeeReceiver, _marketingFeeReceiver, _buyBackFeeReceiver)
}

// SetFees is a paid mutator transaction binding the contract method 0x04a66b48.
//
// Solidity: function setFees(uint256 _ecosystemFee, uint256 _liquidityFee, uint256 _buyBackFee, uint256 _marketingFee, uint256 _feeDenominator) returns()
func (_ERC20 *ERC20Transactor) SetFees(opts *bind.TransactOpts, _ecosystemFee *big.Int, _liquidityFee *big.Int, _buyBackFee *big.Int, _marketingFee *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setFees", _ecosystemFee, _liquidityFee, _buyBackFee, _marketingFee, _feeDenominator)
}

// SetFees is a paid mutator transaction binding the contract method 0x04a66b48.
//
// Solidity: function setFees(uint256 _ecosystemFee, uint256 _liquidityFee, uint256 _buyBackFee, uint256 _marketingFee, uint256 _feeDenominator) returns()
func (_ERC20 *ERC20Session) SetFees(_ecosystemFee *big.Int, _liquidityFee *big.Int, _buyBackFee *big.Int, _marketingFee *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetFees(&_ERC20.TransactOpts, _ecosystemFee, _liquidityFee, _buyBackFee, _marketingFee, _feeDenominator)
}

// SetFees is a paid mutator transaction binding the contract method 0x04a66b48.
//
// Solidity: function setFees(uint256 _ecosystemFee, uint256 _liquidityFee, uint256 _buyBackFee, uint256 _marketingFee, uint256 _feeDenominator) returns()
func (_ERC20 *ERC20TransactorSession) SetFees(_ecosystemFee *big.Int, _liquidityFee *big.Int, _buyBackFee *big.Int, _marketingFee *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetFees(&_ERC20.TransactOpts, _ecosystemFee, _liquidityFee, _buyBackFee, _marketingFee, _feeDenominator)
}

// SetInitialDistributionFinished is a paid mutator transaction binding the contract method 0x2be6937d.
//
// Solidity: function setInitialDistributionFinished() returns()
func (_ERC20 *ERC20Transactor) SetInitialDistributionFinished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setInitialDistributionFinished")
}

// SetInitialDistributionFinished is a paid mutator transaction binding the contract method 0x2be6937d.
//
// Solidity: function setInitialDistributionFinished() returns()
func (_ERC20 *ERC20Session) SetInitialDistributionFinished() (*types.Transaction, error) {
	return _ERC20.Contract.SetInitialDistributionFinished(&_ERC20.TransactOpts)
}

// SetInitialDistributionFinished is a paid mutator transaction binding the contract method 0x2be6937d.
//
// Solidity: function setInitialDistributionFinished() returns()
func (_ERC20 *ERC20TransactorSession) SetInitialDistributionFinished() (*types.Transaction, error) {
	return _ERC20.Contract.SetInitialDistributionFinished(&_ERC20.TransactOpts)
}

// SetLP is a paid mutator transaction binding the contract method 0x2f34d282.
//
// Solidity: function setLP(address _address) returns()
func (_ERC20 *ERC20Transactor) SetLP(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setLP", _address)
}

// SetLP is a paid mutator transaction binding the contract method 0x2f34d282.
//
// Solidity: function setLP(address _address) returns()
func (_ERC20 *ERC20Session) SetLP(_address common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetLP(&_ERC20.TransactOpts, _address)
}

// SetLP is a paid mutator transaction binding the contract method 0x2f34d282.
//
// Solidity: function setLP(address _address) returns()
func (_ERC20 *ERC20TransactorSession) SetLP(_address common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetLP(&_ERC20.TransactOpts, _address)
}

// SetMaster is a paid mutator transaction binding the contract method 0x26fae0d3.
//
// Solidity: function setMaster(address _master) returns()
func (_ERC20 *ERC20Transactor) SetMaster(opts *bind.TransactOpts, _master common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setMaster", _master)
}

// SetMaster is a paid mutator transaction binding the contract method 0x26fae0d3.
//
// Solidity: function setMaster(address _master) returns()
func (_ERC20 *ERC20Session) SetMaster(_master common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaster(&_ERC20.TransactOpts, _master)
}

// SetMaster is a paid mutator transaction binding the contract method 0x26fae0d3.
//
// Solidity: function setMaster(address _master) returns()
func (_ERC20 *ERC20TransactorSession) SetMaster(_master common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaster(&_ERC20.TransactOpts, _master)
}

// SetMaxWalletExempt is a paid mutator transaction binding the contract method 0x2dd5efe7.
//
// Solidity: function setMaxWalletExempt(address _addr) returns()
func (_ERC20 *ERC20Transactor) SetMaxWalletExempt(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setMaxWalletExempt", _addr)
}

// SetMaxWalletExempt is a paid mutator transaction binding the contract method 0x2dd5efe7.
//
// Solidity: function setMaxWalletExempt(address _addr) returns()
func (_ERC20 *ERC20Session) SetMaxWalletExempt(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaxWalletExempt(&_ERC20.TransactOpts, _addr)
}

// SetMaxWalletExempt is a paid mutator transaction binding the contract method 0x2dd5efe7.
//
// Solidity: function setMaxWalletExempt(address _addr) returns()
func (_ERC20 *ERC20TransactorSession) SetMaxWalletExempt(_addr common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaxWalletExempt(&_ERC20.TransactOpts, _addr)
}

// SetMaxWalletToken is a paid mutator transaction binding the contract method 0xc0745ec7.
//
// Solidity: function setMaxWalletToken(uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20Transactor) SetMaxWalletToken(opts *bind.TransactOpts, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setMaxWalletToken", _num, _denom)
}

// SetMaxWalletToken is a paid mutator transaction binding the contract method 0xc0745ec7.
//
// Solidity: function setMaxWalletToken(uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20Session) SetMaxWalletToken(_num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaxWalletToken(&_ERC20.TransactOpts, _num, _denom)
}

// SetMaxWalletToken is a paid mutator transaction binding the contract method 0xc0745ec7.
//
// Solidity: function setMaxWalletToken(uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20TransactorSession) SetMaxWalletToken(_num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetMaxWalletToken(&_ERC20.TransactOpts, _num, _denom)
}

// SetSwapBackSettings is a paid mutator transaction binding the contract method 0xd0889358.
//
// Solidity: function setSwapBackSettings(bool _enabled, uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20Transactor) SetSwapBackSettings(opts *bind.TransactOpts, _enabled bool, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setSwapBackSettings", _enabled, _num, _denom)
}

// SetSwapBackSettings is a paid mutator transaction binding the contract method 0xd0889358.
//
// Solidity: function setSwapBackSettings(bool _enabled, uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20Session) SetSwapBackSettings(_enabled bool, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetSwapBackSettings(&_ERC20.TransactOpts, _enabled, _num, _denom)
}

// SetSwapBackSettings is a paid mutator transaction binding the contract method 0xd0889358.
//
// Solidity: function setSwapBackSettings(bool _enabled, uint256 _num, uint256 _denom) returns()
func (_ERC20 *ERC20TransactorSession) SetSwapBackSettings(_enabled bool, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetSwapBackSettings(&_ERC20.TransactOpts, _enabled, _num, _denom)
}

// SetTargetLiquidity is a paid mutator transaction binding the contract method 0x201e7991.
//
// Solidity: function setTargetLiquidity(uint256 target, uint256 accuracy) returns()
func (_ERC20 *ERC20Transactor) SetTargetLiquidity(opts *bind.TransactOpts, target *big.Int, accuracy *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setTargetLiquidity", target, accuracy)
}

// SetTargetLiquidity is a paid mutator transaction binding the contract method 0x201e7991.
//
// Solidity: function setTargetLiquidity(uint256 target, uint256 accuracy) returns()
func (_ERC20 *ERC20Session) SetTargetLiquidity(target *big.Int, accuracy *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetTargetLiquidity(&_ERC20.TransactOpts, target, accuracy)
}

// SetTargetLiquidity is a paid mutator transaction binding the contract method 0x201e7991.
//
// Solidity: function setTargetLiquidity(uint256 target, uint256 accuracy) returns()
func (_ERC20 *ERC20TransactorSession) SetTargetLiquidity(target *big.Int, accuracy *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.SetTargetLiquidity(&_ERC20.TransactOpts, target, accuracy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC20 *ERC20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC20 *ERC20Session) Receive() (*types.Transaction, error) {
	return _ERC20.Contract.Receive(&_ERC20.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ERC20 *ERC20TransactorSession) Receive() (*types.Transaction, error) {
	return _ERC20.Contract.Receive(&_ERC20.TransactOpts)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LogRebaseIterator is returned from FilterLogRebase and is used to iterate over the raw logs and unpacked data for LogRebase events raised by the ERC20 contract.
type ERC20LogRebaseIterator struct {
	Event *ERC20LogRebase // Event containing the contract specifics and raw log

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
func (it *ERC20LogRebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LogRebase)
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
		it.Event = new(ERC20LogRebase)
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
func (it *ERC20LogRebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LogRebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LogRebase represents a LogRebase event raised by the ERC20 contract.
type ERC20LogRebase struct {
	Epoch       *big.Int
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRebase is a free log retrieval operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_ERC20 *ERC20Filterer) FilterLogRebase(opts *bind.FilterOpts, epoch []*big.Int) (*ERC20LogRebaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LogRebaseIterator{contract: _ERC20.contract, event: "LogRebase", logs: logs, sub: sub}, nil
}

// WatchLogRebase is a free log subscription operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_ERC20 *ERC20Filterer) WatchLogRebase(opts *bind.WatchOpts, sink chan<- *ERC20LogRebase, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LogRebase)
				if err := _ERC20.contract.UnpackLog(event, "LogRebase", log); err != nil {
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

// ParseLogRebase is a log parse operation binding the contract event 0x72725a3b1e5bd622d6bcd1339bb31279c351abe8f541ac7fd320f24e1b1641f2.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 totalSupply)
func (_ERC20 *ERC20Filterer) ParseLogRebase(log types.Log) (*ERC20LogRebase, error) {
	event := new(ERC20LogRebase)
	if err := _ERC20.contract.UnpackLog(event, "LogRebase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20OwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ERC20 contract.
type ERC20OwnershipRenouncedIterator struct {
	Event *ERC20OwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ERC20OwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20OwnershipRenounced)
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
		it.Event = new(ERC20OwnershipRenounced)
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
func (it *ERC20OwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20OwnershipRenounced represents a OwnershipRenounced event raised by the ERC20 contract.
type ERC20OwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ERC20 *ERC20Filterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ERC20OwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20OwnershipRenouncedIterator{contract: _ERC20.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ERC20 *ERC20Filterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ERC20OwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OwnershipRenounced)
				if err := _ERC20.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ERC20 *ERC20Filterer) ParseOwnershipRenounced(log types.Log) (*ERC20OwnershipRenounced, error) {
	event := new(ERC20OwnershipRenounced)
	if err := _ERC20.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20 contract.
type ERC20OwnershipTransferredIterator struct {
	Event *ERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20OwnershipTransferred)
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
		it.Event = new(ERC20OwnershipTransferred)
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
func (it *ERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20OwnershipTransferred represents a OwnershipTransferred event raised by the ERC20 contract.
type ERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20OwnershipTransferredIterator{contract: _ERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OwnershipTransferred)
				if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20 *ERC20Filterer) ParseOwnershipTransferred(log types.Log) (*ERC20OwnershipTransferred, error) {
	event := new(ERC20OwnershipTransferred)
	if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
