// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AirDropper

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

// AirDropperMetaData contains all meta data concerning the AirDropper contract.
var AirDropperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"dests\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"doAirDropMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"dests\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"doAirDropSingle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6109298061007e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063715018a61161005b578063715018a6146100bd5780638da5cb5b146100c5578063cb1fcfb6146100e4578063f2fde38b146100f757600080fd5b80630aba5edc146100825780633aeac4e11461009757806351cff8d9146100aa575b600080fd5b6100956100903660046106e7565b61010a565b005b6100956100a536600461074d565b6101dd565b6100956100b8366004610780565b610330565b61009561038d565b600054604080516001600160a01b039092168252519081900360200190f35b6100956100f23660046107a2565b6103a1565b610095610105366004610780565b6104cb565b610112610544565b60005b82518110156101d657846001600160a01b03166323b872dd858584815181106101405761014061087b565b60209081029190910101516040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604481018590526064016020604051808303816000875af115801561019f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c39190610891565b50806101ce816108b3565b915050610115565b5050505050565b6101e5610544565b6040516370a0823160e01b815230600482015282906000906001600160a01b038316906370a0823190602401602060405180830381865afa15801561022e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025291906108da565b90506000811161029e5760405162461bcd60e51b81526020600482015260126024820152711d1bdad95b881a5b9cdd59999a58da595b9d60721b60448201526064015b60405180910390fd5b6001600160a01b038083169063a9059cbb908516156102bd57846102bf565b335b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303816000875af115801561030c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d69190610891565b610338610544565b476001600160a01b0382161561034e5781610350565b335b6001600160a01b03166108fc829081150290604051600060405180830381858888f19350505050158015610388573d6000803e3d6000fd5b505050565b610395610544565b61039f600061059e565b565b6103a9610544565b80518251146103ed5760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d040dcdee840dac2e8c6d60831b6044820152606401610295565b60005b82518110156101d657846001600160a01b03166323b872dd8585848151811061041b5761041b61087b565b60200260200101518585815181106104355761043561087b565b60209081029190910101516040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303816000875af1158015610494573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b89190610891565b50806104c3816108b3565b9150506103f0565b6104d3610544565b6001600160a01b0381166105385760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610295565b6105418161059e565b50565b6000546001600160a01b0316331461039f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610295565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b038116811461060557600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156106495761064961060a565b604052919050565b600067ffffffffffffffff82111561066b5761066b61060a565b5060051b60200190565b600082601f83011261068657600080fd5b8135602061069b61069683610651565b610620565b82815260059290921b840181019181810190868411156106ba57600080fd5b8286015b848110156106dc576106cf816105ee565b83529183019183016106be565b509695505050505050565b600080600080608085870312156106fd57600080fd5b610706856105ee565b9350610714602086016105ee565b9250604085013567ffffffffffffffff81111561073057600080fd5b61073c87828801610675565b949793965093946060013593505050565b6000806040838503121561076057600080fd5b610769836105ee565b9150610777602084016105ee565b90509250929050565b60006020828403121561079257600080fd5b61079b826105ee565b9392505050565b600080600080608085870312156107b857600080fd5b6107c1856105ee565b935060206107d08187016105ee565b9350604086013567ffffffffffffffff808211156107ed57600080fd5b6107f989838a01610675565b9450606088013591508082111561080f57600080fd5b508601601f8101881361082157600080fd5b803561082f61069682610651565b81815260059190911b8201830190838101908a83111561084e57600080fd5b928401925b8284101561086c57833582529284019290840190610853565b979a9699509497505050505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156108a357600080fd5b8151801515811461079b57600080fd5b6000600182016108d357634e487b7160e01b600052601160045260246000fd5b5060010190565b6000602082840312156108ec57600080fd5b505191905056fea264697066735822122015027b9bab96a12c43df0f247acbe3560f38ac2ae6cc9aad532b6eb72a503deb64736f6c63430008130033",
}

