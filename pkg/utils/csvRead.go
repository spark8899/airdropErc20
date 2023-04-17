package utils

import (
	"encoding/csv"
	"os"
)

type CsvDataLines struct {
	Address string
	Amount  string
}

func ReadCsvFile(filename string) ([]CsvDataLines, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	csvData, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return nil, err
	}

	var dataGroup []CsvDataLines
	for i, line := range csvData {
		if i > 0 {
			dataGroup = append(dataGroup,
				CsvDataLines{
					Address: line[0],
					Amount:  line[1],
				})
		}
	}

	return dataGroup, nil
}
