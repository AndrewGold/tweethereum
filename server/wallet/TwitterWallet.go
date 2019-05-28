// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferEtherToHandle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"etherBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"handle\",\"type\":\"string\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"fundHandleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferTokenToHandle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferTokenToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"handle\",\"type\":\"string\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferEtherToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"fundHandleEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// EtherBalance is a free data retrieval call binding the contract method 0x57a32534.
//
// Solidity: function etherBalance(string handle) constant returns(uint256)
func (_Wallet *WalletCaller) EtherBalance(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "etherBalance", handle)
	return *ret0, err
}

// EtherBalance is a free data retrieval call binding the contract method 0x57a32534.
//
// Solidity: function etherBalance(string handle) constant returns(uint256)
func (_Wallet *WalletSession) EtherBalance(handle string) (*big.Int, error) {
	return _Wallet.Contract.EtherBalance(&_Wallet.CallOpts, handle)
}

// EtherBalance is a free data retrieval call binding the contract method 0x57a32534.
//
// Solidity: function etherBalance(string handle) constant returns(uint256)
func (_Wallet *WalletCallerSession) EtherBalance(handle string) (*big.Int, error) {
	return _Wallet.Contract.EtherBalance(&_Wallet.CallOpts, handle)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Wallet *WalletCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Wallet *WalletSession) IsOwner() (bool, error) {
	return _Wallet.Contract.IsOwner(&_Wallet.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Wallet *WalletCallerSession) IsOwner() (bool, error) {
	return _Wallet.Contract.IsOwner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0xca224313.
//
// Solidity: function tokenBalance(string handle, address token) constant returns(uint256)
func (_Wallet *WalletCaller) TokenBalance(opts *bind.CallOpts, handle string, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "tokenBalance", handle, token)
	return *ret0, err
}

// TokenBalance is a free data retrieval call binding the contract method 0xca224313.
//
// Solidity: function tokenBalance(string handle, address token) constant returns(uint256)
func (_Wallet *WalletSession) TokenBalance(handle string, token common.Address) (*big.Int, error) {
	return _Wallet.Contract.TokenBalance(&_Wallet.CallOpts, handle, token)
}

// TokenBalance is a free data retrieval call binding the contract method 0xca224313.
//
// Solidity: function tokenBalance(string handle, address token) constant returns(uint256)
func (_Wallet *WalletCallerSession) TokenBalance(handle string, token common.Address) (*big.Int, error) {
	return _Wallet.Contract.TokenBalance(&_Wallet.CallOpts, handle, token)
}

// FundHandleEther is a paid mutator transaction binding the contract method 0xffc4582b.
//
// Solidity: function fundHandleEther(string handle) returns(bool)
func (_Wallet *WalletTransactor) FundHandleEther(opts *bind.TransactOpts, handle string) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "fundHandleEther", handle)
}

// FundHandleEther is a paid mutator transaction binding the contract method 0xffc4582b.
//
// Solidity: function fundHandleEther(string handle) returns(bool)
func (_Wallet *WalletSession) FundHandleEther(handle string) (*types.Transaction, error) {
	return _Wallet.Contract.FundHandleEther(&_Wallet.TransactOpts, handle)
}

// FundHandleEther is a paid mutator transaction binding the contract method 0xffc4582b.
//
// Solidity: function fundHandleEther(string handle) returns(bool)
func (_Wallet *WalletTransactorSession) FundHandleEther(handle string) (*types.Transaction, error) {
	return _Wallet.Contract.FundHandleEther(&_Wallet.TransactOpts, handle)
}

// FundHandleToken is a paid mutator transaction binding the contract method 0x7d35da3b.
//
// Solidity: function fundHandleToken(string handle, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactor) FundHandleToken(opts *bind.TransactOpts, handle string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "fundHandleToken", handle, token, value)
}

// FundHandleToken is a paid mutator transaction binding the contract method 0x7d35da3b.
//
// Solidity: function fundHandleToken(string handle, address token, uint256 value) returns(bool)
func (_Wallet *WalletSession) FundHandleToken(handle string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.FundHandleToken(&_Wallet.TransactOpts, handle, token, value)
}

// FundHandleToken is a paid mutator transaction binding the contract method 0x7d35da3b.
//
// Solidity: function fundHandleToken(string handle, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactorSession) FundHandleToken(handle string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.FundHandleToken(&_Wallet.TransactOpts, handle, token, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.RenounceOwnership(&_Wallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wallet *WalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wallet.Contract.RenounceOwnership(&_Wallet.TransactOpts)
}