// AirDropperABI is the input ABI used to generate the binding from.
// Deprecated: Use AirDropperMetaData.ABI instead.
var AirDropperABI = AirDropperMetaData.ABI

// AirDropperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AirDropperMetaData.Bin instead.
var AirDropperBin = AirDropperMetaData.Bin

// DeployAirDropper deploys a new Ethereum contract, binding an instance of AirDropper to it.
func DeployAirDropper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AirDropper, error) {
	parsed, err := AirDropperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AirDropperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AirDropper{AirDropperCaller: AirDropperCaller{contract: contract}, AirDropperTransactor: AirDropperTransactor{contract: contract}, AirDropperFilterer: AirDropperFilterer{contract: contract}}, nil
}

// AirDropper is an auto generated Go binding around an Ethereum contract.
type AirDropper struct {
	AirDropperCaller     // Read-only binding to the contract
	AirDropperTransactor // Write-only binding to the contract
	AirDropperFilterer   // Log filterer for contract events
}

// AirDropperCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirDropperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirDropperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirDropperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirDropperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirDropperSession struct {
	Contract     *AirDropper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirDropperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirDropperCallerSession struct {
	Contract *AirDropperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AirDropperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirDropperTransactorSession struct {
	Contract     *AirDropperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AirDropperRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirDropperRaw struct {
	Contract *AirDropper // Generic contract binding to access the raw methods on
}

// AirDropperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirDropperCallerRaw struct {
	Contract *AirDropperCaller // Generic read-only contract binding to access the raw methods on
}

// AirDropperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirDropperTransactorRaw struct {
	Contract *AirDropperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirDropper creates a new instance of AirDropper, bound to a specific deployed contract.
func NewAirDropper(address common.Address, backend bind.ContractBackend) (*AirDropper, error) {
	contract, err := bindAirDropper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AirDropper{AirDropperCaller: AirDropperCaller{contract: contract}, AirDropperTransactor: AirDropperTransactor{contract: contract}, AirDropperFilterer: AirDropperFilterer{contract: contract}}, nil
}

// NewAirDropperCaller creates a new read-only instance of AirDropper, bound to a specific deployed contract.
func NewAirDropperCaller(address common.Address, caller bind.ContractCaller) (*AirDropperCaller, error) {
	contract, err := bindAirDropper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirDropperCaller{contract: contract}, nil
}

// NewAirDropperTransactor creates a new write-only instance of AirDropper, bound to a specific deployed contract.
func NewAirDropperTransactor(address common.Address, transactor bind.ContractTransactor) (*AirDropperTransactor, error) {
	contract, err := bindAirDropper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirDropperTransactor{contract: contract}, nil
}

// NewAirDropperFilterer creates a new log filterer instance of AirDropper, bound to a specific deployed contract.
func NewAirDropperFilterer(address common.Address, filterer bind.ContractFilterer) (*AirDropperFilterer, error) {
	contract, err := bindAirDropper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirDropperFilterer{contract: contract}, nil
}

// bindAirDropper binds a generic wrapper to an already deployed contract.
func bindAirDropper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AirDropperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirDropper *AirDropperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AirDropper.Contract.AirDropperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirDropper *AirDropperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropper.Contract.AirDropperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirDropper *AirDropperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirDropper.Contract.AirDropperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirDropper *AirDropperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AirDropper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirDropper *AirDropperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirDropper *AirDropperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirDropper.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropper *AirDropperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AirDropper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropper *AirDropperSession) Owner() (common.Address, error) {
	return _AirDropper.Contract.Owner(&_AirDropper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AirDropper *AirDropperCallerSession) Owner() (common.Address, error) {
	return _AirDropper.Contract.Owner(&_AirDropper.CallOpts)
}

// DoAirDropMultiple is a paid mutator transaction binding the contract method 0xcb1fcfb6.
//
// Solidity: function doAirDropMultiple(address token, address from, address[] dests, uint256[] values) returns()
func (_AirDropper *AirDropperTransactor) DoAirDropMultiple(opts *bind.TransactOpts, token common.Address, from common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "doAirDropMultiple", token, from, dests, values)
}

