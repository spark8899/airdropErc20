package airdrop

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/contract"
	"github.com/spark8899/airdropErc20/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var AirDropperMetaData *bind.MetaData

const oneGwei = 1e9

func init() {
	AirDropperMetaData = getAirDropperMetaData()
}

func Airdrop() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	gasPriceLimit := new(big.Int).Mul(big.NewInt(int64(config.GASLIMIT)), big.NewInt(oneGwei))
	increaseGas := big.NewInt(config.INCREASEGAS * oneGwei)
	airdropAddress := common.HexToAddress(config.AIRDROPADDR)
	// connect eth client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatal("Cannot connect rpc, error: ", err)
	}
	// get auth
	auth, err := contract.GetAccountAuth(client, config.PRIVATEKEY, config.GASLIMIT)
	if err != nil {
		log.Fatal("Cannot load abi, error: ", err)
	}
	// get callOpts
	callOpts := &bind.CallOpts{
		BlockNumber: nil,
		Context:     context.Background(),
	}
	// get transacOpts
	transacOpts := bind.TransactOpts{
		From:    auth.From,
		Signer:  auth.Signer,
		Context: context.Background(),
	}

	// get airdrop address
	csvData, err := utils.ReadCsvFile(config.AIRDROPFILE)
	if err != nil {
		log.Fatal("Get csv data error:", err)
	}

	// Calculate the number of airdrops
	addrNum := len(csvData.Address)
	amountNum := len(csvData.Amount)
	if addrNum != amountNum {
		panic("The address and amount of the csv file are not equal.")
	}
	var count int
	modNum := addrNum % config.AIRDROPPER
	divNum := addrNum / config.AIRDROPPER
	fmt.Println("Total number of airdrops:", addrNum)
	fmt.Println("airdrop per:", config.AIRDROPPER)
	if modNum > 0 {
		count = divNum + 1
	} else {
		count = divNum
	}
	fmt.Println("airdrop count:", count)
	fmt.Println("Address:", csvData.Address)
	fmt.Println("Amount:", csvData.Amount)
	fmt.Println("TotalAmount:", csvData.TotalAmount)

	token, err := contract.NewTokenContracts(common.HexToAddress(config.TOKENADDR), client)
	if err != nil {
		log.Fatal("Cannot init token contract, error: ", err)
		panic(err)
	}

	airdrop, err := NewAirDropper(airdropAddress, client)
	if err != nil {
		log.Fatal("Cannot init airdrop contract, error: ", err)
		panic(err)
	}

	// ====aprove
	allowance, err := token.Allowance(callOpts, auth.From, airdropAddress)
	if err != nil {
		log.Fatal("Cannot get token Allowance, error: ", err)
		panic(err)
	}

	// check approve needed
	if csvData.TotalAmount.Cmp(allowance) > 0 {
		var willUseGasPrice *big.Int
		var err error
		for {
			willUseGasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				panic(err)
			}
			willUseGasPrice = new(big.Int).Add(willUseGasPrice, increaseGas)
			if willUseGasPrice.Cmp(gasPriceLimit) > 0 {
				time.Sleep(time.Second * 5)
				continue
			} else {
				break
			}
		}
		transacOpts.GasPrice = willUseGasPrice
		tx, err := token.Approve(&transacOpts, airdropAddress, csvData.TotalAmount)
		if err != nil {
			log.Fatal("approve erc20 error: ", err)
			panic(err)
		}
		fmt.Println("approve tx hash: ", tx.Hash().String())
		for {
			_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
			if err == nil && !pending {
				break
			} else {
				if err != nil {
					fmt.Println("approve tx status: ", err)
				} else {
					fmt.Println("approve tx status: is pending...")
				}
				time.Sleep(time.Second * 8)
				continue
			}
		}
		fmt.Println("approve tx ok: ", tx.Hash().String())
	}

	for i := 1; i <= count; i++ {
		// ==== collect address
		var subAddress []common.Address
		var subAmount []*big.Int
		if i == count {
			subAddress = csvData.Address[(i-1)*config.AIRDROPPER : addrNum]
			subAmount = csvData.Amount[(i-1)*config.AIRDROPPER : amountNum]
		} else {
			subAddress = csvData.Address[(i-1)*config.AIRDROPPER : (i-1)*config.AIRDROPPER+config.AIRDROPPER]
			subAmount = csvData.Amount[(i-1)*config.AIRDROPPER : (i-1)*config.AIRDROPPER+config.AIRDROPPER]
		}

		fmt.Printf("airdrop %dth:\n", i)
		fmt.Println(subAddress)
		fmt.Println(subAmount)

		// === send erc20
		fmt.Println("will send transInfo。")
		var willUseGasPrice *big.Int
		var err error
		for {
			willUseGasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				panic(err)
			}
			fmt.Println("suggest gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))
			willUseGasPrice = new(big.Int).Add(willUseGasPrice, increaseGas)
			if willUseGasPrice.Cmp(gasPriceLimit) > 0 {
				time.Sleep(time.Second * 5)
				continue
			} else {
				break
			}
		}
		transacOpts.GasPrice = willUseGasPrice
		fmt.Println("final use gas price: ", new(big.Int).Div(willUseGasPrice, big.NewInt(oneGwei)))
		tx, err := airdrop.DoAirDropMultiple(
			&transacOpts,
			common.HexToAddress(config.TOKENADDR),
			subAddress,
			subAmount)
		if err != nil {
			log.Fatal("send erc20 error: ", err)
			panic(err)
		}
		fmt.Println("tx hash:", tx.Hash().String())
		for {
			_, pending, err := client.TransactionByHash(context.Background(), tx.Hash())
			if err == nil && !pending {
				break
			} else {
				if err != nil {
					fmt.Println("tx status: ", err)
				} else {
					fmt.Println("tx status: is pending...")
				}
				time.Sleep(time.Second * 8)
				continue
			}
		}
		fmt.Println("tx ok: ", tx.Hash().String())
	}
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

