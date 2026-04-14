// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.ConvertType
)

// VoterRegistryMetaData contains all meta data concerning the VoterRegistry contract.
var VoterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registrars\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"}],\"name\":\"VoterRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRAR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfRegistrationOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"name\":\"setSelfRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VoterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterRegistryMetaData.ABI instead.
var VoterRegistryABI = VoterRegistryMetaData.ABI

// VoterRegistry is an auto generated Go binding around an Ethereum contract.
type VoterRegistry struct {
	VoterRegistryCaller     // Read-only binding to the contract
	VoterRegistryTransactor // Write-only binding to the contract
	VoterRegistryFilterer   // Log filterer for contract events
}

// VoterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterRegistrySession struct {
	Contract     *VoterRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterRegistryCallerSession struct {
	Contract *VoterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VoterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterRegistryTransactorSession struct {
	Contract     *VoterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VoterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterRegistryRaw struct {
	Contract *VoterRegistry // Generic contract binding to access the raw methods on
}

// VoterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterRegistryCallerRaw struct {
	Contract *VoterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// VoterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterRegistryTransactorRaw struct {
	Contract *VoterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterRegistry creates a new instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistry(address common.Address, backend bind.ContractBackend) (*VoterRegistry, error) {
	contract, err := bindVoterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterRegistry{VoterRegistryCaller: VoterRegistryCaller{contract: contract}, VoterRegistryTransactor: VoterRegistryTransactor{contract: contract}, VoterRegistryFilterer: VoterRegistryFilterer{contract: contract}}, nil
}

// NewVoterRegistryCaller creates a new read-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryCaller(address common.Address, caller bind.ContractCaller) (*VoterRegistryCaller, error) {
	contract, err := bindVoterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryCaller{contract: contract}, nil
}

// NewVoterRegistryTransactor creates a new write-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterRegistryTransactor, error) {
	contract, err := bindVoterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryTransactor{contract: contract}, nil
}

// NewVoterRegistryFilterer creates a new log filterer instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterRegistryFilterer, error) {
	contract, err := bindVoterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryFilterer{contract: contract}, nil
}

// bindVoterRegistry binds a generic wrapper to an already deployed contract.
func bindVoterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.VoterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoterRegistry.Contract.DEFAULTADMINROLE(&_VoterRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoterRegistry.Contract.DEFAULTADMINROLE(&_VoterRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistryCaller) REGISTRARROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "REGISTRAR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistrySession) REGISTRARROLE() ([32]byte, error) {
	return _VoterRegistry.Contract.REGISTRARROLE(&_VoterRegistry.CallOpts)
}