// DoAirDropMultiple is a paid mutator transaction binding the contract method 0xcb1fcfb6.
//
// Solidity: function doAirDropMultiple(address token, address from, address[] dests, uint256[] values) returns()
func (_AirDropper *AirDropperSession) DoAirDropMultiple(token common.Address, from common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirDropper.Contract.DoAirDropMultiple(&_AirDropper.TransactOpts, token, from, dests, values)
}

// DoAirDropMultiple is a paid mutator transaction binding the contract method 0xcb1fcfb6.
//
// Solidity: function doAirDropMultiple(address token, address from, address[] dests, uint256[] values) returns()
func (_AirDropper *AirDropperTransactorSession) DoAirDropMultiple(token common.Address, from common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirDropper.Contract.DoAirDropMultiple(&_AirDropper.TransactOpts, token, from, dests, values)
}

// DoAirDropSingle is a paid mutator transaction binding the contract method 0x0aba5edc.
//
// Solidity: function doAirDropSingle(address token, address from, address[] dests, uint256 value) returns()
func (_AirDropper *AirDropperTransactor) DoAirDropSingle(opts *bind.TransactOpts, token common.Address, from common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "doAirDropSingle", token, from, dests, value)
}

// DoAirDropSingle is a paid mutator transaction binding the contract method 0x0aba5edc.
//
// Solidity: function doAirDropSingle(address token, address from, address[] dests, uint256 value) returns()
func (_AirDropper *AirDropperSession) DoAirDropSingle(token common.Address, from common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _AirDropper.Contract.DoAirDropSingle(&_AirDropper.TransactOpts, token, from, dests, value)
}

// DoAirDropSingle is a paid mutator transaction binding the contract method 0x0aba5edc.
//
// Solidity: function doAirDropSingle(address token, address from, address[] dests, uint256 value) returns()
func (_AirDropper *AirDropperTransactorSession) DoAirDropSingle(token common.Address, from common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _AirDropper.Contract.DoAirDropSingle(&_AirDropper.TransactOpts, token, from, dests, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropper *AirDropperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropper *AirDropperSession) RenounceOwnership() (*types.Transaction, error) {
	return _AirDropper.Contract.RenounceOwnership(&_AirDropper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AirDropper *AirDropperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AirDropper.Contract.RenounceOwnership(&_AirDropper.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropper *AirDropperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropper *AirDropperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.TransferOwnership(&_AirDropper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AirDropper *AirDropperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.TransferOwnership(&_AirDropper.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AirDropper *AirDropperTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AirDropper *AirDropperSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.Withdraw(&_AirDropper.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AirDropper *AirDropperTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.Withdraw(&_AirDropper.TransactOpts, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_AirDropper *AirDropperTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "withdrawToken", token, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_AirDropper *AirDropperSession) WithdrawToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.WithdrawToken(&_AirDropper.TransactOpts, token, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_AirDropper *AirDropperTransactorSession) WithdrawToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _AirDropper.Contract.WithdrawToken(&_AirDropper.TransactOpts, token, to)
}

// AirDropperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AirDropper contract.
type AirDropperOwnershipTransferredIterator struct {
	Event *AirDropperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AirDropperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirDropperOwnershipTransferred)
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
		it.Event = new(AirDropperOwnershipTransferred)
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
func (it *AirDropperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirDropperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirDropperOwnershipTransferred represents a OwnershipTransferred event raised by the AirDropper contract.
type AirDropperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AirDropper *AirDropperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AirDropperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AirDropper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AirDropperOwnershipTransferredIterator{contract: _AirDropper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AirDropper *AirDropperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirDropperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AirDropper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirDropperOwnershipTransferred)
				if err := _AirDropper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AirDropper *AirDropperFilterer) ParseOwnershipTransferred(log types.Log) (*AirDropperOwnershipTransferred, error) {
	event := new(AirDropperOwnershipTransferred)
	if err := _AirDropper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
