package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func getAccountAuth(client *ethclient.Client, privateKeyStr string, gasLimit uint64) (*bind.TransactOpts, error) {
	// Build transaction
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKeyStr, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		log.Fatal("cannot load privateKey, error: ", err)
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA, error: ", err)
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("cannot get nonce, error: ", err)
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("cannot get chainID, error: ", err)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("cannot get gasPrice, error: ", err)
		return nil, err
	}

	// Construct transaction data
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("cannot deploy contract, error: ", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}
