// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RentalRecordsABI is the input ABI used to generate the binding from.
const RentalRecordsABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"paramLocator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"paramRenter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"paramAddressHome\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"paramRentalValue\",\"type\":\"uint256\"}],\"name\":\"registerRental\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"locator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"renter\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"addressHome\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rentalValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RentalRecords is an auto generated Go binding around an Ethereum contract.
type RentalRecords struct {
	RentalRecordsCaller     // Read-only binding to the contract
	RentalRecordsTransactor // Write-only binding to the contract
	RentalRecordsFilterer   // Log filterer for contract events
}

// RentalRecordsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RentalRecordsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalRecordsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RentalRecordsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalRecordsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RentalRecordsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalRecordsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RentalRecordsSession struct {
	Contract     *RentalRecords    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RentalRecordsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RentalRecordsCallerSession struct {
	Contract *RentalRecordsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RentalRecordsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RentalRecordsTransactorSession struct {
	Contract     *RentalRecordsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RentalRecordsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RentalRecordsRaw struct {
	Contract *RentalRecords // Generic contract binding to access the raw methods on
}

// RentalRecordsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RentalRecordsCallerRaw struct {
	Contract *RentalRecordsCaller // Generic read-only contract binding to access the raw methods on
}

// RentalRecordsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RentalRecordsTransactorRaw struct {
	Contract *RentalRecordsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRentalRecords creates a new instance of RentalRecords, bound to a specific deployed contract.
func NewRentalRecords(address common.Address, backend bind.ContractBackend) (*RentalRecords, error) {
	contract, err := bindRentalRecords(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RentalRecords{RentalRecordsCaller: RentalRecordsCaller{contract: contract}, RentalRecordsTransactor: RentalRecordsTransactor{contract: contract}, RentalRecordsFilterer: RentalRecordsFilterer{contract: contract}}, nil
}

// NewRentalRecordsCaller creates a new read-only instance of RentalRecords, bound to a specific deployed contract.
func NewRentalRecordsCaller(address common.Address, caller bind.ContractCaller) (*RentalRecordsCaller, error) {
	contract, err := bindRentalRecords(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RentalRecordsCaller{contract: contract}, nil
}

// NewRentalRecordsTransactor creates a new write-only instance of RentalRecords, bound to a specific deployed contract.
func NewRentalRecordsTransactor(address common.Address, transactor bind.ContractTransactor) (*RentalRecordsTransactor, error) {
	contract, err := bindRentalRecords(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RentalRecordsTransactor{contract: contract}, nil
}

// NewRentalRecordsFilterer creates a new log filterer instance of RentalRecords, bound to a specific deployed contract.
func NewRentalRecordsFilterer(address common.Address, filterer bind.ContractFilterer) (*RentalRecordsFilterer, error) {
	contract, err := bindRentalRecords(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RentalRecordsFilterer{contract: contract}, nil
}

// bindRentalRecords binds a generic wrapper to an already deployed contract.
func bindRentalRecords(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RentalRecordsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalRecords *RentalRecordsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalRecords.Contract.RentalRecordsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalRecords *RentalRecordsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalRecords.Contract.RentalRecordsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalRecords *RentalRecordsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalRecords.Contract.RentalRecordsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalRecords *RentalRecordsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalRecords.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalRecords *RentalRecordsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalRecords.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalRecords *RentalRecordsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalRecords.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RentalRecords *RentalRecordsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RentalRecords.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RentalRecords *RentalRecordsSession) Owner() (common.Address, error) {
	return _RentalRecords.Contract.Owner(&_RentalRecords.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RentalRecords *RentalRecordsCallerSession) Owner() (common.Address, error) {
	return _RentalRecords.Contract.Owner(&_RentalRecords.CallOpts)
}

// Rentals is a free data retrieval call binding the contract method 0xaf94a1d1.
//
// Solidity: function rentals(uint256 ) view returns(string locator, string renter, string addressHome, uint256 rentalValue)
func (_RentalRecords *RentalRecordsCaller) Rentals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Locator     string
	Renter      string
	AddressHome string
	RentalValue *big.Int
}, error) {
	var out []interface{}
	err := _RentalRecords.contract.Call(opts, &out, "rentals", arg0)

	outstruct := new(struct {
		Locator     string
		Renter      string
		AddressHome string
		RentalValue *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Locator = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Renter = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AddressHome = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.RentalValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rentals is a free data retrieval call binding the contract method 0xaf94a1d1.
//
// Solidity: function rentals(uint256 ) view returns(string locator, string renter, string addressHome, uint256 rentalValue)
func (_RentalRecords *RentalRecordsSession) Rentals(arg0 *big.Int) (struct {
	Locator     string
	Renter      string
	AddressHome string
	RentalValue *big.Int
}, error) {
	return _RentalRecords.Contract.Rentals(&_RentalRecords.CallOpts, arg0)
}

// Rentals is a free data retrieval call binding the contract method 0xaf94a1d1.
//
// Solidity: function rentals(uint256 ) view returns(string locator, string renter, string addressHome, uint256 rentalValue)
func (_RentalRecords *RentalRecordsCallerSession) Rentals(arg0 *big.Int) (struct {
	Locator     string
	Renter      string
	AddressHome string
	RentalValue *big.Int
}, error) {
	return _RentalRecords.Contract.Rentals(&_RentalRecords.CallOpts, arg0)
}

// RegisterRental is a paid mutator transaction binding the contract method 0xc70a1c01.
//
// Solidity: function registerRental(string paramLocator, string paramRenter, string paramAddressHome, uint256 paramRentalValue) returns(bool)
func (_RentalRecords *RentalRecordsTransactor) RegisterRental(opts *bind.TransactOpts, paramLocator string, paramRenter string, paramAddressHome string, paramRentalValue *big.Int) (*types.Transaction, error) {
	return _RentalRecords.contract.Transact(opts, "registerRental", paramLocator, paramRenter, paramAddressHome, paramRentalValue)
}

// RegisterRental is a paid mutator transaction binding the contract method 0xc70a1c01.
//
// Solidity: function registerRental(string paramLocator, string paramRenter, string paramAddressHome, uint256 paramRentalValue) returns(bool)
func (_RentalRecords *RentalRecordsSession) RegisterRental(paramLocator string, paramRenter string, paramAddressHome string, paramRentalValue *big.Int) (*types.Transaction, error) {
	return _RentalRecords.Contract.RegisterRental(&_RentalRecords.TransactOpts, paramLocator, paramRenter, paramAddressHome, paramRentalValue)
}

// RegisterRental is a paid mutator transaction binding the contract method 0xc70a1c01.
//
// Solidity: function registerRental(string paramLocator, string paramRenter, string paramAddressHome, uint256 paramRentalValue) returns(bool)
func (_RentalRecords *RentalRecordsTransactorSession) RegisterRental(paramLocator string, paramRenter string, paramAddressHome string, paramRentalValue *big.Int) (*types.Transaction, error) {
	return _RentalRecords.Contract.RegisterRental(&_RentalRecords.TransactOpts, paramLocator, paramRenter, paramAddressHome, paramRentalValue)
}
