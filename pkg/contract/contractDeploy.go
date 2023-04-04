package contract

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/utils"
)

func Deploy(contranName string) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config, error: ", err)
	}
	// connect eth client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatal("cannot connect rpc, error: ", err)
	}

	processPath, err := utils.ProcessPath()
	if err != nil {
		log.Fatal("Error get process path ", err)
	}
	processPath = "."

	abiPath := fmt.Sprintf("%s/contracts/%s.abi", processPath, contranName)
	abiString, err := utils.ReadFileStr(abiPath)
	if err != nil {
		log.Fatal("read abi files error: ", err)
	}

	bytecodePath := fmt.Sprintf("%s/contracts/%s.bin", processPath, contranName)
	bytecode, err := utils.ReadFileStr(bytecodePath)
	if err != nil {
		log.Fatal("read abi files error: ", err)
	}

	// get abi
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal("cannot load abi, error: ", err)
	}

	// get auth
	auth, err := getAccountAuth(client, config.PRIVATEKEY, uint64(config.GASLIMIT))
	if err != nil {
		log.Fatal("cannot load abi, error: ", err)
	}

	// deploy contract
	address, tx, _, err := bind.DeployContract(
		auth,
		abi,
		common.FromHex(bytecode),
		client,
	)
	if err != nil {
		log.Fatal("cannot deploy contract, error: ", err)
	}
	fmt.Printf("Contract deployment transaction: %s\n", tx.Hash().Hex())

	// Waiting for the transaction to be confirmed
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("cannot wait mined, error: ", err)
	}

	fmt.Printf("Contract deployment address: %s\n", address.Hex())
}
