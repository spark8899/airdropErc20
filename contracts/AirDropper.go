// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"dests\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"doAirDrop1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"dests\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"doAirDrop2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6109028061007e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637817573b1161005b5780637817573b146100b25780638da5cb5b146100c5578063ce04d744146100e4578063f2fde38b146100f757600080fd5b80633aeac4e11461008257806351cff8d914610097578063715018a6146100aa575b600080fd5b610095610090366004610601565b61010a565b005b6100956100a5366004610634565b610264565b6100956102c1565b6100956100c0366004610733565b6102d5565b600054604080516001600160a01b039092168252519081900360200190f35b6100956100f23660046107fd565b6103fd565b610095610105366004610634565b6104c2565b61011261053b565b6040516370a0823160e01b815230600482015282906000906001600160a01b038316906370a0823190602401602060405180830381865afa15801561015b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017f9190610854565b9050600081116101cb5760405162461bcd60e51b81526020600482015260126024820152711d1bdad95b881a5b9cdd59999a58da595b9d60721b60448201526064015b60405180910390fd5b6001600160a01b038083169063a9059cbb908516156101ea57846101ec565b335b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303816000875af1158015610239573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025d919061086d565b5050505050565b61026c61053b565b476001600160a01b038216156102825781610284565b335b6001600160a01b03166108fc829081150290604051600060405180830381858888f193505050501580156102bc573d6000803e3d6000fd5b505050565b6102c961053b565b6102d36000610595565b565b6102dd61053b565b80518251146103215760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d040dcdee840dac2e8c6d60831b60448201526064016101c2565b60005b82518110156103f757836001600160a01b031663a9059cbb84838151811061034e5761034e61088f565b60200260200101518484815181106103685761036861088f565b60200260200101516040518363ffffffff1660e01b81526004016103a19291906001600160a01b03929092168252602082015260400190565b6020604051808303816000875af11580156103c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e4919061086d565b50806103ef816108a5565b915050610324565b50505050565b61040561053b565b60005b82518110156103f757836001600160a01b031663a9059cbb8483815181106104325761043261088f565b6020026020010151846040518363ffffffff1660e01b815260040161046c9291906001600160a01b03929092168252602082015260400190565b6020604051808303816000875af115801561048b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104af919061086d565b50806104ba816108a5565b915050610408565b6104ca61053b565b6001600160a01b03811661052f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101c2565b61053881610595565b50565b6000546001600160a01b031633146102d35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101c2565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146105fc57600080fd5b919050565b6000806040838503121561061457600080fd5b61061d836105e5565b915061062b602084016105e5565b90509250929050565b60006020828403121561064657600080fd5b61064f826105e5565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561069557610695610656565b604052919050565b600067ffffffffffffffff8211156106b7576106b7610656565b5060051b60200190565b600082601f8301126106d257600080fd5b813560206106e76106e28361069d565b61066c565b82815260059290921b8401810191818101908684111561070657600080fd5b8286015b848110156107285761071b816105e5565b835291830191830161070a565b509695505050505050565b60008060006060848603121561074857600080fd5b610751846105e5565b925060208085013567ffffffffffffffff8082111561076f57600080fd5b61077b888389016106c1565b9450604087013591508082111561079157600080fd5b508501601f810187136107a357600080fd5b80356107b16106e28261069d565b81815260059190911b820183019083810190898311156107d057600080fd5b928401925b828410156107ee578335825292840192908401906107d5565b80955050505050509250925092565b60008060006060848603121561081257600080fd5b61081b846105e5565b9250602084013567ffffffffffffffff81111561083757600080fd5b610843868287016106c1565b925050604084013590509250925092565b60006020828403121561086657600080fd5b5051919050565b60006020828403121561087f57600080fd5b8151801515811461064f57600080fd5b634e487b7160e01b600052603260045260246000fd5b6000600182016108c557634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220ea2729a94e7ce9ef85309f6f1e0bda0205bbbb14e8da7f9a913588a1e4e5b05a64736f6c63430008130033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// DoAirDrop1 is a paid mutator transaction binding the contract method 0xce04d744.
//
// Solidity: function doAirDrop1(address token, address[] dests, uint256 value) returns()
func (_Contracts *ContractsTransactor) DoAirDrop1(opts *bind.TransactOpts, token common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "doAirDrop1", token, dests, value)
}

// DoAirDrop1 is a paid mutator transaction binding the contract method 0xce04d744.
//
// Solidity: function doAirDrop1(address token, address[] dests, uint256 value) returns()
func (_Contracts *ContractsSession) DoAirDrop1(token common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DoAirDrop1(&_Contracts.TransactOpts, token, dests, value)
}

// DoAirDrop1 is a paid mutator transaction binding the contract method 0xce04d744.
//
// Solidity: function doAirDrop1(address token, address[] dests, uint256 value) returns()
func (_Contracts *ContractsTransactorSession) DoAirDrop1(token common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DoAirDrop1(&_Contracts.TransactOpts, token, dests, value)
}

// DoAirDrop2 is a paid mutator transaction binding the contract method 0x7817573b.
//
// Solidity: function doAirDrop2(address token, address[] dests, uint256[] values) returns()
func (_Contracts *ContractsTransactor) DoAirDrop2(opts *bind.TransactOpts, token common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "doAirDrop2", token, dests, values)
}

// DoAirDrop2 is a paid mutator transaction binding the contract method 0x7817573b.
//
// Solidity: function doAirDrop2(address token, address[] dests, uint256[] values) returns()
func (_Contracts *ContractsSession) DoAirDrop2(token common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DoAirDrop2(&_Contracts.TransactOpts, token, dests, values)
}

// DoAirDrop2 is a paid mutator transaction binding the contract method 0x7817573b.
//
// Solidity: function doAirDrop2(address token, address[] dests, uint256[] values) returns()
func (_Contracts *ContractsTransactorSession) DoAirDrop2(token common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DoAirDrop2(&_Contracts.TransactOpts, token, dests, values)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Contracts *ContractsSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_Contracts *ContractsTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawToken", token, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_Contracts *ContractsSession) WithdrawToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, token, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address token, address to) returns()
func (_Contracts *ContractsTransactorSession) WithdrawToken(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, token, to)
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
