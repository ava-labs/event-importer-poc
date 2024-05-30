// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ieventimporter

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IEventImporterMetaData contains all meta data concerning the IEventImporter contract.
var IEventImporterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blockHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"receiptProof\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"importEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEventImporterABI is the input ABI used to generate the binding from.
// Deprecated: Use IEventImporterMetaData.ABI instead.
var IEventImporterABI = IEventImporterMetaData.ABI

// IEventImporter is an auto generated Go binding around an Ethereum contract.
type IEventImporter struct {
	IEventImporterCaller     // Read-only binding to the contract
	IEventImporterTransactor // Write-only binding to the contract
	IEventImporterFilterer   // Log filterer for contract events
}

// IEventImporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEventImporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventImporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEventImporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventImporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEventImporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEventImporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEventImporterSession struct {
	Contract     *IEventImporter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEventImporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEventImporterCallerSession struct {
	Contract *IEventImporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IEventImporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEventImporterTransactorSession struct {
	Contract     *IEventImporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IEventImporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEventImporterRaw struct {
	Contract *IEventImporter // Generic contract binding to access the raw methods on
}

// IEventImporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEventImporterCallerRaw struct {
	Contract *IEventImporterCaller // Generic read-only contract binding to access the raw methods on
}

// IEventImporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEventImporterTransactorRaw struct {
	Contract *IEventImporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEventImporter creates a new instance of IEventImporter, bound to a specific deployed contract.
func NewIEventImporter(address common.Address, backend bind.ContractBackend) (*IEventImporter, error) {
	contract, err := bindIEventImporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEventImporter{IEventImporterCaller: IEventImporterCaller{contract: contract}, IEventImporterTransactor: IEventImporterTransactor{contract: contract}, IEventImporterFilterer: IEventImporterFilterer{contract: contract}}, nil
}

// NewIEventImporterCaller creates a new read-only instance of IEventImporter, bound to a specific deployed contract.
func NewIEventImporterCaller(address common.Address, caller bind.ContractCaller) (*IEventImporterCaller, error) {
	contract, err := bindIEventImporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEventImporterCaller{contract: contract}, nil
}

// NewIEventImporterTransactor creates a new write-only instance of IEventImporter, bound to a specific deployed contract.
func NewIEventImporterTransactor(address common.Address, transactor bind.ContractTransactor) (*IEventImporterTransactor, error) {
	contract, err := bindIEventImporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEventImporterTransactor{contract: contract}, nil
}

// NewIEventImporterFilterer creates a new log filterer instance of IEventImporter, bound to a specific deployed contract.
func NewIEventImporterFilterer(address common.Address, filterer bind.ContractFilterer) (*IEventImporterFilterer, error) {
	contract, err := bindIEventImporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEventImporterFilterer{contract: contract}, nil
}

// bindIEventImporter binds a generic wrapper to an already deployed contract.
func bindIEventImporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEventImporterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEventImporter *IEventImporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEventImporter.Contract.IEventImporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEventImporter *IEventImporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEventImporter.Contract.IEventImporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEventImporter *IEventImporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEventImporter.Contract.IEventImporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEventImporter *IEventImporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEventImporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEventImporter *IEventImporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEventImporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEventImporter *IEventImporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEventImporter.Contract.contract.Transact(opts, method, params...)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_IEventImporter *IEventImporterTransactor) ImportEvent(opts *bind.TransactOpts, blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _IEventImporter.contract.Transact(opts, "importEvent", blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_IEventImporter *IEventImporterSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _IEventImporter.Contract.ImportEvent(&_IEventImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_IEventImporter *IEventImporterTransactorSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _IEventImporter.Contract.ImportEvent(&_IEventImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}
