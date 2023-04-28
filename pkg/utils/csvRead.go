package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CsvDataLines struct {
	Address []string
	Amount  []int64
}

func ReadCsvFile(filename string) (CsvDataLines, error) {
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

	var address []string
	var amount []int64
	for i, line := range csvData {
		if i > 0 {
			address = append(address, line[0])
			n, err := strconv.ParseInt(line[1], 10, 64)
			if err != nil {
				return CsvDataLines{}, err
			}
			amount = append(amount, n)
		}
	}

	dataInfo := CsvDataLines{
		address,
		amount,
	}

	return dataInfo, nil
}
