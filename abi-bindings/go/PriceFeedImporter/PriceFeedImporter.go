// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricefeedimporter

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

// PriceFeedImporterMetaData contains all meta data concerning the PriceFeedImporter contract.
var PriceFeedImporterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceOracleAggregator_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"currentAnswer\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"roundID\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ANSWER_UPDATED_EVENT_SIGNATURE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"blockHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"receiptProof\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"importEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSourceBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSourceLogIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSourceTxIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundID\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceOracleAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warpMessenger\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceFeedImporterABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedImporterMetaData.ABI instead.
var PriceFeedImporterABI = PriceFeedImporterMetaData.ABI

// PriceFeedImporter is an auto generated Go binding around an Ethereum contract.
type PriceFeedImporter struct {
	PriceFeedImporterCaller     // Read-only binding to the contract
	PriceFeedImporterTransactor // Write-only binding to the contract
	PriceFeedImporterFilterer   // Log filterer for contract events
}

// PriceFeedImporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedImporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedImporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedImporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedImporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedImporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedImporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedImporterSession struct {
	Contract     *PriceFeedImporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceFeedImporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedImporterCallerSession struct {
	Contract *PriceFeedImporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceFeedImporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedImporterTransactorSession struct {
	Contract     *PriceFeedImporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceFeedImporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedImporterRaw struct {
	Contract *PriceFeedImporter // Generic contract binding to access the raw methods on
}

// PriceFeedImporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedImporterCallerRaw struct {
	Contract *PriceFeedImporterCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedImporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedImporterTransactorRaw struct {
	Contract *PriceFeedImporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedImporter creates a new instance of PriceFeedImporter, bound to a specific deployed contract.
func NewPriceFeedImporter(address common.Address, backend bind.ContractBackend) (*PriceFeedImporter, error) {
	contract, err := bindPriceFeedImporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedImporter{PriceFeedImporterCaller: PriceFeedImporterCaller{contract: contract}, PriceFeedImporterTransactor: PriceFeedImporterTransactor{contract: contract}, PriceFeedImporterFilterer: PriceFeedImporterFilterer{contract: contract}}, nil
}

// NewPriceFeedImporterCaller creates a new read-only instance of PriceFeedImporter, bound to a specific deployed contract.
func NewPriceFeedImporterCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedImporterCaller, error) {
	contract, err := bindPriceFeedImporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedImporterCaller{contract: contract}, nil
}

// NewPriceFeedImporterTransactor creates a new write-only instance of PriceFeedImporter, bound to a specific deployed contract.
func NewPriceFeedImporterTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedImporterTransactor, error) {
	contract, err := bindPriceFeedImporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedImporterTransactor{contract: contract}, nil
}

// NewPriceFeedImporterFilterer creates a new log filterer instance of PriceFeedImporter, bound to a specific deployed contract.
func NewPriceFeedImporterFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedImporterFilterer, error) {
	contract, err := bindPriceFeedImporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedImporterFilterer{contract: contract}, nil
}

// bindPriceFeedImporter binds a generic wrapper to an already deployed contract.
func bindPriceFeedImporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedImporterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedImporter *PriceFeedImporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedImporter.Contract.PriceFeedImporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedImporter *PriceFeedImporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.PriceFeedImporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedImporter *PriceFeedImporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.PriceFeedImporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedImporter *PriceFeedImporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedImporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedImporter *PriceFeedImporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedImporter *PriceFeedImporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.contract.Transact(opts, method, params...)
}

// ANSWERUPDATEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x042a7c97.
//
// Solidity: function ANSWER_UPDATED_EVENT_SIGNATURE() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterCaller) ANSWERUPDATEDEVENTSIGNATURE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "ANSWER_UPDATED_EVENT_SIGNATURE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ANSWERUPDATEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x042a7c97.
//
// Solidity: function ANSWER_UPDATED_EVENT_SIGNATURE() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterSession) ANSWERUPDATEDEVENTSIGNATURE() ([32]byte, error) {
	return _PriceFeedImporter.Contract.ANSWERUPDATEDEVENTSIGNATURE(&_PriceFeedImporter.CallOpts)
}

