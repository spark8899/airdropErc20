package contract

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/utils"
)

var TokenContractsMetaData *bind.MetaData

func init() {
	TokenContractsMetaData = getTokenMetaData()
}

func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenContracts, error) {
	parsed, err := TokenContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenContractsMetaData.Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenContracts{TokenContractsCaller: TokenContractsCaller{contract: contract}, TokenContractsTransactor: TokenContractsTransactor{contract: contract}, TokenContractsFilterer: TokenContractsFilterer{contract: contract}}, nil
}

type TokenContracts struct {
	TokenContractsCaller     // Read-only binding to the contract
	TokenContractsTransactor // Write-only binding to the contract
	TokenContractsFilterer   // Log filterer for contract events
}

type TokenContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type TokenContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type TokenContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type TokenContractsSession struct {
	Contract     *TokenContracts   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

type TokenContractsCallerSession struct {
	Contract *TokenContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

type TokenContractsTransactorSession struct {
	Contract     *TokenContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

func NewTokenContracts(address common.Address, backend bind.ContractBackend) (*TokenContracts, error) {
	contract, err := bindTokenContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenContracts{TokenContractsCaller: TokenContractsCaller{contract: contract}, TokenContractsTransactor: TokenContractsTransactor{contract: contract}, TokenContractsFilterer: TokenContractsFilterer{contract: contract}}, nil
}

func bindTokenContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
func (_Contracts *TokenContractsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contracts *TokenContractsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

func (_Contracts *TokenContractsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
func (_Contracts *TokenContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Contracts *TokenContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

func (_Contracts *TokenContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
func (_Contracts *TokenContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Contracts *TokenContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

func (_Contracts *TokenContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
func (_Contracts *TokenContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Contracts *TokenContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

func (_Contracts *TokenContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
func (_Contracts *TokenContractsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", to, amount)
}

func (_Contracts *TokenContractsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, amount)
}

func (_Contracts *TokenContractsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
func (_Contracts *TokenContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_Contracts *TokenContractsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, amount)
}

func (_Contracts *TokenContractsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
func (_Contracts *TokenContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, amount)
}

func (_Contracts *TokenContractsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

func (_Contracts *TokenContractsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
func (_Contracts *TokenContractsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contracts *TokenContractsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

func (_Contracts *TokenContractsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

// Get Token Meta Data
func getTokenMetaData() *bind.MetaData {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config, error: ", err)
	}

	processPath, err := utils.ProcessPath()
	if err != nil {
		log.Fatal("Error get process path ", err)
	}
	_, err = os.Stat(fmt.Sprintf("%s/%s", processPath, ".env"))
	if err != nil {
		processPath = "."
	}

	abiPath := fmt.Sprintf("%s/contracts/%s/%s.abi", processPath, config.TOKEN, config.TOKEN)
	abiString, err := utils.ReadFileStr(abiPath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	bytecodePath := fmt.Sprintf("%s/contracts/%s/%s.bin", processPath, config.TOKEN, config.TOKEN)
	bytecode, err := utils.ReadFileStr(bytecodePath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	return &bind.MetaData{ABI: abiString, Bin: bytecode}
}
