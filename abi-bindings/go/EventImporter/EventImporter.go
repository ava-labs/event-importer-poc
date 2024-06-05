// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eventimporter

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

// EventImporterMetaData contains all meta data concerning the EventImporter contract.
var EventImporterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blockHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"receiptProof\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"importEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warpMessenger\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EventImporterABI is the input ABI used to generate the binding from.
// Deprecated: Use EventImporterMetaData.ABI instead.
var EventImporterABI = EventImporterMetaData.ABI

// EventImporter is an auto generated Go binding around an Ethereum contract.
type EventImporter struct {
	EventImporterCaller     // Read-only binding to the contract
	EventImporterTransactor // Write-only binding to the contract
	EventImporterFilterer   // Log filterer for contract events
}

// EventImporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventImporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventImporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventImporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventImporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventImporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventImporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventImporterSession struct {
	Contract     *EventImporter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventImporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventImporterCallerSession struct {
	Contract *EventImporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EventImporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventImporterTransactorSession struct {
	Contract     *EventImporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EventImporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventImporterRaw struct {
	Contract *EventImporter // Generic contract binding to access the raw methods on
}

// EventImporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventImporterCallerRaw struct {
	Contract *EventImporterCaller // Generic read-only contract binding to access the raw methods on
}

// EventImporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventImporterTransactorRaw struct {
	Contract *EventImporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventImporter creates a new instance of EventImporter, bound to a specific deployed contract.
func NewEventImporter(address common.Address, backend bind.ContractBackend) (*EventImporter, error) {
	contract, err := bindEventImporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventImporter{EventImporterCaller: EventImporterCaller{contract: contract}, EventImporterTransactor: EventImporterTransactor{contract: contract}, EventImporterFilterer: EventImporterFilterer{contract: contract}}, nil
}

// NewEventImporterCaller creates a new read-only instance of EventImporter, bound to a specific deployed contract.
func NewEventImporterCaller(address common.Address, caller bind.ContractCaller) (*EventImporterCaller, error) {
	contract, err := bindEventImporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventImporterCaller{contract: contract}, nil
}

// NewEventImporterTransactor creates a new write-only instance of EventImporter, bound to a specific deployed contract.
func NewEventImporterTransactor(address common.Address, transactor bind.ContractTransactor) (*EventImporterTransactor, error) {
	contract, err := bindEventImporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventImporterTransactor{contract: contract}, nil
}

// NewEventImporterFilterer creates a new log filterer instance of EventImporter, bound to a specific deployed contract.
func NewEventImporterFilterer(address common.Address, filterer bind.ContractFilterer) (*EventImporterFilterer, error) {
	contract, err := bindEventImporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventImporterFilterer{contract: contract}, nil
}

// bindEventImporter binds a generic wrapper to an already deployed contract.
func bindEventImporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventImporterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventImporter *EventImporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventImporter.Contract.EventImporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventImporter *EventImporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventImporter.Contract.EventImporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventImporter *EventImporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventImporter.Contract.EventImporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventImporter *EventImporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventImporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventImporter *EventImporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventImporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventImporter *EventImporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventImporter.Contract.contract.Transact(opts, method, params...)
}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_EventImporter *EventImporterCaller) WarpMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventImporter.contract.Call(opts, &out, "warpMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_EventImporter *EventImporterSession) WarpMessenger() (common.Address, error) {
	return _EventImporter.Contract.WarpMessenger(&_EventImporter.CallOpts)
}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_EventImporter *EventImporterCallerSession) WarpMessenger() (common.Address, error) {
	return _EventImporter.Contract.WarpMessenger(&_EventImporter.CallOpts)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_EventImporter *EventImporterTransactor) ImportEvent(opts *bind.TransactOpts, blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _EventImporter.contract.Transact(opts, "importEvent", blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_EventImporter *EventImporterSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _EventImporter.Contract.ImportEvent(&_EventImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_EventImporter *EventImporterTransactorSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _EventImporter.Contract.ImportEvent(&_EventImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}