// ANSWERUPDATEDEVENTSIGNATURE is a free data retrieval call binding the contract method 0x042a7c97.
//
// Solidity: function ANSWER_UPDATED_EVENT_SIGNATURE() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) ANSWERUPDATEDEVENTSIGNATURE() ([32]byte, error) {
	return _PriceFeedImporter.Contract.ANSWERUPDATEDEVENTSIGNATURE(&_PriceFeedImporter.CallOpts)
}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_PriceFeedImporter *PriceFeedImporterCaller) CurrentAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "currentAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_PriceFeedImporter *PriceFeedImporterSession) CurrentAnswer() (*big.Int, error) {
	return _PriceFeedImporter.Contract.CurrentAnswer(&_PriceFeedImporter.CallOpts)
}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) CurrentAnswer() (*big.Int, error) {
	return _PriceFeedImporter.Contract.CurrentAnswer(&_PriceFeedImporter.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_PriceFeedImporter *PriceFeedImporterCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_PriceFeedImporter *PriceFeedImporterSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PriceFeedImporter.Contract.LatestRoundData(&_PriceFeedImporter.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PriceFeedImporter.Contract.LatestRoundData(&_PriceFeedImporter.CallOpts)
}

// LatestSourceBlockNumber is a free data retrieval call binding the contract method 0x8716cde5.
//
// Solidity: function latestSourceBlockNumber() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCaller) LatestSourceBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "latestSourceBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSourceBlockNumber is a free data retrieval call binding the contract method 0x8716cde5.
//
// Solidity: function latestSourceBlockNumber() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterSession) LatestSourceBlockNumber() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceBlockNumber(&_PriceFeedImporter.CallOpts)
}

// LatestSourceBlockNumber is a free data retrieval call binding the contract method 0x8716cde5.
//
// Solidity: function latestSourceBlockNumber() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) LatestSourceBlockNumber() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceBlockNumber(&_PriceFeedImporter.CallOpts)
}

// LatestSourceLogIndex is a free data retrieval call binding the contract method 0x08fd4323.
//
// Solidity: function latestSourceLogIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCaller) LatestSourceLogIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "latestSourceLogIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSourceLogIndex is a free data retrieval call binding the contract method 0x08fd4323.
//
// Solidity: function latestSourceLogIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterSession) LatestSourceLogIndex() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceLogIndex(&_PriceFeedImporter.CallOpts)
}

// LatestSourceLogIndex is a free data retrieval call binding the contract method 0x08fd4323.
//
// Solidity: function latestSourceLogIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) LatestSourceLogIndex() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceLogIndex(&_PriceFeedImporter.CallOpts)
}

// LatestSourceTxIndex is a free data retrieval call binding the contract method 0x4af37a11.
//
// Solidity: function latestSourceTxIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCaller) LatestSourceTxIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "latestSourceTxIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSourceTxIndex is a free data retrieval call binding the contract method 0x4af37a11.
//
// Solidity: function latestSourceTxIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterSession) LatestSourceTxIndex() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceTxIndex(&_PriceFeedImporter.CallOpts)
}

// LatestSourceTxIndex is a free data retrieval call binding the contract method 0x4af37a11.
//
// Solidity: function latestSourceTxIndex() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) LatestSourceTxIndex() (*big.Int, error) {
	return _PriceFeedImporter.Contract.LatestSourceTxIndex(&_PriceFeedImporter.CallOpts)
}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_PriceFeedImporter *PriceFeedImporterCaller) RoundID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "roundID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_PriceFeedImporter *PriceFeedImporterSession) RoundID() (*big.Int, error) {
	return _PriceFeedImporter.Contract.RoundID(&_PriceFeedImporter.CallOpts)
}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) RoundID() (*big.Int, error) {
	return _PriceFeedImporter.Contract.RoundID(&_PriceFeedImporter.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterSession) SourceBlockchainID() ([32]byte, error) {
	return _PriceFeedImporter.Contract.SourceBlockchainID(&_PriceFeedImporter.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _PriceFeedImporter.Contract.SourceBlockchainID(&_PriceFeedImporter.CallOpts)
}

// SourceOracleAggregator is a free data retrieval call binding the contract method 0xce2723fc.
//
// Solidity: function sourceOracleAggregator() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterCaller) SourceOracleAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "sourceOracleAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceOracleAggregator is a free data retrieval call binding the contract method 0xce2723fc.
//
// Solidity: function sourceOracleAggregator() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterSession) SourceOracleAggregator() (common.Address, error) {
	return _PriceFeedImporter.Contract.SourceOracleAggregator(&_PriceFeedImporter.CallOpts)
}

// SourceOracleAggregator is a free data retrieval call binding the contract method 0xce2723fc.
//
// Solidity: function sourceOracleAggregator() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) SourceOracleAggregator() (common.Address, error) {
	return _PriceFeedImporter.Contract.SourceOracleAggregator(&_PriceFeedImporter.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCaller) UpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "updatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterSession) UpdatedAt() (*big.Int, error) {
	return _PriceFeedImporter.Contract.UpdatedAt(&_PriceFeedImporter.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) UpdatedAt() (*big.Int, error) {
	return _PriceFeedImporter.Contract.UpdatedAt(&_PriceFeedImporter.CallOpts)
}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterCaller) WarpMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedImporter.contract.Call(opts, &out, "warpMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterSession) WarpMessenger() (common.Address, error) {
	return _PriceFeedImporter.Contract.WarpMessenger(&_PriceFeedImporter.CallOpts)
}

