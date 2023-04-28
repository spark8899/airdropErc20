package contract

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
)

func AddrBalance(addrStr string) {
	addrStr = strings.Trim(addrStr, " ")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config, error: ", err)
	}
	// connect eth client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatal("Cannot connect rpc, error: ", err)
	}

	account := common.HexToAddress(addrStr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("%s: %.6f\n", addrStr, ethValue) // address: xxx

	//blockNumber := big.NewInt(5532993)
	//balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balanceAt) // 25729324269165216042

	//fbalance := new(big.Float)
	//fbalance.SetString(balanceAt.String())
	//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//fmt.Println(ethValue) // 25.729324269165216041

	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance) // 25729324269165216042
}
