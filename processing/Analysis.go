// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Analysis

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

// AnalysisMetaData contains all meta data concerning the Analysis contract.
var AnalysisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"Benchmark\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101a88061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063239b51bf1461002d575b5f80fd5b610047600480360381019061004291906100ad565b610049565b005b5f805f90505b8281101561007157630129f6509150808061006990610105565b91505061004f565b505050565b5f80fd5b5f819050919050565b61008c8161007a565b8114610096575f80fd5b50565b5f813590506100a781610083565b92915050565b5f602082840312156100c2576100c1610076565b5b5f6100cf84828501610099565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61010f8261007a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610141576101406100d8565b5b60018201905091905056fea26469706673582212209d6355e2c5daca9ee2dbd926c0cfcd1922e1ffcb815cd9b61ba6939f952c590964736f6c637828302e382e32322d646576656c6f702e323032332e392e31362b636f6d6d69742e30323062353936380059",
}

// AnalysisABI is the input ABI used to generate the binding from.
// Deprecated: Use AnalysisMetaData.ABI instead.
var AnalysisABI = AnalysisMetaData.ABI

// AnalysisBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AnalysisMetaData.Bin instead.
var AnalysisBin = AnalysisMetaData.Bin

// DeployAnalysis deploys a new Ethereum contract, binding an instance of Analysis to it.
func DeployAnalysis(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Analysis, error) {
	parsed, err := AnalysisMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AnalysisBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Analysis{AnalysisCaller: AnalysisCaller{contract: contract}, AnalysisTransactor: AnalysisTransactor{contract: contract}, AnalysisFilterer: AnalysisFilterer{contract: contract}}, nil
}

// Analysis is an auto generated Go binding around an Ethereum contract.
type Analysis struct {
	AnalysisCaller     // Read-only binding to the contract
	AnalysisTransactor // Write-only binding to the contract
	AnalysisFilterer   // Log filterer for contract events
}

// AnalysisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnalysisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnalysisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnalysisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnalysisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnalysisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnalysisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnalysisSession struct {
	Contract     *Analysis         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnalysisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnalysisCallerSession struct {
	Contract *AnalysisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AnalysisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnalysisTransactorSession struct {
	Contract     *AnalysisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AnalysisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnalysisRaw struct {
	Contract *Analysis // Generic contract binding to access the raw methods on
}

// AnalysisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnalysisCallerRaw struct {
	Contract *AnalysisCaller // Generic read-only contract binding to access the raw methods on
}

// AnalysisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnalysisTransactorRaw struct {
	Contract *AnalysisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnalysis creates a new instance of Analysis, bound to a specific deployed contract.
func NewAnalysis(address common.Address, backend bind.ContractBackend) (*Analysis, error) {
	contract, err := bindAnalysis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Analysis{AnalysisCaller: AnalysisCaller{contract: contract}, AnalysisTransactor: AnalysisTransactor{contract: contract}, AnalysisFilterer: AnalysisFilterer{contract: contract}}, nil
}

// NewAnalysisCaller creates a new read-only instance of Analysis, bound to a specific deployed contract.
func NewAnalysisCaller(address common.Address, caller bind.ContractCaller) (*AnalysisCaller, error) {
	contract, err := bindAnalysis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnalysisCaller{contract: contract}, nil
}

// NewAnalysisTransactor creates a new write-only instance of Analysis, bound to a specific deployed contract.
func NewAnalysisTransactor(address common.Address, transactor bind.ContractTransactor) (*AnalysisTransactor, error) {
	contract, err := bindAnalysis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnalysisTransactor{contract: contract}, nil
}

// NewAnalysisFilterer creates a new log filterer instance of Analysis, bound to a specific deployed contract.
func NewAnalysisFilterer(address common.Address, filterer bind.ContractFilterer) (*AnalysisFilterer, error) {
	contract, err := bindAnalysis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnalysisFilterer{contract: contract}, nil
}

// bindAnalysis binds a generic wrapper to an already deployed contract.
func bindAnalysis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnalysisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Analysis *AnalysisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Analysis.Contract.AnalysisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Analysis *AnalysisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Analysis.Contract.AnalysisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Analysis *AnalysisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Analysis.Contract.AnalysisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Analysis *AnalysisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Analysis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Analysis *AnalysisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Analysis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Analysis *AnalysisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Analysis.Contract.contract.Transact(opts, method, params...)
}

// Benchmark is a paid mutator transaction binding the contract method 0x239b51bf.
//
// Solidity: function Benchmark(uint256 limit) returns()
func (_Analysis *AnalysisTransactor) Benchmark(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Analysis.contract.Transact(opts, "Benchmark", limit)
}

// Benchmark is a paid mutator transaction binding the contract method 0x239b51bf.
//
// Solidity: function Benchmark(uint256 limit) returns()
func (_Analysis *AnalysisSession) Benchmark(limit *big.Int) (*types.Transaction, error) {
	return _Analysis.Contract.Benchmark(&_Analysis.TransactOpts, limit)
}

// Benchmark is a paid mutator transaction binding the contract method 0x239b51bf.
//
// Solidity: function Benchmark(uint256 limit) returns()
func (_Analysis *AnalysisTransactorSession) Benchmark(limit *big.Int) (*types.Transaction, error) {
	return _Analysis.Contract.Benchmark(&_Analysis.TransactOpts, limit)
}
