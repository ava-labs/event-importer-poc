// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockpricefeedaggregator

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

// MockPriceFeedAggregatorMetaData contains all meta data concerning the MockPriceFeedAggregator contract.
var MockPriceFeedAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundID\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentAnswer_\",\"type\":\"int256\"},{\"internalType\":\"uint80\",\"name\":\"roundID_\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"name\":\"updateAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50336080526080516102a06100366000396000818160d1015261015001526102a06000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806314f8b42414610067578063683d010d146100975780637519ab50146100ac5780637e1b4cb0146100c3578063d5f39488146100cc578063feaf968c1461010b575b600080fd5b60015461007a906001600160501b031681565b6040516001600160501b0390911681526020015b60405180910390f35b6100aa6100a5366004610226565b610145565b005b6100b560025481565b60405190815260200161008e565b6100b560005481565b6100f37f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161008e565b600154600054600254604080516001600160501b03909416808552602085019390935283018190526060830152608082015260a00161008e565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101c15760405162461bcd60e51b815260206004820152601f60248201527f4f6e6c79206465706c6f7965722063616e2075706461746520616e7377657200604482015260640160405180910390fd5b60008390556001805469ffffffffffffffffffff19166001600160501b038416908117909155600282905560405182815284907f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f9060200160405180910390a3505050565b60008060006060848603121561023b57600080fd5b8335925060208401356001600160501b038116811461025957600080fd5b92959294505050604091909101359056fea2646970667358221220d670ac97b1f226886dc490c248ae6db92e3161ebeaba40164c8a8c131233e62c64736f6c63430008120033",
}

// MockPriceFeedAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use MockPriceFeedAggregatorMetaData.ABI instead.
var MockPriceFeedAggregatorABI = MockPriceFeedAggregatorMetaData.ABI

// MockPriceFeedAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockPriceFeedAggregatorMetaData.Bin instead.
var MockPriceFeedAggregatorBin = MockPriceFeedAggregatorMetaData.Bin

// DeployMockPriceFeedAggregator deploys a new Ethereum contract, binding an instance of MockPriceFeedAggregator to it.
func DeployMockPriceFeedAggregator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockPriceFeedAggregator, error) {
	parsed, err := MockPriceFeedAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockPriceFeedAggregatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockPriceFeedAggregator{MockPriceFeedAggregatorCaller: MockPriceFeedAggregatorCaller{contract: contract}, MockPriceFeedAggregatorTransactor: MockPriceFeedAggregatorTransactor{contract: contract}, MockPriceFeedAggregatorFilterer: MockPriceFeedAggregatorFilterer{contract: contract}}, nil
}

// MockPriceFeedAggregator is an auto generated Go binding around an Ethereum contract.
type MockPriceFeedAggregator struct {
	MockPriceFeedAggregatorCaller     // Read-only binding to the contract
	MockPriceFeedAggregatorTransactor // Write-only binding to the contract
	MockPriceFeedAggregatorFilterer   // Log filterer for contract events
}

// MockPriceFeedAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockPriceFeedAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPriceFeedAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockPriceFeedAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPriceFeedAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockPriceFeedAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPriceFeedAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockPriceFeedAggregatorSession struct {
	Contract     *MockPriceFeedAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MockPriceFeedAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockPriceFeedAggregatorCallerSession struct {
	Contract *MockPriceFeedAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MockPriceFeedAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockPriceFeedAggregatorTransactorSession struct {
	Contract     *MockPriceFeedAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MockPriceFeedAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockPriceFeedAggregatorRaw struct {
	Contract *MockPriceFeedAggregator // Generic contract binding to access the raw methods on
}

// MockPriceFeedAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockPriceFeedAggregatorCallerRaw struct {
	Contract *MockPriceFeedAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// MockPriceFeedAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockPriceFeedAggregatorTransactorRaw struct {
	Contract *MockPriceFeedAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockPriceFeedAggregator creates a new instance of MockPriceFeedAggregator, bound to a specific deployed contract.
func NewMockPriceFeedAggregator(address common.Address, backend bind.ContractBackend) (*MockPriceFeedAggregator, error) {
	contract, err := bindMockPriceFeedAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockPriceFeedAggregator{MockPriceFeedAggregatorCaller: MockPriceFeedAggregatorCaller{contract: contract}, MockPriceFeedAggregatorTransactor: MockPriceFeedAggregatorTransactor{contract: contract}, MockPriceFeedAggregatorFilterer: MockPriceFeedAggregatorFilterer{contract: contract}}, nil
}

// NewMockPriceFeedAggregatorCaller creates a new read-only instance of MockPriceFeedAggregator, bound to a specific deployed contract.
func NewMockPriceFeedAggregatorCaller(address common.Address, caller bind.ContractCaller) (*MockPriceFeedAggregatorCaller, error) {
	contract, err := bindMockPriceFeedAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockPriceFeedAggregatorCaller{contract: contract}, nil
}

// NewMockPriceFeedAggregatorTransactor creates a new write-only instance of MockPriceFeedAggregator, bound to a specific deployed contract.
func NewMockPriceFeedAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MockPriceFeedAggregatorTransactor, error) {
	contract, err := bindMockPriceFeedAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockPriceFeedAggregatorTransactor{contract: contract}, nil
}

// NewMockPriceFeedAggregatorFilterer creates a new log filterer instance of MockPriceFeedAggregator, bound to a specific deployed contract.
func NewMockPriceFeedAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MockPriceFeedAggregatorFilterer, error) {
	contract, err := bindMockPriceFeedAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockPriceFeedAggregatorFilterer{contract: contract}, nil
}

