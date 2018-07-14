// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PermissionsABI is the input ABI used to generate the binding from.
const PermissionsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_canWrite\",\"type\":\"bool\"},{\"name\":\"_canLead\",\"type\":\"bool\"}],\"name\":\"AddNode\",\"outputs\":[{\"name\":\"_nodeId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_nodeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_canWrite\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_canLead\",\"type\":\"bool\"}],\"name\":\"NewNodeAdded\",\"type\":\"event\"}]"

// Permissions is an auto generated Go binding around an Ethereum contract.
type Permissions struct {
	PermissionsCaller     // Read-only binding to the contract
	PermissionsTransactor // Write-only binding to the contract
}

// PermissionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionsSession struct {
	Contract     *Permissions      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsCallerSession struct {
	Contract *PermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsTransactorSession struct {
	Contract     *PermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PermissionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionsRaw struct {
	Contract *Permissions // Generic contract binding to access the raw methods on
}

// PermissionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionsCallerRaw struct {
	Contract *PermissionsCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionsTransactorRaw struct {
	Contract *PermissionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissions creates a new instance of Permissions, bound to a specific deployed contract.
func NewPermissions(address common.Address, backend bind.ContractBackend) (*Permissions, error) {
	contract, err := bindPermissions(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissions{PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}}, nil
}

// NewPermissionsCaller creates a new read-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsCaller(address common.Address, caller bind.ContractCaller) (*PermissionsCaller, error) {
	contract, err := bindPermissions(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsCaller{contract: contract}, nil
}

// NewPermissionsTransactor creates a new write-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionsTransactor, error) {
	contract, err := bindPermissions(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PermissionsTransactor{contract: contract}, nil
}

// bindPermissions binds a generic wrapper to an already deployed contract.
func bindPermissions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.PermissionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transact(opts, method, params...)
}

// AddNode is a paid mutator transaction binding the contract method 0x378d5115.
//
// Solidity: function AddNode(_enodeId string, _canWrite bool, _canLead bool) returns(_nodeId uint256)
func (_Permissions *PermissionsTransactor) AddNode(opts *bind.TransactOpts, _enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "AddNode", _enodeId, _canWrite, _canLead)
}

// AddNode is a paid mutator transaction binding the contract method 0x378d5115.
//
// Solidity: function AddNode(_enodeId string, _canWrite bool, _canLead bool) returns(_nodeId uint256)
func (_Permissions *PermissionsSession) AddNode(_enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _enodeId, _canWrite, _canLead)
}

// AddNode is a paid mutator transaction binding the contract method 0x378d5115.
//
// Solidity: function AddNode(_enodeId string, _canWrite bool, _canLead bool) returns(_nodeId uint256)
func (_Permissions *PermissionsTransactorSession) AddNode(_enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _enodeId, _canWrite, _canLead)
}
