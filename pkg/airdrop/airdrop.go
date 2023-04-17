package airdrop

import (
	"fmt"
	"log"

	"github.com/spark8899/airdropErc20/pkg/config"
	"github.com/spark8899/airdropErc20/pkg/utils"
)

func Airdrop() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println("airdrop files: ", config.AIRDROPFILE)
	fmt.Println("airdrop module.")

	csvData, err := utils.ReadCsvFile(config.AIRDROPFILE)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	for _, info := range csvData {
		fmt.Printf("%s: %s\n", info.Address, info.Amount)
	}
}
