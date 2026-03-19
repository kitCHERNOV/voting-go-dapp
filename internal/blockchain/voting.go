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

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerCandidateId\",\"type\":\"uint256\"}],\"name\":\"ProposalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_durationSec\",\"type\":\"uint256\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"finalizeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getResults\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"votes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xb1ff97c1.
//
// Solidity: function candidateCount(uint256 ) view returns(uint256)
func (_Voting *VotingCaller) CandidateCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidateCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateCount is a free data retrieval call binding the contract method 0xb1ff97c1.
//
// Solidity: function candidateCount(uint256 ) view returns(uint256)
func (_Voting *VotingSession) CandidateCount(arg0 *big.Int) (*big.Int, error) {
	return _Voting.Contract.CandidateCount(&_Voting.CallOpts, arg0)
}

// CandidateCount is a free data retrieval call binding the contract method 0xb1ff97c1.
//
// Solidity: function candidateCount(uint256 ) view returns(uint256)
func (_Voting *VotingCallerSession) CandidateCount(arg0 *big.Int) (*big.Int, error) {
	return _Voting.Contract.CandidateCount(&_Voting.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x7de14242.
//
// Solidity: function candidates(uint256 , uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidates", arg0, arg1)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x7de14242.
//
// Solidity: function candidates(uint256 , uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingSession) Candidates(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0, arg1)
}

// Candidates is a free data retrieval call binding the contract method 0x7de14242.
//
// Solidity: function candidates(uint256 , uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCallerSession) Candidates(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0, arg1)
}