// NewAirDropper creates a new instance of AirDropper, bound to a specific deployed contract.
func NewAirDropper(address common.Address, backend bind.ContractBackend) (*AirDropper, error) {
	contract, err := bindAirDropper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AirDropper{AirDropperCaller: AirDropperCaller{contract: contract}, AirDropperTransactor: AirDropperTransactor{contract: contract}, AirDropperFilterer: AirDropperFilterer{contract: contract}}, nil
}

// bindAirDropper binds a generic wrapper to an already deployed contract.
func bindAirDropper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AirDropperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Solidity: function doAirDropSingle(address token, address[] dests, uint256 value) returns()
func (_AirDropper *AirDropperTransactor) DoAirDropSingle(opts *bind.TransactOpts, token common.Address, dests []common.Address, value *big.Int) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "doAirDropSingle", token, dests, value)
}

// Solidity: function doAirDropMultiple(address token, address[] dests, uint256[] values) returns()
func (_AirDropper *AirDropperTransactor) DoAirDropMultiple(opts *bind.TransactOpts, token common.Address, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirDropper.contract.Transact(opts, "doAirDropMultiple", token, dests, values)
}

// Get AirDropper Meta Data
func getAirDropperMetaData() *bind.MetaData {
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

	abiPath := fmt.Sprintf("%s/contracts/%s/%s.abi", processPath, config.AIRDROP, config.AIRDROP)
	abiString, err := utils.ReadFileStr(abiPath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	bytecodePath := fmt.Sprintf("%s/contracts/%s/%s.bin", processPath, config.AIRDROP, config.AIRDROP)
	bytecode, err := utils.ReadFileStr(bytecodePath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	return &bind.MetaData{ABI: abiString, Bin: bytecode}
}