// REGISTRARROLE is a free data retrieval call binding the contract method 0xf68e9553.
//
// Solidity: function REGISTRAR_ROLE() view returns(bytes32)
func (_VoterRegistry *VoterRegistryCallerSession) REGISTRARROLE() ([32]byte, error) {
	return _VoterRegistry.Contract.REGISTRARROLE(&_VoterRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoterRegistry *VoterRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoterRegistry *VoterRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoterRegistry.Contract.GetRoleAdmin(&_VoterRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoterRegistry *VoterRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoterRegistry.Contract.GetRoleAdmin(&_VoterRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoterRegistry *VoterRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoterRegistry.Contract.HasRole(&_VoterRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoterRegistry.Contract.HasRole(&_VoterRegistry.CallOpts, role, account)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_VoterRegistry *VoterRegistrySession) IsRegistered(arg0 common.Address) (bool, error) {
	return _VoterRegistry.Contract.IsRegistered(&_VoterRegistry.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _VoterRegistry.Contract.IsRegistered(&_VoterRegistry.CallOpts, arg0)
}

// SelfRegistrationOpen is a free data retrieval call binding the contract method 0x3bbbfa99.
//
// Solidity: function selfRegistrationOpen() view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) SelfRegistrationOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "selfRegistrationOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SelfRegistrationOpen is a free data retrieval call binding the contract method 0x3bbbfa99.
//
// Solidity: function selfRegistrationOpen() view returns(bool)
func (_VoterRegistry *VoterRegistrySession) SelfRegistrationOpen() (bool, error) {
	return _VoterRegistry.Contract.SelfRegistrationOpen(&_VoterRegistry.CallOpts)
}

// SelfRegistrationOpen is a free data retrieval call binding the contract method 0x3bbbfa99.
//
// Solidity: function selfRegistrationOpen() view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) SelfRegistrationOpen() (bool, error) {
	return _VoterRegistry.Contract.SelfRegistrationOpen(&_VoterRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoterRegistry *VoterRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoterRegistry.Contract.SupportsInterface(&_VoterRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoterRegistry.Contract.SupportsInterface(&_VoterRegistry.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.GrantRole(&_VoterRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.GrantRole(&_VoterRegistry.TransactOpts, role, account)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _voter) returns()
func (_VoterRegistry *VoterRegistryTransactor) Register(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "register", _voter)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _voter) returns()
func (_VoterRegistry *VoterRegistrySession) Register(_voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.Register(&_VoterRegistry.TransactOpts, _voter)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address _voter) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) Register(_voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.Register(&_VoterRegistry.TransactOpts, _voter)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VoterRegistry *VoterRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VoterRegistry *VoterRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RenounceRole(&_VoterRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RenounceRole(&_VoterRegistry.TransactOpts, role, callerConfirmation)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _voter) returns()
func (_VoterRegistry *VoterRegistryTransactor) Revoke(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "revoke", _voter)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _voter) returns()
func (_VoterRegistry *VoterRegistrySession) Revoke(_voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.Revoke(&_VoterRegistry.TransactOpts, _voter)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _voter) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) Revoke(_voter common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.Revoke(&_VoterRegistry.TransactOpts, _voter)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RevokeRole(&_VoterRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RevokeRole(&_VoterRegistry.TransactOpts, role, account)
}

// SelfRegister is a paid mutator transaction binding the contract method 0xd86a28bb.
//
// Solidity: function selfRegister() returns()
func (_VoterRegistry *VoterRegistryTransactor) SelfRegister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "selfRegister")
}

// SelfRegister is a paid mutator transaction binding the contract method 0xd86a28bb.
//
// Solidity: function selfRegister() returns()
func (_VoterRegistry *VoterRegistrySession) SelfRegister() (*types.Transaction, error) {
	return _VoterRegistry.Contract.SelfRegister(&_VoterRegistry.TransactOpts)
}

// SelfRegister is a paid mutator transaction binding the contract method 0xd86a28bb.
//
// Solidity: function selfRegister() returns()
func (_VoterRegistry *VoterRegistryTransactorSession) SelfRegister() (*types.Transaction, error) {
	return _VoterRegistry.Contract.SelfRegister(&_VoterRegistry.TransactOpts)
}

// SetSelfRegistration is a paid mutator transaction binding the contract method 0x85bb6c46.
//
// Solidity: function setSelfRegistration(bool _open) returns()
func (_VoterRegistry *VoterRegistryTransactor) SetSelfRegistration(opts *bind.TransactOpts, _open bool) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "setSelfRegistration", _open)
}

// SetSelfRegistration is a paid mutator transaction binding the contract method 0x85bb6c46.
//
// Solidity: function setSelfRegistration(bool _open) returns()
func (_VoterRegistry *VoterRegistrySession) SetSelfRegistration(_open bool) (*types.Transaction, error) {
	return _VoterRegistry.Contract.SetSelfRegistration(&_VoterRegistry.TransactOpts, _open)
}

// SetSelfRegistration is a paid mutator transaction binding the contract method 0x85bb6c46.
//
// Solidity: function setSelfRegistration(bool _open) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) SetSelfRegistration(_open bool) (*types.Transaction, error) {
	return _VoterRegistry.Contract.SetSelfRegistration(&_VoterRegistry.TransactOpts, _open)
}

// VoterRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VoterRegistry contract.
type VoterRegistryRoleAdminChangedIterator struct {
	Event *VoterRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoterRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryRoleAdminChanged)
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
		it.Event = new(VoterRegistryRoleAdminChanged)
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
func (it *VoterRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the VoterRegistry contract.
type VoterRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoterRegistry *VoterRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VoterRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryRoleAdminChangedIterator{contract: _VoterRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoterRegistry *VoterRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VoterRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryRoleAdminChanged)
				if err := _VoterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoterRegistry *VoterRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*VoterRegistryRoleAdminChanged, error) {
	event := new(VoterRegistryRoleAdminChanged)
	if err := _VoterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VoterRegistry contract.
type VoterRegistryRoleGrantedIterator struct {
	Event *VoterRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *VoterRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryRoleGranted)
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
		it.Event = new(VoterRegistryRoleGranted)
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
func (it *VoterRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryRoleGranted represents a RoleGranted event raised by the VoterRegistry contract.
type VoterRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoterRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryRoleGrantedIterator{contract: _VoterRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VoterRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryRoleGranted)
				if err := _VoterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) ParseRoleGranted(log types.Log) (*VoterRegistryRoleGranted, error) {
	event := new(VoterRegistryRoleGranted)
	if err := _VoterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VoterRegistry contract.
type VoterRegistryRoleRevokedIterator struct {
	Event *VoterRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VoterRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryRoleRevoked)
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
		it.Event = new(VoterRegistryRoleRevoked)
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
func (it *VoterRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryRoleRevoked represents a RoleRevoked event raised by the VoterRegistry contract.
type VoterRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoterRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryRoleRevokedIterator{contract: _VoterRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VoterRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryRoleRevoked)
				if err := _VoterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoterRegistry *VoterRegistryFilterer) ParseRoleRevoked(log types.Log) (*VoterRegistryRoleRevoked, error) {
	event := new(VoterRegistryRoleRevoked)
	if err := _VoterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the VoterRegistry contract.
type VoterRegistryVoterRegisteredIterator struct {
	Event *VoterRegistryVoterRegistered // Event containing the contract specifics and raw log

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
func (it *VoterRegistryVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryVoterRegistered)
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
		it.Event = new(VoterRegistryVoterRegistered)
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
func (it *VoterRegistryVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryVoterRegistered represents a VoterRegistered event raised by the VoterRegistry contract.
type VoterRegistryVoterRegistered struct {
	Voter     common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0x76f495f538ae23446b46f9ea62c1a4d8a985e56002509d25fd87a6dbbcff75ab.
//
// Solidity: event VoterRegistered(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) FilterVoterRegistered(opts *bind.FilterOpts, voter []common.Address, registrar []common.Address) (*VoterRegistryVoterRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "VoterRegistered", voterRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryVoterRegisteredIterator{contract: _VoterRegistry.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0x76f495f538ae23446b46f9ea62c1a4d8a985e56002509d25fd87a6dbbcff75ab.
//
// Solidity: event VoterRegistered(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *VoterRegistryVoterRegistered, voter []common.Address, registrar []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "VoterRegistered", voterRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryVoterRegistered)
				if err := _VoterRegistry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
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

// ParseVoterRegistered is a log parse operation binding the contract event 0x76f495f538ae23446b46f9ea62c1a4d8a985e56002509d25fd87a6dbbcff75ab.
//
// Solidity: event VoterRegistered(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) ParseVoterRegistered(log types.Log) (*VoterRegistryVoterRegistered, error) {
	event := new(VoterRegistryVoterRegistered)
	if err := _VoterRegistry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryVoterRevokedIterator is returned from FilterVoterRevoked and is used to iterate over the raw logs and unpacked data for VoterRevoked events raised by the VoterRegistry contract.
type VoterRegistryVoterRevokedIterator struct {
	Event *VoterRegistryVoterRevoked // Event containing the contract specifics and raw log

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
func (it *VoterRegistryVoterRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryVoterRevoked)
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
		it.Event = new(VoterRegistryVoterRevoked)
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
func (it *VoterRegistryVoterRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryVoterRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryVoterRevoked represents a VoterRevoked event raised by the VoterRegistry contract.
type VoterRegistryVoterRevoked struct {
	Voter     common.Address
	Registrar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterRevoked is a free log retrieval operation binding the contract event 0x3f504b50e6d35791a3b87cb4c64218c995a97a3fdcfca176f770ebcc344effd2.
//
// Solidity: event VoterRevoked(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) FilterVoterRevoked(opts *bind.FilterOpts, voter []common.Address, registrar []common.Address) (*VoterRegistryVoterRevokedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "VoterRevoked", voterRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryVoterRevokedIterator{contract: _VoterRegistry.contract, event: "VoterRevoked", logs: logs, sub: sub}, nil
}

// WatchVoterRevoked is a free log subscription operation binding the contract event 0x3f504b50e6d35791a3b87cb4c64218c995a97a3fdcfca176f770ebcc344effd2.
//
// Solidity: event VoterRevoked(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) WatchVoterRevoked(opts *bind.WatchOpts, sink chan<- *VoterRegistryVoterRevoked, voter []common.Address, registrar []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "VoterRevoked", voterRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryVoterRevoked)
				if err := _VoterRegistry.contract.UnpackLog(event, "VoterRevoked", log); err != nil {
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

// ParseVoterRevoked is a log parse operation binding the contract event 0x3f504b50e6d35791a3b87cb4c64218c995a97a3fdcfca176f770ebcc344effd2.
//
// Solidity: event VoterRevoked(address indexed voter, address indexed registrar)
func (_VoterRegistry *VoterRegistryFilterer) ParseVoterRevoked(log types.Log) (*VoterRegistryVoterRevoked, error) {
	event := new(VoterRegistryVoterRevoked)
	if err := _VoterRegistry.contract.UnpackLog(event, "VoterRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