// bindMockPriceFeedAggregator binds a generic wrapper to an already deployed contract.
func bindMockPriceFeedAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockPriceFeedAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPriceFeedAggregator.Contract.MockPriceFeedAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.MockPriceFeedAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.MockPriceFeedAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPriceFeedAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.contract.Transact(opts, method, params...)
}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCaller) CurrentAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPriceFeedAggregator.contract.Call(opts, &out, "currentAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) CurrentAnswer() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.CurrentAnswer(&_MockPriceFeedAggregator.CallOpts)
}

// CurrentAnswer is a free data retrieval call binding the contract method 0x7e1b4cb0.
//
// Solidity: function currentAnswer() view returns(int256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerSession) CurrentAnswer() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.CurrentAnswer(&_MockPriceFeedAggregator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockPriceFeedAggregator.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) Deployer() (common.Address, error) {
	return _MockPriceFeedAggregator.Contract.Deployer(&_MockPriceFeedAggregator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerSession) Deployer() (common.Address, error) {
	return _MockPriceFeedAggregator.Contract.Deployer(&_MockPriceFeedAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MockPriceFeedAggregator.contract.Call(opts, &out, "latestRoundData")

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
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MockPriceFeedAggregator.Contract.LatestRoundData(&_MockPriceFeedAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MockPriceFeedAggregator.Contract.LatestRoundData(&_MockPriceFeedAggregator.CallOpts)
}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCaller) RoundID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPriceFeedAggregator.contract.Call(opts, &out, "roundID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) RoundID() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.RoundID(&_MockPriceFeedAggregator.CallOpts)
}

// RoundID is a free data retrieval call binding the contract method 0x14f8b424.
//
// Solidity: function roundID() view returns(uint80)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerSession) RoundID() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.RoundID(&_MockPriceFeedAggregator.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCaller) UpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockPriceFeedAggregator.contract.Call(opts, &out, "updatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) UpdatedAt() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.UpdatedAt(&_MockPriceFeedAggregator.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorCallerSession) UpdatedAt() (*big.Int, error) {
	return _MockPriceFeedAggregator.Contract.UpdatedAt(&_MockPriceFeedAggregator.CallOpts)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0x683d010d.
//
// Solidity: function updateAnswer(int256 currentAnswer_, uint80 roundID_, uint256 updatedAt_) returns()
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorTransactor) UpdateAnswer(opts *bind.TransactOpts, currentAnswer_ *big.Int, roundID_ *big.Int, updatedAt_ *big.Int) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.contract.Transact(opts, "updateAnswer", currentAnswer_, roundID_, updatedAt_)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0x683d010d.
//
// Solidity: function updateAnswer(int256 currentAnswer_, uint80 roundID_, uint256 updatedAt_) returns()
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorSession) UpdateAnswer(currentAnswer_ *big.Int, roundID_ *big.Int, updatedAt_ *big.Int) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.UpdateAnswer(&_MockPriceFeedAggregator.TransactOpts, currentAnswer_, roundID_, updatedAt_)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0x683d010d.
//
// Solidity: function updateAnswer(int256 currentAnswer_, uint80 roundID_, uint256 updatedAt_) returns()
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorTransactorSession) UpdateAnswer(currentAnswer_ *big.Int, roundID_ *big.Int, updatedAt_ *big.Int) (*types.Transaction, error) {
	return _MockPriceFeedAggregator.Contract.UpdateAnswer(&_MockPriceFeedAggregator.TransactOpts, currentAnswer_, roundID_, updatedAt_)
}

// MockPriceFeedAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the MockPriceFeedAggregator contract.
type MockPriceFeedAggregatorAnswerUpdatedIterator struct {
	Event *MockPriceFeedAggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *MockPriceFeedAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockPriceFeedAggregatorAnswerUpdated)
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
		it.Event = new(MockPriceFeedAggregatorAnswerUpdated)
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
func (it *MockPriceFeedAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockPriceFeedAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockPriceFeedAggregatorAnswerUpdated represents a AnswerUpdated event raised by the MockPriceFeedAggregator contract.
type MockPriceFeedAggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*MockPriceFeedAggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockPriceFeedAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &MockPriceFeedAggregatorAnswerUpdatedIterator{contract: _MockPriceFeedAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *MockPriceFeedAggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockPriceFeedAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockPriceFeedAggregatorAnswerUpdated)
				if err := _MockPriceFeedAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_MockPriceFeedAggregator *MockPriceFeedAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*MockPriceFeedAggregatorAnswerUpdated, error) {
	event := new(MockPriceFeedAggregatorAnswerUpdated)
	if err := _MockPriceFeedAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