// TransferEtherToAddress is a paid mutator transaction binding the contract method 0xd15ee728.
//
// Solidity: function transferEtherToAddress(string from, address to, uint256 value) returns(bool)
func (_Wallet *WalletTransactor) TransferEtherToAddress(opts *bind.TransactOpts, from string, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferEtherToAddress", from, to, value)
}

// TransferEtherToAddress is a paid mutator transaction binding the contract method 0xd15ee728.
//
// Solidity: function transferEtherToAddress(string from, address to, uint256 value) returns(bool)
func (_Wallet *WalletSession) TransferEtherToAddress(from string, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEtherToAddress(&_Wallet.TransactOpts, from, to, value)
}

// TransferEtherToAddress is a paid mutator transaction binding the contract method 0xd15ee728.
//
// Solidity: function transferEtherToAddress(string from, address to, uint256 value) returns(bool)
func (_Wallet *WalletTransactorSession) TransferEtherToAddress(from string, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEtherToAddress(&_Wallet.TransactOpts, from, to, value)
}

// TransferEtherToHandle is a paid mutator transaction binding the contract method 0x0068c72f.
//
// Solidity: function transferEtherToHandle(string from, string to, uint256 value) returns(bool)
func (_Wallet *WalletTransactor) TransferEtherToHandle(opts *bind.TransactOpts, from string, to string, value *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferEtherToHandle", from, to, value)
}

// TransferEtherToHandle is a paid mutator transaction binding the contract method 0x0068c72f.
//
// Solidity: function transferEtherToHandle(string from, string to, uint256 value) returns(bool)
func (_Wallet *WalletSession) TransferEtherToHandle(from string, to string, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEtherToHandle(&_Wallet.TransactOpts, from, to, value)
}

// TransferEtherToHandle is a paid mutator transaction binding the contract method 0x0068c72f.
//
// Solidity: function transferEtherToHandle(string from, string to, uint256 value) returns(bool)
func (_Wallet *WalletTransactorSession) TransferEtherToHandle(from string, to string, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEtherToHandle(&_Wallet.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wallet *WalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wallet *WalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wallet *WalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.TransferOwnership(&_Wallet.TransactOpts, newOwner)
}

// TransferTokenToAddress is a paid mutator transaction binding the contract method 0xc5cb811f.
//
// Solidity: function transferTokenToAddress(string from, address to, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactor) TransferTokenToAddress(opts *bind.TransactOpts, from string, to common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferTokenToAddress", from, to, token, value)
}

// TransferTokenToAddress is a paid mutator transaction binding the contract method 0xc5cb811f.
//
// Solidity: function transferTokenToAddress(string from, address to, address token, uint256 value) returns(bool)
func (_Wallet *WalletSession) TransferTokenToAddress(from string, to common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferTokenToAddress(&_Wallet.TransactOpts, from, to, token, value)
}

// TransferTokenToAddress is a paid mutator transaction binding the contract method 0xc5cb811f.
//
// Solidity: function transferTokenToAddress(string from, address to, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactorSession) TransferTokenToAddress(from string, to common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferTokenToAddress(&_Wallet.TransactOpts, from, to, token, value)
}

// TransferTokenToHandle is a paid mutator transaction binding the contract method 0xa5b7606b.
//
// Solidity: function transferTokenToHandle(string from, string to, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactor) TransferTokenToHandle(opts *bind.TransactOpts, from string, to string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferTokenToHandle", from, to, token, value)
}

// TransferTokenToHandle is a paid mutator transaction binding the contract method 0xa5b7606b.
//
// Solidity: function transferTokenToHandle(string from, string to, address token, uint256 value) returns(bool)
func (_Wallet *WalletSession) TransferTokenToHandle(from string, to string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferTokenToHandle(&_Wallet.TransactOpts, from, to, token, value)
}

// TransferTokenToHandle is a paid mutator transaction binding the contract method 0xa5b7606b.
//
// Solidity: function transferTokenToHandle(string from, string to, address token, uint256 value) returns(bool)
func (_Wallet *WalletTransactorSession) TransferTokenToHandle(from string, to string, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferTokenToHandle(&_Wallet.TransactOpts, from, to, token, value)
}

// WalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wallet contract.
type WalletOwnershipTransferredIterator struct {
	Event *WalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletOwnershipTransferred)
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
		it.Event = new(WalletOwnershipTransferred)
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
func (it *WalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletOwnershipTransferred represents a OwnershipTransferred event raised by the Wallet contract.
type WalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wallet *WalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletOwnershipTransferredIterator{contract: _Wallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wallet *WalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletOwnershipTransferred)
				if err := _Wallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
