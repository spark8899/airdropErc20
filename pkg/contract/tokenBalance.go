package contract

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
)

func TokenBalance(addrStr string) (*big.Float, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config, error: ", err)
	}
	// connect eth client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatal("Cannot connect rpc, error: ", err)
	}

	conn, err := NewTokenContracts(common.HexToAddress(config.TOKENADDR), client)
	if err != nil {
		log.Fatal("Cannot init contract, error: ", err)
		panic(err)
	}

	balance, err := conn.BalanceOf(&bind.CallOpts{}, common.HexToAddress(addrStr))
	if err != nil {
		log.Fatal("Cannot get balanceOf, error: ", err)
	}
	decimals, err := conn.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Cannot get decimals, error: ", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("%s: %.6f\n", addrStr, ethValue)

	return ethValue, nil
}
