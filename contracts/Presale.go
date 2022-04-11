// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package presale

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
)

// PresaleMetaData contains all meta data concerning the Presale contract.
var PresaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Closed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OverHardCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverMaxContribution\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnapprovedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderMinContribution\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Contribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"TokenState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HARD_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"approved\",\"type\":\"bool[]\"}],\"name\":\"setApprovedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setMaxContribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newState\",\"type\":\"bool\"}],\"name\":\"setSaleOpen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PresaleABI is the input ABI used to generate the binding from.
// Deprecated: Use PresaleMetaData.ABI instead.
var PresaleABI = PresaleMetaData.ABI

// Presale is an auto generated Go binding around an Ethereum contract.
type Presale struct {
	PresaleCaller     // Read-only binding to the contract
	PresaleTransactor // Write-only binding to the contract
	PresaleFilterer   // Log filterer for contract events
}

// PresaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PresaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PresaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PresaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PresaleSession struct {
	Contract     *Presale          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PresaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PresaleCallerSession struct {
	Contract *PresaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PresaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PresaleTransactorSession struct {
	Contract     *PresaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PresaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PresaleRaw struct {
	Contract *Presale // Generic contract binding to access the raw methods on
}

// PresaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PresaleCallerRaw struct {
	Contract *PresaleCaller // Generic read-only contract binding to access the raw methods on
}

// PresaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PresaleTransactorRaw struct {
	Contract *PresaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPresale creates a new instance of Presale, bound to a specific deployed contract.
func NewPresale(address common.Address, backend bind.ContractBackend) (*Presale, error) {
	contract, err := bindPresale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Presale{PresaleCaller: PresaleCaller{contract: contract}, PresaleTransactor: PresaleTransactor{contract: contract}, PresaleFilterer: PresaleFilterer{contract: contract}}, nil
}

// NewPresaleCaller creates a new read-only instance of Presale, bound to a specific deployed contract.
func NewPresaleCaller(address common.Address, caller bind.ContractCaller) (*PresaleCaller, error) {
	contract, err := bindPresale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleCaller{contract: contract}, nil
}

// NewPresaleTransactor creates a new write-only instance of Presale, bound to a specific deployed contract.
func NewPresaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PresaleTransactor, error) {
	contract, err := bindPresale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleTransactor{contract: contract}, nil
}

// NewPresaleFilterer creates a new log filterer instance of Presale, bound to a specific deployed contract.
func NewPresaleFilterer(address common.Address, filterer bind.ContractFilterer) (*PresaleFilterer, error) {
	contract, err := bindPresale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PresaleFilterer{contract: contract}, nil
}

// bindPresale binds a generic wrapper to an already deployed contract.
func bindPresale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PresaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presale *PresaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presale.Contract.PresaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presale *PresaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presale.Contract.PresaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presale *PresaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presale.Contract.PresaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presale *PresaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presale *PresaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presale *PresaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presale.Contract.contract.Transact(opts, method, params...)
}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() view returns(uint256)
func (_Presale *PresaleCaller) HARDCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "HARD_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() view returns(uint256)
func (_Presale *PresaleSession) HARDCAP() (*big.Int, error) {
	return _Presale.Contract.HARDCAP(&_Presale.CallOpts)
}

// HARDCAP is a free data retrieval call binding the contract method 0x3a03171c.
//
// Solidity: function HARD_CAP() view returns(uint256)
func (_Presale *PresaleCallerSession) HARDCAP() (*big.Int, error) {
	return _Presale.Contract.HARDCAP(&_Presale.CallOpts)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_Presale *PresaleCaller) ApprovedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "approvedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_Presale *PresaleSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _Presale.Contract.ApprovedTokens(&_Presale.CallOpts, arg0)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_Presale *PresaleCallerSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _Presale.Contract.ApprovedTokens(&_Presale.CallOpts, arg0)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Presale *PresaleCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Presale *PresaleSession) Beneficiary() (common.Address, error) {
	return _Presale.Contract.Beneficiary(&_Presale.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Presale *PresaleCallerSession) Beneficiary() (common.Address, error) {
	return _Presale.Contract.Beneficiary(&_Presale.CallOpts)
}