// WarpMessenger is a free data retrieval call binding the contract method 0xc9572e14.
//
// Solidity: function warpMessenger() view returns(address)
func (_PriceFeedImporter *PriceFeedImporterCallerSession) WarpMessenger() (common.Address, error) {
	return _PriceFeedImporter.Contract.WarpMessenger(&_PriceFeedImporter.CallOpts)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_PriceFeedImporter *PriceFeedImporterTransactor) ImportEvent(opts *bind.TransactOpts, blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _PriceFeedImporter.contract.Transact(opts, "importEvent", blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_PriceFeedImporter *PriceFeedImporterSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.ImportEvent(&_PriceFeedImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}

// ImportEvent is a paid mutator transaction binding the contract method 0xa47b555e.
//
// Solidity: function importEvent(bytes blockHeader, uint256 txIndex, bytes[] receiptProof, uint256 logIndex) returns()
func (_PriceFeedImporter *PriceFeedImporterTransactorSession) ImportEvent(blockHeader []byte, txIndex *big.Int, receiptProof [][]byte, logIndex *big.Int) (*types.Transaction, error) {
	return _PriceFeedImporter.Contract.ImportEvent(&_PriceFeedImporter.TransactOpts, blockHeader, txIndex, receiptProof, logIndex)
}

// PriceFeedImporterAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the PriceFeedImporter contract.
type PriceFeedImporterAnswerUpdatedIterator struct {
	Event *PriceFeedImporterAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedImporterAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedImporterAnswerUpdated)
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
		it.Event = new(PriceFeedImporterAnswerUpdated)
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
func (it *PriceFeedImporterAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedImporterAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedImporterAnswerUpdated represents a AnswerUpdated event raised by the PriceFeedImporter contract.
type PriceFeedImporterAnswerUpdated struct {
	CurrentAnswer *big.Int
	RoundID       *big.Int
	UpdatedAt     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x5b196ccf306f345de8745dffeaf185f4cafe74a334e2b2466d04c880071533b9.
//
// Solidity: event AnswerUpdated(int256 currentAnswer, uint80 roundID, uint256 updatedAt)
func (_PriceFeedImporter *PriceFeedImporterFilterer) FilterAnswerUpdated(opts *bind.FilterOpts) (*PriceFeedImporterAnswerUpdatedIterator, error) {

	logs, sub, err := _PriceFeedImporter.contract.FilterLogs(opts, "AnswerUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceFeedImporterAnswerUpdatedIterator{contract: _PriceFeedImporter.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x5b196ccf306f345de8745dffeaf185f4cafe74a334e2b2466d04c880071533b9.
//
// Solidity: event AnswerUpdated(int256 currentAnswer, uint80 roundID, uint256 updatedAt)
func (_PriceFeedImporter *PriceFeedImporterFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedImporterAnswerUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceFeedImporter.contract.WatchLogs(opts, "AnswerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedImporterAnswerUpdated)
				if err := _PriceFeedImporter.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x5b196ccf306f345de8745dffeaf185f4cafe74a334e2b2466d04c880071533b9.
//
// Solidity: event AnswerUpdated(int256 currentAnswer, uint80 roundID, uint256 updatedAt)
func (_PriceFeedImporter *PriceFeedImporterFilterer) ParseAnswerUpdated(log types.Log) (*PriceFeedImporterAnswerUpdated, error) {
	event := new(PriceFeedImporterAnswerUpdated)
	if err := _PriceFeedImporter.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
