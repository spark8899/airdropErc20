package contract

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/contract"
	"github.com/spark8899/airdropErc20/pkg/utils"
)

func Deploy(contranName string) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config, error: ", err)
	}
	// connect eth client
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		log.Fatal("Cannot connect rpc, error: ", err)
	}

	processPath, err := utils.ProcessPath()
	if err != nil {
		log.Fatal("Error get process path ", err)
	}
	_, err = os.Stat(fmt.Sprintf("%s/%s", processPath, ".env"))
	if err != nil {
		processPath = "."
	}

	abiPath := fmt.Sprintf("%s/contracts/%s/%s.abi", processPath, contranName, contranName)
	abiString, err := utils.ReadFileStr(abiPath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	bytecodePath := fmt.Sprintf("%s/contracts/%s/%s.bin", processPath, contranName, contranName)
	bytecode, err := utils.ReadFileStr(bytecodePath)
	if err != nil {
		log.Fatal("Read abi files error: ", err)
	}

	// get abi
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal("Cannot load abi, error: ", err)
	}

	// get auth
	auth, err := contract.GetAccountAuth(client, config.PRIVATEKEY, config.GASLIMIT)
	if err != nil {
		log.Fatal("Cannot load abi, error: ", err)
	}

	// deploy contract
	address, tx, _, err := bind.DeployContract(
		auth,
		abi,
		common.FromHex(bytecode),
		client,
	)
	if err != nil {
		log.Fatal("Cannot deploy contract, error: ", err)
	}
	fmt.Printf("Deployment transaction: %s\n", tx.Hash().Hex())

	// Waiting for the transaction to be confirmed
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Cannot wait mined, error: ", err)
	}

	fmt.Printf("Contract address: %s\n", address.Hex())
}