// GetResults is a free data retrieval call binding the contract method 0x81a60c0d.
//
// Solidity: function getResults(uint256 _proposalId) view returns(uint256[] ids, uint256[] votes)
func (_Voting *VotingCaller) GetResults(opts *bind.CallOpts, _proposalId *big.Int) (struct {
	Ids   []*big.Int
	Votes []*big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getResults", _proposalId)

	outstruct := new(struct {
		Ids   []*big.Int
		Votes []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Votes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetResults is a free data retrieval call binding the contract method 0x81a60c0d.
//
// Solidity: function getResults(uint256 _proposalId) view returns(uint256[] ids, uint256[] votes)
func (_Voting *VotingSession) GetResults(_proposalId *big.Int) (struct {
	Ids   []*big.Int
	Votes []*big.Int
}, error) {
	return _Voting.Contract.GetResults(&_Voting.CallOpts, _proposalId)
}

// GetResults is a free data retrieval call binding the contract method 0x81a60c0d.
//
// Solidity: function getResults(uint256 _proposalId) view returns(uint256[] ids, uint256[] votes)
func (_Voting *VotingCallerSession) GetResults(_proposalId *big.Int) (struct {
	Ids   []*big.Int
	Votes []*big.Int
}, error) {
	return _Voting.Contract.GetResults(&_Voting.CallOpts, _proposalId)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Voting *VotingCaller) HasVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "hasVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Voting *VotingSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0, arg1)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Voting *VotingCallerSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0, arg1)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Voting *VotingCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Voting *VotingSession) ProposalCount() (*big.Int, error) {
	return _Voting.Contract.ProposalCount(&_Voting.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Voting *VotingCallerSession) ProposalCount() (*big.Int, error) {
	return _Voting.Contract.ProposalCount(&_Voting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string title, string description, address creator, uint256 startTime, uint256 endTime, bool finalized, uint256 totalVotes)
func (_Voting *VotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Creator     common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Finalized   bool
	TotalVotes  *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Title       string
		Description string
		Creator     common.Address
		StartTime   *big.Int
		EndTime     *big.Int
		Finalized   bool
		TotalVotes  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Finalized = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.TotalVotes = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string title, string description, address creator, uint256 startTime, uint256 endTime, bool finalized, uint256 totalVotes)
func (_Voting *VotingSession) Proposals(arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Creator     common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Finalized   bool
	TotalVotes  *big.Int
}, error) {
	return _Voting.Contract.Proposals(&_Voting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string title, string description, address creator, uint256 startTime, uint256 endTime, bool finalized, uint256 totalVotes)
func (_Voting *VotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Id          *big.Int
	Title       string
	Description string
	Creator     common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Finalized   bool
	TotalVotes  *big.Int
}, error) {
	return _Voting.Contract.Proposals(&_Voting.CallOpts, arg0)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1750a3d0.
//
// Solidity: function addCandidate(uint256 _proposalId, string _name) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, _proposalId *big.Int, _name string) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", _proposalId, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1750a3d0.
//
// Solidity: function addCandidate(uint256 _proposalId, string _name) returns()
func (_Voting *VotingSession) AddCandidate(_proposalId *big.Int, _name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _proposalId, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x1750a3d0.
//
// Solidity: function addCandidate(uint256 _proposalId, string _name) returns()
func (_Voting *VotingTransactorSession) AddCandidate(_proposalId *big.Int, _name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _proposalId, _name)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x6b5e1766.
//
// Solidity: function createProposal(string _title, string _description, uint256 _startDelay, uint256 _durationSec) returns(uint256 proposalId)
func (_Voting *VotingTransactor) CreateProposal(opts *bind.TransactOpts, _title string, _description string, _startDelay *big.Int, _durationSec *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "createProposal", _title, _description, _startDelay, _durationSec)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x6b5e1766.
//
// Solidity: function createProposal(string _title, string _description, uint256 _startDelay, uint256 _durationSec) returns(uint256 proposalId)
func (_Voting *VotingSession) CreateProposal(_title string, _description string, _startDelay *big.Int, _durationSec *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.CreateProposal(&_Voting.TransactOpts, _title, _description, _startDelay, _durationSec)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x6b5e1766.
//
// Solidity: function createProposal(string _title, string _description, uint256 _startDelay, uint256 _durationSec) returns(uint256 proposalId)
func (_Voting *VotingTransactorSession) CreateProposal(_title string, _description string, _startDelay *big.Int, _durationSec *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.CreateProposal(&_Voting.TransactOpts, _title, _description, _startDelay, _durationSec)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 _proposalId) returns()
func (_Voting *VotingTransactor) FinalizeProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "finalizeProposal", _proposalId)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 _proposalId) returns()
func (_Voting *VotingSession) FinalizeProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.FinalizeProposal(&_Voting.TransactOpts, _proposalId)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0x5652077c.
//
// Solidity: function finalizeProposal(uint256 _proposalId) returns()
func (_Voting *VotingTransactorSession) FinalizeProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.FinalizeProposal(&_Voting.TransactOpts, _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _proposalId, uint256 _candidateId) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, _proposalId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", _proposalId, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _proposalId, uint256 _candidateId) returns()
func (_Voting *VotingSession) Vote(_proposalId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _proposalId, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 _proposalId, uint256 _candidateId) returns()
func (_Voting *VotingTransactorSession) Vote(_proposalId *big.Int, _candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, _proposalId, _candidateId)
}

// VotingCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Voting contract.
type VotingCandidateAddedIterator struct {
	Event *VotingCandidateAdded // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAdded)
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
		it.Event = new(VotingCandidateAdded)
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
func (it *VotingCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAdded represents a CandidateAdded event raised by the Voting contract.
type VotingCandidateAdded struct {
	ProposalId  *big.Int
	CandidateId *big.Int
	Name        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0xed8911b3df733b7d5f75724158e54478ea12e30f49c9d31b5261879f5b76586f.
//
// Solidity: event CandidateAdded(uint256 indexed proposalId, uint256 indexed candidateId, string name)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts, proposalId []*big.Int, candidateId []*big.Int) (*VotingCandidateAddedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded", proposalIdRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0xed8911b3df733b7d5f75724158e54478ea12e30f49c9d31b5261879f5b76586f.
//
// Solidity: event CandidateAdded(uint256 indexed proposalId, uint256 indexed candidateId, string name)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded, proposalId []*big.Int, candidateId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded", proposalIdRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAdded)
				if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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

// ParseCandidateAdded is a log parse operation binding the contract event 0xed8911b3df733b7d5f75724158e54478ea12e30f49c9d31b5261879f5b76586f.
//
// Solidity: event CandidateAdded(uint256 indexed proposalId, uint256 indexed candidateId, string name)
func (_Voting *VotingFilterer) ParseCandidateAdded(log types.Log) (*VotingCandidateAdded, error) {
	event := new(VotingCandidateAdded)
	if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Voting contract.
type VotingProposalCreatedIterator struct {
	Event *VotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *VotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingProposalCreated)
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
		it.Event = new(VotingProposalCreated)
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
func (it *VotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingProposalCreated represents a ProposalCreated event raised by the Voting contract.
type VotingProposalCreated struct {
	ProposalId *big.Int
	Title      string
	Creator    common.Address
	StartTime  *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x331b791e5a0526c6d6f17a46c1d4139542c8579c245d9f8082d0b1462942f4fa.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, string title, address indexed creator, uint256 startTime, uint256 endTime)
func (_Voting *VotingFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalId []*big.Int, creator []common.Address) (*VotingProposalCreatedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "ProposalCreated", proposalIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &VotingProposalCreatedIterator{contract: _Voting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x331b791e5a0526c6d6f17a46c1d4139542c8579c245d9f8082d0b1462942f4fa.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, string title, address indexed creator, uint256 startTime, uint256 endTime)
func (_Voting *VotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *VotingProposalCreated, proposalId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "ProposalCreated", proposalIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingProposalCreated)
				if err := _Voting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x331b791e5a0526c6d6f17a46c1d4139542c8579c245d9f8082d0b1462942f4fa.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, string title, address indexed creator, uint256 startTime, uint256 endTime)
func (_Voting *VotingFilterer) ParseProposalCreated(log types.Log) (*VotingProposalCreated, error) {
	event := new(VotingProposalCreated)
	if err := _Voting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingProposalFinalizedIterator is returned from FilterProposalFinalized and is used to iterate over the raw logs and unpacked data for ProposalFinalized events raised by the Voting contract.
type VotingProposalFinalizedIterator struct {
	Event *VotingProposalFinalized // Event containing the contract specifics and raw log

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
func (it *VotingProposalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingProposalFinalized)
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
		it.Event = new(VotingProposalFinalized)
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
func (it *VotingProposalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingProposalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingProposalFinalized represents a ProposalFinalized event raised by the Voting contract.
type VotingProposalFinalized struct {
	ProposalId        *big.Int
	WinnerCandidateId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterProposalFinalized is a free log retrieval operation binding the contract event 0x7701dba108504636e1c3edeae33d8c8894d878b94ace52c051e49cb5aeb0fe05.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, uint256 winnerCandidateId)
func (_Voting *VotingFilterer) FilterProposalFinalized(opts *bind.FilterOpts, proposalId []*big.Int) (*VotingProposalFinalizedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingProposalFinalizedIterator{contract: _Voting.contract, event: "ProposalFinalized", logs: logs, sub: sub}, nil
}

// WatchProposalFinalized is a free log subscription operation binding the contract event 0x7701dba108504636e1c3edeae33d8c8894d878b94ace52c051e49cb5aeb0fe05.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, uint256 winnerCandidateId)
func (_Voting *VotingFilterer) WatchProposalFinalized(opts *bind.WatchOpts, sink chan<- *VotingProposalFinalized, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingProposalFinalized)
				if err := _Voting.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
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

// ParseProposalFinalized is a log parse operation binding the contract event 0x7701dba108504636e1c3edeae33d8c8894d878b94ace52c051e49cb5aeb0fe05.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, uint256 winnerCandidateId)
func (_Voting *VotingFilterer) ParseProposalFinalized(log types.Log) (*VotingProposalFinalized, error) {
	event := new(VotingProposalFinalized)
	if err := _Voting.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Voting contract.
type VotingVoteCastIterator struct {
	Event *VotingVoteCast // Event containing the contract specifics and raw log

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
func (it *VotingVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoteCast)
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
		it.Event = new(VotingVoteCast)
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
func (it *VotingVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoteCast represents a VoteCast event raised by the Voting contract.
type VotingVoteCast struct {
	ProposalId  *big.Int
	Voter       common.Address
	CandidateId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x2acce567deca3aabf56327adbb4524bd5318936eaefa69e3a5208ffda0cfec09.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, uint256 indexed candidateId)
func (_Voting *VotingFilterer) FilterVoteCast(opts *bind.FilterOpts, proposalId []*big.Int, voter []common.Address, candidateId []*big.Int) (*VotingVoteCastIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VoteCast", proposalIdRule, voterRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &VotingVoteCastIterator{contract: _Voting.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x2acce567deca3aabf56327adbb4524bd5318936eaefa69e3a5208ffda0cfec09.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, uint256 indexed candidateId)
func (_Voting *VotingFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *VotingVoteCast, proposalId []*big.Int, voter []common.Address, candidateId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VoteCast", proposalIdRule, voterRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoteCast)
				if err := _Voting.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x2acce567deca3aabf56327adbb4524bd5318936eaefa69e3a5208ffda0cfec09.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address indexed voter, uint256 indexed candidateId)
func (_Voting *VotingFilterer) ParseVoteCast(log types.Log) (*VotingVoteCast, error) {
	event := new(VotingVoteCast)
	if err := _Voting.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
