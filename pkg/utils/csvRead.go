package utils

import (
	"encoding/csv"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/spark8899/airdropErc20/pkg/config"
)

type CsvDataLines struct {
	Address     []common.Address
	Amount      []*big.Int
	TotalAmount *big.Int
}

func ReadCsvFile(filename string) (CsvDataLines, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return CsvDataLines{}, err
	}
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return CsvDataLines{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	csvData, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return CsvDataLines{}, err
	}

	var address []common.Address
	var amount []*big.Int
	totalAmount := big.NewInt(0)
	decimalsDeci := decimal.New(1, int32(config.TOKENDECIMALS))
	for i, line := range csvData {
		if i > 0 {
			address = append(address, common.HexToAddress(line[0]))
			amountDec, err := decimal.NewFromString(line[1])
			if err != nil {
				return CsvDataLines{}, nil
			}
			amountBig := amountDec.Mul(decimalsDeci).BigInt()
			amount = append(amount, amountBig)
			totalAmount = new(big.Int).Add(totalAmount, amountBig)
		}
	}

	dataInfo := CsvDataLines{
		address,
		amount,
		totalAmount,
	}

	return dataInfo, nil
}
