package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spark8899/airdropErc20/pkg/config"
)

func GetAccountAuth(client *ethclient.Client, privateKeyStr string, gasLimit uint64) (*bind.TransactOpts, error) {
	// Build transaction
	privateKeyWithoutPrefix := strings.TrimPrefix(privateKeyStr, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyWithoutPrefix)
	if err != nil {
		log.Fatal("Cannot load privateKey, error: ", err)
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA, error: ", err)
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Cannot get nonce, error: ", err)
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Cannot get chainID, error: ", err)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Cannot get gasPrice, error: ", err)
		return nil, err
	}

	// Construct transaction data
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Cannot get auth, error: ", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func TokenDecimals() (uint8, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config, error: ", err)
		return 0, err
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

	decimals, err := conn.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Cannot get decimals, error: ", err)
		return 0, err
	}
	return decimals, nil
}