// Contributed is a free data retrieval call binding the contract method 0x112f1907.
//
// Solidity: function contributed() view returns(uint256)
func (_Presale *PresaleCaller) Contributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "contributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributed is a free data retrieval call binding the contract method 0x112f1907.
//
// Solidity: function contributed() view returns(uint256)
func (_Presale *PresaleSession) Contributed() (*big.Int, error) {
	return _Presale.Contract.Contributed(&_Presale.CallOpts)
}

// Contributed is a free data retrieval call binding the contract method 0x112f1907.
//
// Solidity: function contributed() view returns(uint256)
func (_Presale *PresaleCallerSession) Contributed() (*big.Int, error) {
	return _Presale.Contract.Contributed(&_Presale.CallOpts)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Presale *PresaleCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "contributions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Presale *PresaleSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Presale.Contract.Contributions(&_Presale.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Presale *PresaleCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Presale.Contract.Contributions(&_Presale.CallOpts, arg0)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Presale *PresaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Presale *PresaleSession) IsOpen() (bool, error) {
	return _Presale.Contract.IsOpen(&_Presale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Presale *PresaleCallerSession) IsOpen() (bool, error) {
	return _Presale.Contract.IsOpen(&_Presale.CallOpts)
}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() view returns(uint256)
func (_Presale *PresaleCaller) MaxContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "maxContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() view returns(uint256)
func (_Presale *PresaleSession) MaxContribution() (*big.Int, error) {
	return _Presale.Contract.MaxContribution(&_Presale.CallOpts)
}

// MaxContribution is a free data retrieval call binding the contract method 0x8d3d6576.
//
// Solidity: function maxContribution() view returns(uint256)
func (_Presale *PresaleCallerSession) MaxContribution() (*big.Int, error) {
	return _Presale.Contract.MaxContribution(&_Presale.CallOpts)
}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() view returns(uint256)
func (_Presale *PresaleCaller) MinContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "minContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() view returns(uint256)
func (_Presale *PresaleSession) MinContribution() (*big.Int, error) {
	return _Presale.Contract.MinContribution(&_Presale.CallOpts)
}

// MinContribution is a free data retrieval call binding the contract method 0xaaffadf3.
//
// Solidity: function minContribution() view returns(uint256)
func (_Presale *PresaleCallerSession) MinContribution() (*big.Int, error) {
	return _Presale.Contract.MinContribution(&_Presale.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Presale *PresaleCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Presale *PresaleSession) Oracle() (common.Address, error) {
	return _Presale.Contract.Oracle(&_Presale.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Presale *PresaleCallerSession) Oracle() (common.Address, error) {
	return _Presale.Contract.Oracle(&_Presale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presale *PresaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Presale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presale *PresaleSession) Owner() (common.Address, error) {
	return _Presale.Contract.Owner(&_Presale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presale *PresaleCallerSession) Owner() (common.Address, error) {
	return _Presale.Contract.Owner(&_Presale.CallOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0x8418cd99.
//
// Solidity: function contribute(address token, uint256 amount) returns()
func (_Presale *PresaleTransactor) Contribute(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "contribute", token, amount)
}

// Contribute is a paid mutator transaction binding the contract method 0x8418cd99.
//
// Solidity: function contribute(address token, uint256 amount) returns()
func (_Presale *PresaleSession) Contribute(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Presale.Contract.Contribute(&_Presale.TransactOpts, token, amount)
}

// Contribute is a paid mutator transaction binding the contract method 0x8418cd99.
//
// Solidity: function contribute(address token, uint256 amount) returns()
func (_Presale *PresaleTransactorSession) Contribute(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Presale.Contract.Contribute(&_Presale.TransactOpts, token, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presale *PresaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presale *PresaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Presale.Contract.RenounceOwnership(&_Presale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presale *PresaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Presale.Contract.RenounceOwnership(&_Presale.TransactOpts)
}

// SetApprovedTokens is a paid mutator transaction binding the contract method 0x92463bed.
//
// Solidity: function setApprovedTokens(address[] tokens, bool[] approved) returns()
func (_Presale *PresaleTransactor) SetApprovedTokens(opts *bind.TransactOpts, tokens []common.Address, approved []bool) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "setApprovedTokens", tokens, approved)
}

// SetApprovedTokens is a paid mutator transaction binding the contract method 0x92463bed.
//
// Solidity: function setApprovedTokens(address[] tokens, bool[] approved) returns()
func (_Presale *PresaleSession) SetApprovedTokens(tokens []common.Address, approved []bool) (*types.Transaction, error) {
	return _Presale.Contract.SetApprovedTokens(&_Presale.TransactOpts, tokens, approved)
}

// SetApprovedTokens is a paid mutator transaction binding the contract method 0x92463bed.
//
// Solidity: function setApprovedTokens(address[] tokens, bool[] approved) returns()
func (_Presale *PresaleTransactorSession) SetApprovedTokens(tokens []common.Address, approved []bool) (*types.Transaction, error) {
	return _Presale.Contract.SetApprovedTokens(&_Presale.TransactOpts, tokens, approved)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address account) returns()
func (_Presale *PresaleTransactor) SetBeneficiary(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "setBeneficiary", account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address account) returns()
func (_Presale *PresaleSession) SetBeneficiary(account common.Address) (*types.Transaction, error) {
	return _Presale.Contract.SetBeneficiary(&_Presale.TransactOpts, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address account) returns()
func (_Presale *PresaleTransactorSession) SetBeneficiary(account common.Address) (*types.Transaction, error) {
	return _Presale.Contract.SetBeneficiary(&_Presale.TransactOpts, account)
}

// SetMaxContribution is a paid mutator transaction binding the contract method 0x03ed9d21.
//
// Solidity: function setMaxContribution(uint256 max) returns()
func (_Presale *PresaleTransactor) SetMaxContribution(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "setMaxContribution", max)
}

// SetMaxContribution is a paid mutator transaction binding the contract method 0x03ed9d21.
//
// Solidity: function setMaxContribution(uint256 max) returns()
func (_Presale *PresaleSession) SetMaxContribution(max *big.Int) (*types.Transaction, error) {
	return _Presale.Contract.SetMaxContribution(&_Presale.TransactOpts, max)
}

// SetMaxContribution is a paid mutator transaction binding the contract method 0x03ed9d21.
//
// Solidity: function setMaxContribution(uint256 max) returns()
func (_Presale *PresaleTransactorSession) SetMaxContribution(max *big.Int) (*types.Transaction, error) {
	return _Presale.Contract.SetMaxContribution(&_Presale.TransactOpts, max)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address account) returns()
func (_Presale *PresaleTransactor) SetOracle(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "setOracle", account)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address account) returns()
func (_Presale *PresaleSession) SetOracle(account common.Address) (*types.Transaction, error) {
	return _Presale.Contract.SetOracle(&_Presale.TransactOpts, account)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address account) returns()
func (_Presale *PresaleTransactorSession) SetOracle(account common.Address) (*types.Transaction, error) {
	return _Presale.Contract.SetOracle(&_Presale.TransactOpts, account)
}

// SetSaleOpen is a paid mutator transaction binding the contract method 0xaf979f25.
//
// Solidity: function setSaleOpen(bool newState) returns()
func (_Presale *PresaleTransactor) SetSaleOpen(opts *bind.TransactOpts, newState bool) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "setSaleOpen", newState)
}

// SetSaleOpen is a paid mutator transaction binding the contract method 0xaf979f25.
//
// Solidity: function setSaleOpen(bool newState) returns()
func (_Presale *PresaleSession) SetSaleOpen(newState bool) (*types.Transaction, error) {
	return _Presale.Contract.SetSaleOpen(&_Presale.TransactOpts, newState)
}

// SetSaleOpen is a paid mutator transaction binding the contract method 0xaf979f25.
//
// Solidity: function setSaleOpen(bool newState) returns()
func (_Presale *PresaleTransactorSession) SetSaleOpen(newState bool) (*types.Transaction, error) {
	return _Presale.Contract.SetSaleOpen(&_Presale.TransactOpts, newState)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presale *PresaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Presale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presale *PresaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presale.Contract.TransferOwnership(&_Presale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presale *PresaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presale.Contract.TransferOwnership(&_Presale.TransactOpts, newOwner)
}

// PresaleContributionIterator is returned from FilterContribution and is used to iterate over the raw logs and unpacked data for Contribution events raised by the Presale contract.
type PresaleContributionIterator struct {
	Event *PresaleContribution // Event containing the contract specifics and raw log

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
func (it *PresaleContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleContribution)
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
		it.Event = new(PresaleContribution)
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
func (it *PresaleContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleContribution represents a Contribution event raised by the Presale contract.
type PresaleContribution struct {
	Buyer  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContribution is a free log retrieval operation binding the contract event 0x71be8766c1bd9bb33623f518aecc1d98612c75b596f4a0c38408f72239ef4e60.
//
// Solidity: event Contribution(address indexed buyer, address indexed token, uint256 amount)
func (_Presale *PresaleFilterer) FilterContribution(opts *bind.FilterOpts, buyer []common.Address, token []common.Address) (*PresaleContributionIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Presale.contract.FilterLogs(opts, "Contribution", buyerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PresaleContributionIterator{contract: _Presale.contract, event: "Contribution", logs: logs, sub: sub}, nil
}

// WatchContribution is a free log subscription operation binding the contract event 0x71be8766c1bd9bb33623f518aecc1d98612c75b596f4a0c38408f72239ef4e60.
//
// Solidity: event Contribution(address indexed buyer, address indexed token, uint256 amount)
func (_Presale *PresaleFilterer) WatchContribution(opts *bind.WatchOpts, sink chan<- *PresaleContribution, buyer []common.Address, token []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Presale.contract.WatchLogs(opts, "Contribution", buyerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleContribution)
				if err := _Presale.contract.UnpackLog(event, "Contribution", log); err != nil {
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

// ParseContribution is a log parse operation binding the contract event 0x71be8766c1bd9bb33623f518aecc1d98612c75b596f4a0c38408f72239ef4e60.
//
// Solidity: event Contribution(address indexed buyer, address indexed token, uint256 amount)
func (_Presale *PresaleFilterer) ParseContribution(log types.Log) (*PresaleContribution, error) {
	event := new(PresaleContribution)
	if err := _Presale.contract.UnpackLog(event, "Contribution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PresaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Presale contract.
type PresaleOwnershipTransferredIterator struct {
	Event *PresaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PresaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleOwnershipTransferred)
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
		it.Event = new(PresaleOwnershipTransferred)
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
func (it *PresaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleOwnershipTransferred represents a OwnershipTransferred event raised by the Presale contract.
type PresaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presale *PresaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PresaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleOwnershipTransferredIterator{contract: _Presale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presale *PresaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PresaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleOwnershipTransferred)
				if err := _Presale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presale *PresaleFilterer) ParseOwnershipTransferred(log types.Log) (*PresaleOwnershipTransferred, error) {
	event := new(PresaleOwnershipTransferred)
	if err := _Presale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PresaleTokenStateIterator is returned from FilterTokenState and is used to iterate over the raw logs and unpacked data for TokenState events raised by the Presale contract.
type PresaleTokenStateIterator struct {
	Event *PresaleTokenState // Event containing the contract specifics and raw log

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
func (it *PresaleTokenStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleTokenState)
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
		it.Event = new(PresaleTokenState)
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
func (it *PresaleTokenStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleTokenStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleTokenState represents a TokenState event raised by the Presale contract.
type PresaleTokenState struct {
	Token common.Address
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenState is a free log retrieval operation binding the contract event 0xb68249cfdf717e902bb48a62af12f230e222538141616fcf91517900ae354d0b.
//
// Solidity: event TokenState(address indexed token, bool state)
func (_Presale *PresaleFilterer) FilterTokenState(opts *bind.FilterOpts, token []common.Address) (*PresaleTokenStateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Presale.contract.FilterLogs(opts, "TokenState", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PresaleTokenStateIterator{contract: _Presale.contract, event: "TokenState", logs: logs, sub: sub}, nil
}

// WatchTokenState is a free log subscription operation binding the contract event 0xb68249cfdf717e902bb48a62af12f230e222538141616fcf91517900ae354d0b.
//
// Solidity: event TokenState(address indexed token, bool state)
func (_Presale *PresaleFilterer) WatchTokenState(opts *bind.WatchOpts, sink chan<- *PresaleTokenState, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Presale.contract.WatchLogs(opts, "TokenState", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleTokenState)
				if err := _Presale.contract.UnpackLog(event, "TokenState", log); err != nil {
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

// ParseTokenState is a log parse operation binding the contract event 0xb68249cfdf717e902bb48a62af12f230e222538141616fcf91517900ae354d0b.
//
// Solidity: event TokenState(address indexed token, bool state)
func (_Presale *PresaleFilterer) ParseTokenState(log types.Log) (*PresaleTokenState, error) {
	event := new(PresaleTokenState)
	if err := _Presale.contract.UnpackLog(event, "TokenState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
